package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// Dump - format object hierarchically
func Dump(s interface{}, lvl int) string {
	res := ""
	c := [...]rune{
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
	}

	rtyp := reflect.TypeOf(s)
	rval := reflect.ValueOf(s)
	ltab := string(c[0:lvl])

	for i := 0; i < rtyp.NumField(); i++ {
		fname := rtyp.Field(i).Name
		//		ftag := rtyp.Field(i).Tag

		//		res += fmt.Sprintf("%s%s %s = ", ltab, fname, rval.Field(i).Type())
		res += fmt.Sprintf("%s%s = ", ltab, fname)
		switch rval.Field(i).Kind() {
		case reflect.Struct:
			res += "{\n" + Dump(rval.Field(i).Interface(), lvl+2) + ltab + "}"
		case reflect.Int, reflect.Int32, reflect.Int64:
			res += strconv.FormatInt(rval.Field(i).Int(), 10)
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			res += strconv.FormatUint(rval.Field(i).Uint(), 10)
		case reflect.Bool:
			res += strconv.FormatBool(rval.Field(i).Bool())
		case reflect.Slice:
			res += fmt.Sprint(rval.Field(i).Slice(0, rval.Field(i).Len()))
		default:
			res += rval.Field(i).String()
		}
		res += "\n"
	}

	return res
}

func BoolToString(v bool, s1, s0 string) string {
	if v {
		return s1
	} else {
		return s0
	}
}

type RuneChar struct {
	RuneUp rune
	RuneLo rune
	CharUp byte
	CharLo byte
}

var RuneChars = []RuneChar{
	{'А', 'а', 0xC0, 0xE0},
	{'Б', 'б', 0xC1, 0xE1},
	{'В', 'в', 0xC2, 0xE2},
	{'Г', 'г', 0xC3, 0xE3},
	{'Д', 'д', 0xC4, 0xE4},
	{'Е', 'е', 0xC5, 0xE5},
	{'Ж', 'ж', 0xC6, 0xE6},
	{'З', 'з', 0xC7, 0xE7},
	{'И', 'и', 0xC8, 0xE8},
	{'Й', 'й', 0xC9, 0xE9},
	{'К', 'к', 0xCA, 0xEA},
	{'Л', 'л', 0xCB, 0xEB},
	{'М', 'м', 0xCC, 0xEC},
	{'Н', 'н', 0xCD, 0xED},
	{'О', 'о', 0xCE, 0xEE},
	{'П', 'п', 0xCF, 0xEF},
	{'Р', 'р', 0xD0, 0xF0},
	{'С', 'с', 0xD1, 0xF1},
	{'Т', 'т', 0xD2, 0xF2},
	{'У', 'у', 0xD3, 0xF3},
	{'Ф', 'ф', 0xD4, 0xF4},
	{'Х', 'х', 0xD5, 0xF5},
	{'Ц', 'ц', 0xD6, 0xF6},
	{'Ч', 'ч', 0xD7, 0xF7},
	{'Ш', 'ш', 0xD8, 0xF8},
	{'Щ', 'щ', 0xD9, 0xF9},
	{'Ъ', 'ъ', 0xDA, 0xFA},
	{'Ы', 'ы', 0xDB, 0xFB},
	{'Ь', 'ь', 0xDC, 0xFC},
	{'Э', 'э', 0xDD, 0xFD},
	{'Ю', 'ю', 0xDE, 0xFE},
	{'Я', 'я', 0xDF, 0xFF},
}

func RuneToChar(r rune) byte {
	for _, rc := range RuneChars {
		switch r {
		case rc.RuneUp, rc.RuneLo:
			return rc.CharUp
		}
	}
	return byte(r)
}

func CharToRune(c byte) rune {
	for _, rc := range RuneChars {
		switch c {
		case rc.CharUp:
			return rc.RuneUp
		case rc.CharLo:
			return rc.RuneLo
		}
	}
	return rune(c)
}
