/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * (LIGHTS ONLY) SET COLOR COMMAND
 *-----------------------------------------------------------------*/
package main

import (
	"fmt"

	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type ColorCommand struct {
	GoveeCommand
	color veex.Color
}

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

func newCmdColor(clientPtr *veex.Client, address, model string, color veex.Color) *ColorCommand {
	// verify it is a light device
	if dev := clientPtr.Device(address, model); dev != nil {
		if hasLightControl(*dev) {
			o := &ColorCommand{
				GoveeCommand: GoveeCommand{
					Client:  clientPtr,
					Address: address,
					Model:   model,
				},
				color: color,
			}

			return o
		}
	}

	die(RETVAL_CMD_EXEC_ABORT, "Device model %s %q is not a LIGHT\n", model, address)
	return nil
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (c *ColorCommand) name() string {
	return fmt.Sprintf("Color #%02x%02x%02x", c.color.R, c.color.G, c.color.B)
}

// implements ICommand for OnCommand
func (c *ColorCommand) execute() error {
	var controlRequest veex.DeviceControlRequest
	var err error
	controlRequest, err = c.Client.Device(c.Address, c.Model).SetColor(c.color)
	if err == nil {
		//var rsp veex.GoveeResponse
		_, err = c.Client.Run(controlRequest) // GoveeResponse, error
		//fmt.Printf("Response %#+v\n", rsp)
	}

	return err
}
