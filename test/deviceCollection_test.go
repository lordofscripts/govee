/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * TESTS GoveeDevice and DeviceCollection
 *-----------------------------------------------------------------*/
package test

import (
	"fmt"
	"testing"

	"github.com/lordofscripts/govee"
)

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/
const (
	emptyMac string = "00:00:00:00:00:00:00:00"
	anyMac1  string = "AA:BB:CC:DD:EE:FF:00:11"
	anyMac2  string = "DD:EE:AA:DD:FF:EE:77:77"
	anyMac3  string = "DD:EE:AA:DD:FF:EE:CC:33"
)

var (
	validDevice1 = govee.GoveeDevice{"H6088", anyMac1, "device1", "garage"}
	validDevice2 = govee.GoveeDevice{"H6052", anyMac2, "device2", ""}
	validDevice3 = govee.GoveeDevice{"H6053", anyMac2, "device3", "bedroom"}

	invalidDevice1 = govee.GoveeDevice{"H6040", anyMac3, "", ""}
	invalidDevice2 = govee.GoveeDevice{"H6045", emptyMac, "", ""}
	invalidDevice3 = govee.GoveeDevice{"", anyMac3, "", ""}
)

/* ----------------------------------------------------------------
 *				M o d u l e   I n i t i a l i z a t i o n
 *-----------------------------------------------------------------*/
func init() {
	fmt.Println("Device/Collection Tests")
}

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

// must have Model, Alias/ID and a non-empty MAC that is not zeroes
func TestValidDevice(t *testing.T) {
	const Template string = "Should have been valid: %s"

	if !validDevice1.IsValid() {
		t.Errorf(Template, validDevice1.String())
	}

	if !validDevice2.IsValid() {
		t.Errorf(Template, validDevice2.String())
	}
}

// must have Model, Alias/ID and a non-empty MAC that is not zeroes to be valid
func TestInvalidDevice(t *testing.T) {
	const Template string = "Should have been invalid: %s"

	if invalidDevice1.IsValid() {
		t.Errorf(Template, invalidDevice1.String())
	}

	if invalidDevice2.IsValid() {
		t.Errorf(Template, invalidDevice2.String())
	}

	if invalidDevice3.IsValid() {
		t.Errorf(Template, invalidDevice3.String())
	}
}

func TestDeviceCollectionCount(t *testing.T) {
	var collection govee.DeviceCollection = govee.DeviceCollection{
		validDevice1,
		validDevice2,
		validDevice3, // same MAC as previous
		invalidDevice1,
		invalidDevice2,
		invalidDevice3,
	}

	exp := len(collection)
	got := collection.Count()
	if got != exp {
		t.Errorf("Unexpected collection Count() %d != %d", got, exp)
	}
}

func TestDeviceCollectionWhere(t *testing.T) {
	var collection govee.DeviceCollection = govee.DeviceCollection{
		validDevice1,
		validDevice2,
		validDevice3, // same MAC as previous
		invalidDevice1,
		invalidDevice2,
		invalidDevice3,
	}

	var filtered govee.DeviceCollection
	var seek string

	// there is only one H6052
	seek = "H6052"
	filtered = collection.Where(govee.FieldMODEL, seek)
	if len(filtered) != 1 {
		t.Errorf("There should have been only ONE by that Model %q", seek)
	} else {
		fmt.Println("\t[].Where(MODEL) OK")
	}

	// there are two with the same MAC (testing purposes only)
	seek = anyMac2
	filtered = collection.Where(govee.FieldMAC, seek)
	if len(filtered) != 2 {
		t.Errorf("There should have been only TWO by that MAC %q", seek)
	} else {
		fmt.Println("\t[].Where(MAC) OK")
	}

	// there is ONE with that Alias
	seek = "device3"
	filtered = collection.Where(govee.FieldALIAS, seek)
	if len(filtered) != 1 {
		t.Errorf("There should have been only ONE with ID %q", seek)
	} else {
		fmt.Println("\t[].Where(ID) OK")
	}

	// there is ONE with that Location
	seek = "bedroom"
	filtered = collection.Where(govee.FieldLOCATION, seek)
	if len(filtered) != 1 {
		t.Errorf("There should have been only ONE with Location %q", seek)
	} else {
		fmt.Println("\t[].Where(LOCATION) OK")
	}
}
