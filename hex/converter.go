package hex

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func converter(kind interface{}, rawvalue interface{}) (out reflect.Value, err error) {
	v := fmt.Sprintf("%v", rawvalue)

	switch kind.(type) {
	case bool:
		var tmp bool
		tmp, err = strconv.ParseBool(v)
		out = reflect.ValueOf(tmp)
	case int:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(tmp)
	case int8:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int8(tmp))
	case int16:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int16(tmp))
	case int32:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int32(tmp))
	case int64:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int64(tmp))
	case uint:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint(tmp))
	case uint8:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint8(tmp))
	case uint16:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint16(tmp))
	case uint32:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint32(tmp))
	case uint64:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint64(tmp))
	case float32:
		var tmp float64
		tmp, err = strconv.ParseFloat(v, 32)
		out = reflect.ValueOf(float32(tmp))
	case float64:
		var tmp float64
		tmp, err = strconv.ParseFloat(v, 64)
		out = reflect.ValueOf(tmp)
	case time.Time:
		tmp, ok := rawvalue.(time.Time)
		if !ok {
			err = errors.New(fmt.Sprintf("type %s is unexpected, it should be type time.Time", reflect.TypeOf(rawvalue)))
		}
		out = reflect.ValueOf(tmp)
	default:
		out = reflect.ValueOf(v)
	}
	return
}
