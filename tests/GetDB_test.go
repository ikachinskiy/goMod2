package goMod2

import (
	goMod2 "smartnet.ru/goMod2"
	"testing"
)

func TestGetDB(t *testing.T) {
	want := "GetDB - работает"
	if got := goMod2.GetDB(); got != want {
		t.Errorf("GetDB() = %q, а надо %q", got, want)
	}
}
