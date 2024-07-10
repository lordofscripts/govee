/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package govee

import (
	"fmt"
	"log"
	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/


/* ----------------------------------------------------------------
 *						I n t e r f a c e s
 *-----------------------------------------------------------------*/

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type GoveeDataProperties struct {
	Address		string
	Model		string
	Online		bool
	Powered		string
	Brightness	int
	Temperature	int
	Color		string
}

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

// takes a Govee API response and parses the Data payload
func NewGoveeDataProperties(data veex.ResponseData) *GoveeDataProperties {
	dprop := &GoveeDataProperties{data.Device, data.Model, false, "?", -1, -1, ""}

	if data.Properties != nil {
		if err := dprop.Parse(data.Properties); err != nil {
			log.Println("ERROR PARSING GoveeDataProperty")
			dprop = nil
		}
	}

	return dprop
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (g *GoveeDataProperties) Parse(properties []map[string]any) error {
	var funcColorToHex = func(color veex.Color) string {
		return fmt.Sprintf("RGB(#%02x%02x%02x)", color.R, color.G, color.B)
	}
	var err error

	for _, prop := range properties {
		//fmt.Printf("\t#%d %#v\n", i, prop)
		for key,value := range prop {
			switch key {
				case "online":
					g.Online = value.(bool)
					break
				case "powerState":
					g.Powered = value.(string)
					break
				case "brightness":
					g.Brightness= int(value.(float64))
					break
				case "color":
					if colormap, ok := value.(map[string]any); ok {
						color := veex.Color{
							R: int(colormap["r"].(float64)),
							G: int(colormap["g"].(float64)),
							B: int(colormap["b"].(float64)),
						}
						g.Color = funcColorToHex(color)
					} else {
						err = fmt.Errorf("Not a color! %#v", value)
					}
				case "colorTem":
					err = fmt.Errorf("colorTem property parsing not supported")
					break
			}
		}
	}

	return err
}

// whether the device in the response appears to be a light
func (g *GoveeDataProperties) IsLight() bool {
	return g.Brightness > -1 && g.Color != ""
}

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/


