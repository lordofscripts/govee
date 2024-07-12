/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * (LIGHTS ONLY) SET BRIGHTNESS COMMAND
 *-----------------------------------------------------------------*/
package main

import (
	"fmt"

	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type BrightnessCommand struct {
	GoveeCommand
	brightness int
}

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

func newCmdBrightness(clientPtr *veex.Client, address, model string, value uint) *BrightnessCommand {
	// verify it is a light device
	if dev := clientPtr.Device(address, model); dev != nil {
		if hasLightControl(*dev) {
			o := &BrightnessCommand{
				GoveeCommand: GoveeCommand{
					Client:  clientPtr,
					Address: address,
					Model:   model,
				},
				brightness: int(value),
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

func (c *BrightnessCommand) name() string {
	return fmt.Sprintf("Brightness %d", c.brightness)
}

// implements ICommand for OnCommand
func (c *BrightnessCommand) execute() error {
	var controlRequest veex.DeviceControlRequest
	var err error
	controlRequest, err = c.Client.Device(c.Address, c.Model).SetBrightness(c.brightness)
	if err == nil {
		//var rsp veex.GoveeResponse
		_, err = c.Client.Run(controlRequest) // GoveeResponse, error
		//fmt.Printf("Response %#+v\n", rsp)
	}

	return err
}
