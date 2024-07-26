package test_generator

import (
	"github.com/google/uuid"
	"math/rand"
	"reflect"
)

var CustomTypeHandlers = make(map[reflect.Type]func() reflect.Value)

func RegisterCustomTypeHandler(t reflect.Type, handler func() reflect.Value) {
	CustomTypeHandlers[t] = handler
}

func randomValue(t reflect.Type) reflect.Value {
	if handler, found := CustomTypeHandlers[t]; found {
		return handler()
	}

	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(uuid.NewString())
	case reflect.Int:
		return reflect.ValueOf(int(rand.Intn(100)))
	case reflect.Int8:
		return reflect.ValueOf(int8(rand.Intn(100)))
	case reflect.Int16:
		return reflect.ValueOf(int16(rand.Intn(100)))
	case reflect.Int32:
		return reflect.ValueOf(int32(rand.Intn(100)))
	case reflect.Int64:
		return reflect.ValueOf(rand.Int63())
	case reflect.Uint:
		return reflect.ValueOf(uint(rand.Intn(100)))
	case reflect.Uint8:
		return reflect.ValueOf(uint8(rand.Intn(100)))
	case reflect.Uint16:
		return reflect.ValueOf(uint16(rand.Intn(100)))
	case reflect.Uint32:
		return reflect.ValueOf(uint32(rand.Intn(100)))
	case reflect.Uint64:
		return reflect.ValueOf(rand.Uint64())
	case reflect.Float32:
		return reflect.ValueOf(rand.Float32())
	case reflect.Float64:
		return reflect.ValueOf(rand.Float64())
	case reflect.Bool:
		return reflect.ValueOf(rand.Intn(2) == 1)
	case reflect.Struct:
		return randomStruct(t)
	case reflect.Ptr:
		elem := randomValue(t.Elem())
		ptr := reflect.New(t.Elem())
		ptr.Elem().Set(elem)
		return ptr
	default:
		return reflect.Zero(t)
	}
}

func randomStruct(t reflect.Type) reflect.Value {
	v := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if field.CanSet() {
			field.Set(randomValue(field.Type()))
		}
	}
	return v
}

func PopulateStruct(s interface{}) {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if field.CanSet() {
			field.Set(randomValue(field.Type()))
		}
	}
}

func ConstructExpected[T any, R any](args T, modifyFunc func(*R)) R {
	var expected R
	populateFromArgs(&expected, args)
	if modifyFunc != nil {
		modifyFunc(&expected)
	}
	return expected
}

func populateFromArgs[R any, T any](expected *R, args T) {
	eVal := reflect.ValueOf(expected).Elem()
	aVal := reflect.ValueOf(args)

	for i := 0; i < eVal.NumField(); i++ {
		fieldName := eVal.Type().Field(i).Name
		aField := aVal.FieldByName(fieldName)
		if aField.IsValid() {
			eField := eVal.FieldByName(fieldName)
			if eField.CanSet() {
				eField.Set(aField)
			}
		}
	}
}
