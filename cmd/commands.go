/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package main

import (
	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/
const (
	// Govee API commands
	GOVEE_CMD_TURN       string = "turn"       // TurnOn & TurnOff
	GOVEE_CMD_BRIGHTNESS string = "brightness" // SetBrightness(0..100)
	GOVEE_CMD_COLOR      string = "color"      // SetColor(rgbColor)
	GOVEE_CMD_COLORTEM   string = "colorTem"   // SetColorTem(colorTemperature)
)

/* ----------------------------------------------------------------
 *						I n t e r f a c e s
 *-----------------------------------------------------------------*/

type ICommand interface {
	name() string
	execute() error
}

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type GoveeCommand struct {
	Client  *veex.Client
	Address string
	Model   string
}

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

// determine whether device is a light
func isLight(d veex.Device) (bool, veex.Range) {
	var light bool = false
	var lrange veex.Range
	colorTem := d.Properties.ColorTem
	if colorTem.Range.Min > 0 && colorTem.Range.Max > 0 {
		light = true
		lrange = colorTem.Range
	}
	return light, lrange
}

func hasLightControl(dev veex.Device) bool {
	for _, v := range dev.SupportCmds {
		if v == GOVEE_CMD_BRIGHTNESS ||
			v == GOVEE_CMD_COLOR ||
			v == GOVEE_CMD_COLORTEM {
			return true
		}
	}
	return false
}
