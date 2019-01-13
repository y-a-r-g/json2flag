package json2flag

import (
	"flag"
	"fmt"
	"reflect"
	"time"
)

const (
	structDelimiter = "."
)

var (
	durationType = reflect.TypeOf(time.Duration(0))
)

func Flag(s interface{}, usage map[string]string) {
	FlagPrefixed(s, usage, "")
}

func FlagPrefixed(s interface{}, usage map[string]string, prefix string) {
	elem := reflect.ValueOf(s).Elem()
	t := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		f := elem.Field(i)
		name := t.Field(i).Name
		flagName := name
		if prefix != "" {
			flagName = fmt.Sprintf("%s%s%s", prefix, structDelimiter, name)
		}
		fieldType := t.Field(i).Type

		if fieldType == durationType {
			flag.DurationVar(f.Addr().Interface().(*time.Duration), flagName, f.Interface().(time.Duration), usage[name])
		} else {
			switch fieldType.Kind() {
			case reflect.String:
				flag.StringVar(f.Addr().Interface().(*string), flagName, f.Interface().(string), usage[name])
				break
			case reflect.Bool:
				flag.BoolVar(f.Addr().Interface().(*bool), flagName, f.Interface().(bool), usage[name])
				break
			case reflect.Int:
				flag.IntVar(f.Addr().Interface().(*int), flagName, f.Interface().(int), usage[name])
				break
			case reflect.Uint:
				flag.UintVar(f.Addr().Interface().(*uint), flagName, f.Interface().(uint), usage[name])
				break
			case reflect.Int64:
				flag.Int64Var(f.Addr().Interface().(*int64), flagName, f.Interface().(int64), usage[name])
				break
			case reflect.Uint64:
				flag.Uint64Var(f.Addr().Interface().(*uint64), flagName, f.Interface().(uint64), usage[name])
				break
			case reflect.Float64:
				flag.Float64Var(f.Addr().Interface().(*float64), flagName, f.Interface().(float64), usage[name])
				break
			default:
				flag.Var(f.Addr().Interface().(flag.Value), flagName, usage[name])
			}
		}
	}
}
