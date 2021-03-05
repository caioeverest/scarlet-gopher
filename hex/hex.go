package hex

import (
	"reflect"
	"time"

	"github.com/caioeverest/scarlet-gopher/hex/strategy"
)

type Hex struct {
	Decade int
	Size   int
}

const (
	tag = "hex"
)

func Make(decade, size int) *Hex { return &Hex{Decade: decade, Size: size} }

func (h *Hex) Enter(objs ...interface{}) {
	for _, obj := range objs {
		v := reflect.ValueOf(obj).Elem()
		if v.Kind() == reflect.Struct {
			h.changeStructure(v)
		}
	}
}

func (h *Hex) changeStructure(v reflect.Value) {
	var (
		t   = v.Type()
		ctx = strategy.Context{strategy.BaseYearCtx: h.Decade}
		err error
	)

	for i := 0; i < v.NumField(); i++ {
		var (
			fieldV           = v.Field(i)
			fieldT           = t.Field(i)
			tagvalue, exists = fieldT.Tag.Lookup(tag)
			strg             strategy.Strategy
			newval           reflect.Value
		)

		if (!exists || fieldV.IsZero() || !fieldV.CanSet()) && fieldV.Kind() != reflect.Struct && fieldV.Kind() != reflect.Slice {
			continue
		}

		ctx[tagvalue] = fieldV.Interface()
		if strg, err = strategy.GetStrategy(ctx, tagvalue); err != nil && fieldV.Kind() != reflect.Struct && fieldV.Kind() != reflect.Slice {
			continue
		}

		switch fieldV.Kind() {
		case reflect.Slice:
			h.handleSlice(fieldV, strg)
		case reflect.Struct:
			h.handleStruct(fieldV, strg)
		case reflect.Ptr:
			h.handlePrt(fieldV, strg)
		default:
			newRawVal := strg.Change(fieldV.Interface())
			if newval, err = converter(fieldV.Interface(), newRawVal); err != nil {
				continue
			}
			fieldV.Set(newval)
		}

	}
}

func (h *Hex) handleTime(fieldV reflect.Value, strg strategy.Strategy) {
	var (
		newval reflect.Value
		err    error
	)

	newRawVal := strg.Change(fieldV.Interface())
	if newval, err = converter(fieldV.Interface(), newRawVal); err != nil {
		return
	}
	fieldV.Set(newval)
}

func (h *Hex) handleStruct(fieldV reflect.Value, strg strategy.Strategy) {
	switch fieldV.Interface().(type) {
	case time.Time:
		h.handleTime(fieldV, strg)
	default:
		h.changeStructure(fieldV)
	}
}

func (h *Hex) handleSlice(fieldV reflect.Value, _ strategy.Strategy) {
	for i := 0; i < fieldV.Len(); i++ {
		elem := fieldV.Index(i)
		if elem.Kind() == reflect.Struct {
			h.changeStructure(elem)
			return
		}
	}
}

func (h *Hex) handlePrt(fieldV reflect.Value, strg strategy.Strategy) {
	var (
		err    error
		newval reflect.Value
	)
	newRawVal := strg.Change(fieldV.Elem().Interface())
	if newval, err = converter(fieldV.Elem().Interface(), newRawVal); err != nil {
		return
	}
	initializer := reflect.New(fieldV.Type().Elem())
	initializer.Elem().Set(newval)
	fieldV.Set(initializer)
}
