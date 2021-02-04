package test

import (
	"reflect"
	"testing"
)

type User struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *User
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf", st.Kind().String()) // ptr

	st = st.Elem()
	t.Log("reflect.ValueOf.Elem", st.Kind().String()) // struct

	elem = reflect.New(st) // empty here point to address which point to a object
	t.Log("reflect.New", elem.Kind().String()) // ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) // ptr

	model = elem.Interface().(*User) // model point to the address which point to a object

	elem = elem.Elem() // elem point to actual object

	elem.FieldByName("UserId").SetString("0000000")
	elem.FieldByName("Name").SetString("admin")

	t.Log("model model.Name", model, model.Name)

}
