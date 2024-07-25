package test_generator

import (
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
		return reflect.ValueOf(RandomString(10))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(RandomInt(1, 100))
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(RandomFloat64(1.0, 100.0))
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
