/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * DEVICE INFO & STATE COMMAND
 *-----------------------------------------------------------------*/
package main

import (
	"fmt"

	veex "github.com/loxhill/go-vee"

	"lordofscripts/govee"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

// the State commands obtains the online and powered state of a device
// along its h/w address and model-
type StateCommand struct {
	GoveeCommand
}


/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/

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

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (c *StateCommand) name() string {
	return fmt.Sprintf("State")
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
		if dprop := govee.NewGoveeDataProperties(rsp.Data); dprop != nil {
			fmt.Println("\tMAC    :", dprop.Address)	// rsp.Data.Device
			fmt.Println("\tModel  :", dprop.Model)		// rsp.Data.Model
			// rsp.Data.Properties[]
			fmt.Println("\tOnline :", dprop.Online)
			fmt.Println("\tPowered:", dprop.Powered)
			light := dprop.IsLight()
			fmt.Println("\tLight? :", light)
			if light {
				fmt.Println("\t\tBrightness :", dprop.Brightness)
				fmt.Println("\t\tColor      :", dprop.Color)
				fmt.Println("\t\tTemperature:", dprop.Temperature)
			}
		}
	}
	return err
}

