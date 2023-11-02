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

    "lordofscripts/govee"
    "lordofscripts/govee/util"
    veex "github.com/loxhill/go-vee"
)

const (
	HOME_ENV = "HOME"	// environment var. holding user home directory
	MY_CONFIG = ".config/govee.json" // config file relative to HOME_ENV
	API_KEY = ""
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/
type ICommand interface {
	execute() error
}

type GoveeCommand struct {
	Client	*veex.Client
	Address	string
	Model	string
}

/* ----------------------------------------------------------------
 *							T y p e s
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *					OnCommand::GoveeCommand
 *-----------------------------------------------------------------*/
type OnCommand struct {
	GoveeCommand
}

func newCmdOn(clientPtr *veex.Client, address, model string) *OnCommand {
	 o := &OnCommand{
	 	GoveeCommand: GoveeCommand{
	 		Client: clientPtr,
		 	Address: address,
		 	Model: model,
	 	},
	 }
	 return o
}

// implements ICommand for OnCommand
func (c *OnCommand) execute() error {
	controlRequest, _ := c.Client.Device(c.Address, c.Model).TurnOn()
	_, err := c.Client.Run(controlRequest)
	return err
}

/* ----------------------------------------------------------------
 *							T y p e s
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *					OffCommand::GoveeCommand
 *-----------------------------------------------------------------*/
type OffCommand struct {
	GoveeCommand
}

func newCmdOff(clientPtr *veex.Client, address, model string) *OffCommand {
	 o := &OffCommand{
	 	GoveeCommand: GoveeCommand{
	 		Client: clientPtr,
		 	Address: address,
		 	Model: model,
	 	},
	 }
	 return o
}

// implements ICommand for OffCommand
func (c *OffCommand) execute() error {
	controlRequest, _ := c.Client.Device(c.Address, c.Model).TurnOff()
	_, err := c.Client.Run(controlRequest)
	return err
}

/* ----------------------------------------------------------------
 *							T y p e s
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *					StateCommand::GoveeCommand
 *-----------------------------------------------------------------*/

// the State commands obtains the online and powered state of a device
// along its h/w address and model-
type StateCommand struct {
	GoveeCommand
}

func newCmdState(clientPtr *veex.Client, address, model string) *StateCommand {
	 o := &StateCommand{
	 	GoveeCommand: GoveeCommand{
	 		Client: clientPtr,
		 	Address: address,
		 	Model: model,
	 	},
	 }
	 return o
}

// implements ICommand for StateCommand
func (c *StateCommand) execute() error {
	stateRequest := c.Client.Device(c.Address, c.Model).State()
	// {Code:200 Message:Success Data:{Device:D7:B6:60:74:F4:02:D5:A2 Model:H5083 Properties:[map[online:true] map[powerState:off]] Devices:[]}}
	rsp, err := c.Client.Run(stateRequest)	// govee.GoveeResponse
	//fmt.Printf("%+v\n", rsp)
	if rsp.Code != 200 {
		err = fmt.Errorf("State request failed: %q", rsp.Message)
	} else {
		fmt.Println("\tMAC    :", rsp.Data.Device)
		fmt.Println("\tModel  :", rsp.Data.Model)
		fmt.Println("\tOnline :", rsp.Data.Properties[0]["online"])
		fmt.Println("\tPowered:", rsp.Data.Properties[1]["powerState"])
	}
	return err
}
/* ----------------------------------------------------------------
 *							T y p e s
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *					ListCommand::GoveeCommand
 *-----------------------------------------------------------------*/
type ListCommand struct {
	GoveeCommand
}

func newCmdList(clientPtr *veex.Client) *ListCommand {
	 o := &ListCommand{
	 	GoveeCommand: GoveeCommand{
	 		Client: clientPtr,
		 	Address: "",
		 	Model: "",
	 	},
	 }
	 return o
}

// implements ICommand for ListCommand
func (c *ListCommand) execute() error {
	listRequest := c.Client.ListDevices()
	if response, err := c.Client.Run(listRequest); err != nil {
		return err
	} else {
		devices := response.Devices()
		for i,d := range devices {
			fmt.Printf("#%d\n\tDevice: %s\n\tModel : %s\n\tType  : %s\n", i+1, d.Device, d.Model, d.DeviceName)
		}
	}
	return nil
}

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

func getHelp() {
	fmt.Printf("Govee v%s by Lord of Scripts\n", govee.CURRENT_VERSION)
	fmt.Println("Usage:")
	fmt.Println("   govee -list")
	fmt.Println("   govee -mac {MAC_ADDRESS} -model {MODEL_NUMBER} -on")
	fmt.Println("   govee -mac {MAC_ADDRESS} -model {MODEL_NUMBER} -off")
	fmt.Println("   govee -mac {MAC_ADDRESS} -model {MODEL_NUMBER} -state")
	// these two need a config file with entries
	fmt.Println("   govee -id {ALIAS} -on")
	fmt.Println("   govee -id {ALIAS} -off")
	fmt.Println("   govee -id {ALIAS} -state")
	flag.PrintDefaults()
	fmt.Println("\t*** ***")
}

/* ----------------------------------------------------------------
 *							M A I N
 *-----------------------------------------------------------------*/
func main() {
	fmt.Printf("\t\t../ GoveeLux v%s (c)2023 Lord of Scripts \\..\n", govee.CURRENT_VERSION)
	fmt.Println("\t\t    https://allmylinks.com/lordofscripts")

	var optHelp, optList, optOn, optOff, optState bool
	var inDevice, inModel, inAlias string
	flag.BoolVar(&optHelp, "help", false, "This help")
	flag.BoolVar(&optList, "list", false, "List devices")
	flag.BoolVar(&optOn, "on", false, "Turn ON")
	flag.BoolVar(&optOff, "off", false, "Turn OFF")
	flag.BoolVar(&optState, "state", false, "Request online/power state of device")
	flag.StringVar(&inDevice, "mac", "", "Device MAC")
	flag.StringVar(&inModel, "model", "", "Device Model")
	flag.StringVar(&inAlias, "id", "", "Device alias (derive Model & MAC from this)")
	flag.Parse()

	// any command given or at least help?
	if optHelp || !(optList || optOn || optOff || optState) {
		getHelp()
		if optHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
	// state && on & off are mutually exclusive
	if !util.One(optOn, optOff, optState, optList) {
		getHelp()
		fmt.Println("Decide either -on OR -off OR -state OR -list")
		os.Exit(1)
	}

	// Config
	cfgFilename := path.Join(os.Getenv(HOME_ENV), MY_CONFIG)

	if len(inAlias) != 0 {
		cfg := govee.Read(cfgFilename)
		candidates := cfg.Devices.Where(govee.ALIAS, inAlias)
		cnt := candidates.Count()
		if cnt == 0 {
			fmt.Printf("Could not find alias %q in repository\n", inAlias)
			os.Exit(2)
		}

		if cnt > 1 {
			fmt.Printf("Found %d entries. Alias not unique, please correct config\n", cnt)
			os.Exit(2)
		}

		fmt.Println("\tFound  ", candidates[0], "\n\tAt     :", candidates[0].Location)
		inDevice = candidates[0].MacAddress
		inModel = candidates[0].Model
	}

	// with STATE, ON & OFF commands DEVICE & MODEL are required
	if (optOn || optOff || optState) && ((len(inDevice) == 0) && (len(inModel) == 0)) {
		getHelp()
		fmt.Println("-dev MAC and -model MODEL options are required!")
		os.Exit(1)
	} else {
		inDevice = strings.ToUpper(inDevice)
		inModel = strings.ToUpper(inModel)
	}

	key := govee.GetAPI(cfgFilename)
	if len(key) == 0 {
		fmt.Println("You forgot to set your GOVEE API Key!")
		os.Exit(2)
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
	}

	if err := command.execute(); err != nil {
		//fmt.Printf("Error type: %T\nMessage: %s\n%v\n", err, err, err)
		fmt.Printf("\tError  : %v\n", err)
	}
}