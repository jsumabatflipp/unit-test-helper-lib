package test_generator

import (
	"reflect"
	"testing"
)

func TestPopulateStruct(t *testing.T) {
	type MyStruct struct {
		Name  string
		Value int
	}

	var s MyStruct
	PopulateStruct(&s)
	if s.Name == "" || s.Value == 0 {
		t.Errorf("Expected struct to be populated, got %+v", s)
	}
}

func TestRandomString(t *testing.T) {
	str := RandomString(10)
	if len(str) != 10 {
		t.Errorf("Expected string of length 10, got %s", str)
	}
}

func TestRandomUUID(t *testing.T) {
	uuid := RandomUUID()
	if len(uuid) != 36 {
		t.Errorf("Expected UUID of length 36, got %s", uuid)
	}
}

func TestCustomTypeHandler(t *testing.T) {
	RegisterCustomTypeHandler(reflect.TypeOf(""), func() reflect.Value {
		return reflect.ValueOf("custom_string")
	})

	var s struct {
		CustomField string
	}

	PopulateStruct(&s)
	if s.CustomField != "custom_string" {
		t.Errorf("Expected custom_string, got %s", s.CustomField)
	}
}
