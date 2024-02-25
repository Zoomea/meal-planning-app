package fsdatabase

import "reflect"

// If v is a pointer to a struct containing a field "ID" of type int64,
// it will set its value to that of the given "id" parameters
func setIDIfExists[T any](v *T, id int64) {
	s := reflect.ValueOf(v)
	s = s.Elem()
	if s.Type().Kind() != reflect.Struct {
		return
	}

	idField := s.FieldByName("ID")
	if !idField.IsValid() {
		return
	}

	if !idField.CanSet() {
		return
	}

	if idField.Type().String() != "int64" {
		return
	}

	idField.SetInt(id)
}
