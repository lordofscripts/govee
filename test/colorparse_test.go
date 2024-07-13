/* -----------------------------------------------------------------
 *				C o r a l y s   T e c h n o l o g i e s
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 *						U n i t   T e s t
 *-----------------------------------------------------------------*/
package test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/lordofscripts/govee/util"
)

func init() {
	fmt.Println("Color Parse Test")
}

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/
const (
	RED_SHORT    uint8 = 0x4
	RED_NORMAL   uint8 = 0x35
	GREEN_SHORT  uint8 = 0x7
	GREEN_NORMAL uint8 = 0x69
	BLUE_SHORT   uint8 = 0xa
	BLUE_NORMAL  uint8 = 0x9b
	ALPHA_SHORT  uint8 = 0x8
	ALPHA_NORMAL uint8 = 0x8e

	TAB = "\t"
)

/* ----------------------------------------------------------------
 *				U n i t  T e s t   F u n c t i o n s
 *-----------------------------------------------------------------*/
func Test_ColorParseShort(t *testing.T) {
	col := toShortColor(RED_SHORT, GREEN_SHORT, BLUE_SHORT, ALPHA_SHORT)
	exp := toColorString(col, true)
	var got string

	// Without Alpha
	if gotCol, err := util.ParseHexColor(exp); err != nil {
		t.Errorf("Oops! %s", err.Error())
	} else {
		got = toColorString(gotCol, true)

		if exp != got {
			t.Errorf("Failed expect: %q != Got %q", exp, got)
		} else {
			fmt.Printf("\tParseHexColor(%q)\n", exp)
		}
	}

	// With Alpha
	exp = toColorStringAlpha(col, true)
	if gotCol, err := util.ParseHexColor(exp); err != nil {
		t.Errorf("Oops! %s", err.Error())
	} else {
		got = toColorStringAlpha(gotCol, true)
		if exp != got {
			t.Errorf("Failed expect: %q != Got %q", exp, got)
		} else {
			fmt.Printf("\tParseHexColor(%q)\n", exp)
		}
	}
}

func Test_ColorParseStandard(t *testing.T) {
	col := toShortColor(RED_NORMAL, GREEN_NORMAL, BLUE_NORMAL, ALPHA_NORMAL)
	exp := toColorString(col, false)
	var got string

	// Without Alpha
	if gotCol, err := util.ParseHexColor(exp); err != nil {
		t.Errorf("Oops! %s", err.Error())
	} else {
		got = toColorString(gotCol, false)
		if exp != got {
			t.Errorf("Failed expect: %q != Got %q", exp, got)
		} else {
			fmt.Printf("\tParseHexColor(%q)\n", exp)
		}
	}

	// With Alpha
	exp = toColorStringAlpha(col, false)
	if gotCol, err := util.ParseHexColor(exp); err != nil {
		t.Errorf("Oops! %s", err.Error())
	} else {
		got = toColorStringAlpha(gotCol, false)
		if exp != got {
			t.Errorf("Failed expect: %q != Got %q", exp, got)
		} else {
			fmt.Printf("\tParseHexColor(%q)\n", exp)
		}
	}
}

func Test_ColorInvalidFormat(t *testing.T) {
	var err error
	var invalid string

	// missing leading #
	invalid = "2e9"
	_, err = util.ParseHexColor(invalid)
	if err == nil {
		t.Errorf("Expected error due to missing leading #")
	} else {
		fmt.Println(TAB, err.Error())
	}

	// has a non-hex digit
	invalid = "#aa3g5f"
	_, err = util.ParseHexColor(invalid)
	if err == nil {
		t.Errorf("Expected error due to non-hex")
	} else {
		fmt.Println(TAB, err.Error())
	}
}

/* ----------------------------------------------------------------
 *					H e l p e r   F u n c t i o n s
 *-----------------------------------------------------------------*/

func toShortColor(r, g, b, a uint8) color.RGBA {
	r = r & 0x0f
	r = r | (r << 4)
	g = g & 0x0f
	g = g | (g << 4)
	b = b & 0x0f
	b = b | (b << 4)
	a = a & 0x0f
	a = a | (a << 4)
	return color.RGBA{r, g, b, a}
}

func toColor(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

// color to String without Alpha: #RGB or #RRGGBB
func toColorString(c color.RGBA, short bool) string {
	var template = "#%02X%02X%02X"
	if short {
		template = "#%01X%01X%01X"
	}
	return fmt.Sprintf(template, c.R, c.G, c.B)
}

// color to String with Alpha: #RGBA or #RRGGBBAA
func toColorStringAlpha(c color.RGBA, short bool) string {
	var template = "#%02X%02X%02X%02X"
	if short {
		template = "#%01X%01X%01X%01X"
	}
	return fmt.Sprintf(template, c.R, c.G, c.B, c.A)
}
