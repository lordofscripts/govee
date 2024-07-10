/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * TURN OFF DEVICE COMMAND
 *-----------------------------------------------------------------*/
package main

import (
	"fmt"

	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type OffCommand struct {
	GoveeCommand
}

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

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

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (c *OffCommand) name() string {
	return fmt.Sprintf("TurnOff")
}

// implements ICommand for OffCommand
func (c *OffCommand) execute() error {
	controlRequest, _ := c.Client.Device(c.Address, c.Model).TurnOff()
	_, err := c.Client.Run(controlRequest)
	return err
}
