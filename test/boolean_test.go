/* -----------------------------------------------------------------
 *				  https://allmylinks.com/lordofscripts
 *				  Copyright (C)2023 Lord of Scripts(tm)
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package test

import (
	"fmt"
	"github.com/lordofscripts/govee/util"
	"testing"
)

func init() {
	fmt.Println("Boolean Test")
}

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/
type testCase struct {
	display string
	values  []bool
	expect  bool
}

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/
const (
	// test bit patterns (readable)
	PAT_O = "000"
	PAT_1 = "001"
	PAT_2 = "010"
	PAT_3 = "011"
	PAT_4 = "100"
	PAT_5 = "101"
	PAT_6 = "110"
	PAT_7 = "111"
)

var (
	// test bit pattern values
	value0 = []bool{false, false, false}
	value1 = []bool{false, false, true}
	value2 = []bool{false, true, false}
	value3 = []bool{false, true, true}
	value4 = []bool{true, false, false}
	value5 = []bool{true, false, true}
	value6 = []bool{true, true, false}
	value7 = []bool{true, true, true}
)

/* ----------------------------------------------------------------
 *							C o n s t r u c t o r s
 *-----------------------------------------------------------------*/
func newCase(disp string, exp bool, vals []bool) testCase {
	return testCase{
		display: disp,
		values:  vals,
		expect:  exp,
	}
}

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/
func TestOne(t *testing.T) {
	testCases := []testCase{
		newCase(PAT_O, false, value0),
		newCase(PAT_1, true, value1),
		newCase(PAT_2, true, value2),
		newCase(PAT_3, false, value3),
		newCase(PAT_4, true, value4),
		newCase(PAT_5, false, value5),
		newCase(PAT_6, false, value6),
		newCase(PAT_7, false, value7),
	}

	for i, tCase := range testCases {
		// https://stackoverflow.com/questions/23723955/how-can-i-pass-a-slice-as-a-variadic-input
		result := util.One(tCase.values...)
		if result != tCase.expect {
			t.Errorf("#%d FAIL One-of %s needs:%t got %t\n", i, tCase.display, tCase.expect, result)
		} else {
			fmt.Printf("#%d OK One-of %s = %t\n", i, tCase.display, result)
		}
	}
}

func TestAny(t *testing.T) {
	testCases := []testCase{
		newCase(PAT_O, false, value0),
		newCase(PAT_1, true, value1),
		newCase(PAT_2, true, value2),
		newCase(PAT_3, true, value3),
		newCase(PAT_4, true, value4),
		newCase(PAT_5, true, value5),
		newCase(PAT_6, true, value6),
		newCase(PAT_7, true, value7),
	}

	for i, tCase := range testCases {
		result := util.Any(tCase.values...)
		if result != tCase.expect {
			t.Errorf("#%d FAIL Any-of %s needs:%t got %t\n", i, tCase.display, tCase.expect, result)
		} else {
			fmt.Printf("#%d OK Any-of %s = %t\n", i, tCase.display, result)
		}
	}
}

func TestNone(t *testing.T) {
	testCases := []testCase{
		newCase(PAT_O, true, value0),
		newCase(PAT_1, false, value1),
		newCase(PAT_2, false, value2),
		newCase(PAT_3, false, value3),
		newCase(PAT_4, false, value4),
		newCase(PAT_5, false, value5),
		newCase(PAT_6, false, value6),
		newCase(PAT_7, false, value7),
	}

	for i, tCase := range testCases {
		result := util.None(tCase.values...)
		if result != tCase.expect {
			t.Errorf("#%d FAIL None-of %s needs:%t got %t\n", i, tCase.display, tCase.expect, result)
		} else {
			fmt.Printf("#%d OK None-of %s = %t\n", i, tCase.display, result)
		}
	}
}

func TestExclusiveOr(t *testing.T) {
	testCases := []testCase{
		newCase(PAT_O, false, value0),
		newCase(PAT_1, true, value1),
		newCase(PAT_2, true, value2),
		newCase(PAT_3, false, value3),
		newCase(PAT_4, true, value4),
		newCase(PAT_5, false, value5),
		newCase(PAT_6, false, value6),
		newCase(PAT_7, true, value7),
	}

	for i, tCase := range testCases {
		result := util.ExclusiveOr(tCase.values...)
		//checkmark := "\u2713"
		if result != tCase.expect {
			t.Errorf("#%d FAIL ExclusiveOr-of %s needs:%t got %t\n", i, tCase.display, tCase.expect, result)
		} else {
			fmt.Printf("#%d OK ExclusiveOr-of %s = %t\n", i, tCase.display, result)
		}
	}
}
