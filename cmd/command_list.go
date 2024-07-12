/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * LIST LOCAL DEVICES COMMAND
 *-----------------------------------------------------------------*/
package main

import (
	"fmt"
	"strings"

	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type ListCommand struct {
	GoveeCommand
}

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

func newCmdList(clientPtr *veex.Client) *ListCommand {
	o := &ListCommand{
		GoveeCommand: GoveeCommand{
			Client:  clientPtr,
			Address: "",
			Model:   "",
		},
	}
	return o
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (c *ListCommand) name() string {
	return fmt.Sprintf("List")
}

// implements ICommand for ListCommand
func (c *ListCommand) execute() error {
	listRequest := c.Client.ListDevices()
	if response, err := c.Client.Run(listRequest); err != nil {
		return err
	} else {
		devices := response.Devices()
		for i, d := range devices {
			dev := findByMAC(d.Device)
			var alias string = "(please add to config file)"
			var controllable string = "(can't be controlled)"
			if dev != nil {
				alias = dev.Alias
			}
			if d.Controllable {
				controllable = ""
			}
			fmt.Printf("#%d %q\n\tDevice: %s\n\tModel : %s %s\n\tType  : %s\n",
				i+1,
				alias,
				d.Device,
				d.Model,
				controllable,
				d.DeviceName)
			fmt.Printf("\tCommands: %s\n", strings.Join(d.SupportCmds, "|"))
			if is, rng := isLight(d); is {
				fmt.Printf("\tLights: %t Temp. Range:(%d, %d) Kelvin\n", is, rng.Min, rng.Max)
			}
		}
	}
	return nil
}
