/* -----------------------------------------------------------------
 *				  https://allmylinks.com/lordofscripts
 *				  Copyright (C)2023 Lord of Scripts(tm)
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package util

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

// Boolean OR: Any(b0,b1,b2...bn)
func Any(params ...bool) bool {
	var value uint = 0
	for i, v := range params {
		var operand uint = 0
		if v {
			operand = 1
		}
		value = value | (operand << i)
	}
	return value != 0
}

// Boolean AND: None(b0,b1,b2...bn)
func None(params ...bool) bool {
	var value uint = 0
	for i, v := range params {
		var operand uint = 0
		if v {
			operand = 1
		}
		value = value | (operand << i)
	}
	return value == 0
}

// Just one set
func One(params ...bool) bool {
	cnt := 0
	for _, v := range params {
		if v {
			cnt += 1
		}
	}
	return cnt == 1
}

// Boolean exclusive OR (XOR)
func ExclusiveOr(params ...bool) bool {
	cnt := 0
	for _, v := range params {
		if v {
			cnt += 1
		}
	}
	return !((cnt % 2) == 0)
}
