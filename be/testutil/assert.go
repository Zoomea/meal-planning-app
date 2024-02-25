package testutil

import (
	"reflect"
	"testing"
)

func Assert[T any](t *testing.T, got, expect T) {
	t.Helper()
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf(`
-- expected --
%v
-- got --
%v`, expect, got)
	}
}
