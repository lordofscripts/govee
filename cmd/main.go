/* -----------------------------------------------------------------
 *				  https://allmylinks.com/lordofscripts
 *				  Copyright (C)2023 Lord of Scripts(tm)
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * Utility program that uses GOVEE API to control (On/Off/List/State)
 * GOVEE smart devices in your network.
 *-----------------------------------------------------------------*/
package main
// https://govee-public.s3.amazonaws.com/developer-docs/GoveeDeveloperAPIReference.pdf

import (
    "flag"
    "fmt"
    "os"
    "path"
    "strings"
    "time"

    "github.com/lordofscripts/govee"
    "github.com/lordofscripts/govee/util"
    veex "github.com/loxhill/go-vee"
)

const (
	HOME_ENV = "HOME"	// environment var. holding user home directory
	MY_CONFIG = ".config/govee.json" // config file relative to HOME_ENV
	API_KEY = ""

	RETVAL_OK = 0
	RETVAL_CLI_MISSING int = 1
	RETVAL_CLI_INVALID int = 2
	RETVAL_CFG_ALIAS int = 3
	RETVAL_CFG_MODEL int = 4
	RETVAL_CFG_API int = 5
	RETVAL_CMD_EXEC_ABORT int = 6
	RETVAL_CLI_INIT int = 7
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

// look up our ~/.config/govee.json for entries matching MAC
func findByMAC(mac string) *govee.GoveeDevice {
	cfgFilename := path.Join(os.Getenv(HOME_ENV), MY_CONFIG)
	mac = strings.ToUpper(mac)
	cfg := govee.Read(cfgFilename)
	candidates := cfg.Devices.Where(govee.MAC, mac)
	cnt := candidates.Count()
	if cnt > 0 {
		for _,v := range candidates {
			if strings.ToUpper(v.MacAddress) == mac {
				return &v
			}
		}
	}

	return nil
}

// print a message and terminate application execution
func die(retVal int, format string, v ...any) {
	fmt.Printf(format + "\n", v...)
	os.Exit(retVal)
}

func getHelp() {
	fmt.Printf("%s by Lord of Scripts\n", govee.Version)
	fmt.Println("Usage:")
	fmt.Println("   govee -list")
	fmt.Println("   govee -mac {MAC_ADDRESS} -model {MODEL_NUMBER} [device command]")
	// these two need a config file with entries
	fmt.Println("   govee -id {ALIAS} [device command]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println("\t*** ***")
}

/* ----------------------------------------------------------------
 *							M A I N
 *-----------------------------------------------------------------*/
func main() {
	const (
		DEF_BRIGHT uint = 101
	)
	fmt.Printf("\t\t../ %s (c)2023-%d Lord of Scripts \\..\n", govee.Version, time.Now().Year())
	fmt.Println("\t\t    https://allmylinks.com/lordofscripts")
	// declare real CLI options
	var optHelp, optList, optOn, optOff, optState, optInit bool
	// declare CLI options which have an explicit value
	var inDevice, inModel, inAlias, inColor string
	var inBright uint
	flag.BoolVar(&optHelp, "help", false, "This help")
	flag.BoolVar(&optInit, "init", false, "Create sample config file")
	flag.BoolVar(&optList, "list", false, "List devices")
	flag.BoolVar(&optOn, "on", false, "Turn ON [device command]")
	flag.BoolVar(&optOff, "off", false, "Turn OFF [device command]")
	flag.BoolVar(&optState, "state", false, "Request online/power state of device [device command]")
	flag.StringVar(&inDevice, "mac", "", "Device MAC")
	flag.StringVar(&inModel, "model", "", "Device Model")
	flag.StringVar(&inAlias, "id", "", "Device alias (derive Model & MAC from this)")
	flag.StringVar(&inColor, "color", "", "Hexadecimal RGB color code, i.e. ffa512")
	flag.UintVar(&inBright, "brightness", DEF_BRIGHT, "Light brightness (0..100) [device command]")
	flag.Parse()

	// declare virtual options
	var voptBright, voptColor bool
	// virtual options based on values
	if inBright < DEF_BRIGHT {
		voptBright = true
	}
	if inColor != "" {
		voptColor = true
	}

	// any command given or at least help?
	if optHelp || !(optList || optOn || optOff || optState || optInit || voptBright || voptColor) {
		getHelp()
		if optHelp {
			os.Exit(RETVAL_OK)
		}

		die(RETVAL_CLI_MISSING, "Missing CLI options")
	}

	if optInit {
		if created, err := govee.CreateSampleConfigFile(); err != nil {
			die(RETVAL_CLI_INIT, err.Error())
		} else {
			if created {
				die(RETVAL_OK, "Created configuration file. Please edit it.")
			} else {
				die(RETVAL_OK, "There is already a config file.")
			}
		}
	}

	// state && on & off are mutually exclusive
	if !util.One(optOn, optOff, optState, optList, voptBright, voptColor) {
		getHelp()

		die(RETVAL_CLI_INVALID, "Decide either -on OR -off OR -state OR -list")
	}

	// Config
	cfgFilename := path.Join(os.Getenv(HOME_ENV), MY_CONFIG)

	if len(inAlias) != 0 {
		cfg := govee.Read(cfgFilename)
		candidates := cfg.Devices.Where(govee.ALIAS, inAlias)
		cnt := candidates.Count()
		if cnt == 0 {
			die(RETVAL_CFG_ALIAS, "Could not find alias %q in repository\n", inAlias)
		}

		if cnt > 1 {
			die(RETVAL_CFG_ALIAS, "Found %d entries. Alias not unique, please correct config\n", cnt)
		}

		fmt.Println("\tFound  ", candidates[0], "\n\tAt     :", candidates[0].Location)
		inDevice = candidates[0].MacAddress
		inModel = candidates[0].Model
	}

	// with STATE, ON & OFF commands DEVICE & MODEL are required
	if (optOn || optOff || optState) && ((len(inDevice) == 0) && (len(inModel) == 0)) {
		getHelp()

		die(RETVAL_CFG_MODEL, "-dev MAC and -model MODEL options are required!")
	} else {
		inDevice = strings.ToUpper(inDevice)
		inModel = strings.ToUpper(inModel)
	}

	key := govee.GetAPI(cfgFilename)
	if len(key) == 0 {
		die(RETVAL_CFG_API, "You forgot to set your GOVEE API Key!")
	}

	// TurnOff, TurnOn, SetBrightness, SetColor, SetColorTem
	client := veex.New(key)
	var command ICommand
	switch {
		case optList:
			command = newCmdList(client)

		case optOn:
			command = newCmdOn(client, inDevice, inModel)

		case optOff:
			command = newCmdOff(client, inDevice, inModel)

		case optState:
			command = newCmdState(client, inDevice, inModel)

		case voptBright:
			command = newCmdBrightness(client, inDevice, inModel, inBright)

		case voptColor:
			if !strings.HasPrefix(inColor, "#") {
				inColor = "#" + inColor
			}
			if color, err := util.ParseHexColorGovee(inColor); err == nil {
				command = newCmdColor(client, inDevice, inModel, color)
			} else {
				die(RETVAL_CLI_INVALID, "Invalid hex RGB color code: %s", inColor)
			}
	}

	if err := command.execute(); err != nil {
		//fmt.Printf("Error type: %T\nMessage: %s\n%v\n", err, err, err)
		fmt.Printf("\tCommand: %s\n\tError  : %v\n", command.name(), err)
	} else {
		fmt.Printf("\tCommand: %s\n", command.name())
	}
}
