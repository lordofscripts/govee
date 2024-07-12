/* -----------------------------------------------------------------
 *				  https://allmylinks.com/lordofscripts
 *				  Copyright (C)2023 Lord of Scripts(tm)
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package govee

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	GOVEE_ENV          = "GOVEE_API" // environment variable holding API KEY
	MIN_CONFIG_VERSION = "1.0" // minimum configuration file version
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type GoveeConfig struct {
	Version string           `json:"version"`
	ApiKey  string           `json:"apiKey"`
	Devices DeviceCollection `json:"devices"`
}

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/
func NewConfig() *GoveeConfig {
	c := &GoveeConfig{
		Version: MIN_CONFIG_VERSION,
		ApiKey:  os.Getenv(GOVEE_ENV),
		Devices: make([]GoveeDevice, 0),
	}
	return c
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/


/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

func CreateSampleConfigFile() (bool, error) {
	created := false
	filename := path.Join(os.Getenv("HOME"), ".config/govee.json")
	if _, err := os.Stat(filename); err != nil {
		sample := NewConfig()
		if key, set := getEnvAPI(); set {
			sample.ApiKey = key
		}

		// sampleDev.IsValid() would return false
		sampleDev := GoveeDevice{"", cEMPTY_MAC, "", "Garage"}
		sample.Devices = append(sample.Devices, sampleDev)

		if fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0640); err != nil {
			return false, err
		} else {
			defer fd.Close()

			jsonEncoder := json.NewEncoder(fd)
			jsonEncoder.SetIndent("", "\t")
			if err := jsonEncoder.Encode(sample); err != nil {
				return false, err
			} else {
				created = true
			}
		}
	}

	return created, nil
}

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
	key, set := getEnvAPI()
	if len(key) == 0 {
		// try to read config file
		if len(filename) == 0 && !set {
			fmt.Printf("Please set your API key on environment %q or create config file\n", GOVEE_ENV)
			return ""
		}

		cfg := Read(filename)
		return cfg.ApiKey
	}
	return key
}

// attempt to retrieve Govee API key from environment variable
func getEnvAPI() (string, bool) {
	key, set := os.LookupEnv(GOVEE_ENV)
	if !set {
		return "", false
	}

	key = strings.Trim(key, " \t")
	return key, true
}
