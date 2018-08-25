package utils

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// Imported from: https://github.com/dbohdan/remarshal/blob/10e552b63d3f10231269c8151ee86ddf9f8ff317/remarshal.go#L39

// ConvertMapsToStringMaps recursively converts values of type
// map[interface{}]interface{} contained in item to map[string]interface{}. This
// is needed before the encoders for TOML and JSON can accept data returned by
// the YAML decoder.
func ConvertMapsToStringMaps(item interface{}) (res interface{}, err error) {
	switch item.(type) {
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, v := range item.(map[interface{}]interface{}) {
			res[k.(string)], err = ConvertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = ConvertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	default:
		return item, nil
	}
}

// toMap converts a struct defined by `in` to a map[string]interface{}. It only
// extract data that is defined by the given tag.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only struct is allowd got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil

}

// toUint tries to convert the given to uint type
func ToUint(x interface{}) uint {
	switch i := x.(type) {
	case float64:
		return uint(i)
	case uint:
		return i
	case int:
		return uint(i)
	case string:
		s, err := strconv.Atoi(i)
		if err != nil {
			log.Println("cannot convert", i)
			return 0
		}

		return uint(s)
	case int64:
		return uint(i)
	default:
		return 0
	}
}
