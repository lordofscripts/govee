/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Lord of Scripts(tm)
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * GOVEE DEVICE & DEVICE COLLECTION
 *-----------------------------------------------------------------*/
package govee

import (
	"fmt"
	"strings"
)

const (
	FieldMODEL    Field = 'M'
	FieldMAC      Field = 'N'
	FieldALIAS    Field = 'A'
	FieldLOCATION Field = 'L'

	cEMPTY_MAC = "00:00:00:00:00:00:00:00"
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

type Field rune

type DeviceCollection []GoveeDevice

type GoveeDevice struct {
	Model      string `json:"model"`
	MacAddress string `json:"mac"`
	Alias      string `json:"alias"`
	Location   string `json:"location"`
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func (d GoveeDevice) String() string {
	return fmt.Sprintf("%s @%s %q", d.Model, d.MacAddress, d.Alias)
}

func (d GoveeDevice) IsValid() bool {
	if len(d.Model) > 0 && len(d.Alias) > 0 &&
		len(d.MacAddress) > 0 && d.MacAddress != cEMPTY_MAC {
		return true
	}

	return false
}

func (q DeviceCollection) Where(fld Field, value string) DeviceCollection {
	var selected DeviceCollection
	selected = make([]GoveeDevice, 0)
	value = strings.Trim(value, " \t")
	if len(q) > 0 {
		for _, v := range q {
			switch fld {
			case FieldMODEL:
				if strings.ToUpper(v.Model) == strings.ToUpper(value) {
					selected = append(selected, v)
				}
			case FieldMAC:
				if strings.ToUpper(v.MacAddress) == strings.ToUpper(value) {
					selected = append(selected, v)
				}
			case FieldALIAS:
				if strings.ToUpper(v.Alias) == strings.ToUpper(value) {
					selected = append(selected, v)
				}
			case FieldLOCATION:
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
