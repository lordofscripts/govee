/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * TURN ON DEVICE COMMAND
 *-----------------------------------------------------------------*/
package main

import (
	"fmt"

	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type OnCommand struct {
	GoveeCommand
}

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

func newCmdOn(clientPtr *veex.Client, address, model string) *OnCommand {
	o := &OnCommand{
		GoveeCommand: GoveeCommand{
			Client:  clientPtr,
			Address: address,
			Model:   model,
		},
	}
	return o
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (c *OnCommand) name() string {
	return fmt.Sprintf("TurnOn")
}

// implements ICommand for OnCommand
func (c *OnCommand) execute() error {
	controlRequest, _ := c.Client.Device(c.Address, c.Model).TurnOn()
	_, err := c.Client.Run(controlRequest)
	return err
}
