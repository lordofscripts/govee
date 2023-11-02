/* -----------------------------------------------------------------
 *				  https://allmylinks.com/lordofscripts
 *				  Copyright (C)2023 Lord of Scripts(tm)
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package govee

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
)

const (
	GOVEE_ENV = "GOVEE_API"		// environment variable holding API KEY
	CURRENT_VERSION = "1.0"
	MODEL	= 'M'
	MAC		= 'N'
	ALIAS	= 'A'
	LOCATION = 'L'
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/
type Field	rune
type DeviceCollection []GoveeDevice

type GoveeDevice struct {
	Model		string	`json:"model"`
	MacAddress	string	`json:"mac"`
	Alias		string	`json:"alias"`
	Location	string	`json:"location"`
}

type GoveeConfig struct {
	Version	string			 `json:"version"`
	ApiKey	string			 `json:"apiKey"`
	Devices	DeviceCollection `json:"devices"`
}

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/
func NewConfig() *GoveeConfig {
	c := &GoveeConfig{
		Version: CURRENT_VERSION,
		ApiKey: os.Getenv(GOVEE_ENV),
		Devices: make([]GoveeDevice,0),
	}
	return c
}
/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/
func (d GoveeDevice) String() string {
	return fmt.Sprintf("%s @%s %q", d.Model, d.MacAddress, d.Alias)
}

func (q DeviceCollection) Where(fld Field, value string) DeviceCollection {
	var selected DeviceCollection
	selected = make([]GoveeDevice,0)
	value = strings.Trim(value, " \t")
	if len(q) > 0 {
		for _,v := range q {
			switch fld {
				case MODEL:
					if strings.ToUpper(v.Model) == strings.ToUpper(value) {
						selected = append(selected, v)
					}
				case MAC:
					if strings.ToUpper(v.MacAddress) == strings.ToUpper(value) {
						selected = append(selected, v)
					}
				case ALIAS:
					if strings.ToUpper(v.Alias) == strings.ToUpper(value) {
						selected = append(selected, v)
					}
				case LOCATION:
					if strings.ToUpper(v.Location) == strings.ToUpper(value) {
						selected = append(selected, v)
					}
			}
		}
	}
	return selected
}

func (q DeviceCollection) Count() int {
	if q == nil {
		return 0
	}
	return len(q)
}
/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/
// Read the configureation file. Else return a default configuration.
// The default configuration reads the API Key from the environment.
func Read(filename string) *GoveeConfig {
	if _, err := os.Stat(filename); err != nil {
		fmt.Printf("WARN: Cannot read config %s", filename)
		return NewConfig()
	}

	// read existing
    fd, err := os.Open(filename)
    defer fd.Close()

    if err != nil {
        fmt.Println(err)
        return NewConfig()
    }

	var cfg GoveeConfig
    jsonParser := json.NewDecoder(fd)
    if err := jsonParser.Decode(&cfg); err != nil {
		fmt.Printf("read JSON ERR %v\n", err)
		return NewConfig()
    }

    if len(cfg.ApiKey) == 0 {
		cfg.ApiKey = os.Getenv(GOVEE_ENV)
    }
    return &cfg
}

// Get the API Key only. If not present in the environment, then fallback
// to the configuration file (if any).
func GetAPI(filename string) string {
	key,set := os.LookupEnv(GOVEE_ENV)
	key = strings.Trim(key," \t")
	if len(key) == 0 {
		// try to read config file
		if len(filename) == 0 && !set{
			fmt.Printf("Please set your API key on environment %q or create config file\n", GOVEE_ENV)
			return ""
		}

		cfg := Read(filename)
		return cfg.ApiKey
	}
	return key
}

 