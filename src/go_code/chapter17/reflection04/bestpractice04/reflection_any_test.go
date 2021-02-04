package test

import (
	"reflect"
	"testing"
)

type User struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *User
		sv    reflect.Value
	)

	model = &User{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String()) // address
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String()) // value

	sv.FieldByName("UserId").SetString("0000000")
	sv.FieldByName("Name").SetString("admin")

	t.Log("model", model)

}
