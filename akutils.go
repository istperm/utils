package akutils

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
