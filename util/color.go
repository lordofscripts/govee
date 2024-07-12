/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2024 Dídimo Grimaldo T.
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *
 *-----------------------------------------------------------------*/
package util

import (
	"errors"
	"image/color"

	veex "github.com/loxhill/go-vee"
)

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/

var ErrInvalidFormat = errors.New("invalid format")

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

func ParseHexColorGovee(s string) (veex.Color, error) {
	if color, err := ParseHexColor(s); err == nil {
		return veex.Color{int(color.R), int(color.G), int(color.B)}, nil
	} else {
		return veex.Color{}, err
	}
}

// Parses a hexadecimal color string and returns a Go image/color.
// The string must begin with '#' and be followed by 3,4,6 or 8 hex digits
// representing a Red/Green/Blue triad with optional Opacity/Transparency.
// Therefore it can be #rgb #rgba #rrggbbaa or #rrggbb
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, ErrInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = ErrInvalidFormat
		return 0
	}

	switch len(s) {
	case 9: // #aabbccee (#RRGGBBAA)
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		c.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7: // #aabbcc (#RRGGBB)
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 5: // #123e (#RGBA)
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
		c.A = hexToByte(s[4]) * 17
	case 4: // #123 (#RGB)
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = ErrInvalidFormat

	}
	return
}
