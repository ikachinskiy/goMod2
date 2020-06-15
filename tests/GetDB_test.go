package goMod2

import (
	goMod2 "github.com/ikachinskiy/goMod2"
	"testing"
)

func TestGetDB(t *testing.T) {
	want := "GetDB - работает"
	if got := goMod2.GetDB(); got != want {
		t.Errorf("GetDB() = %q, а надо %q", got, want)
	}
}

func TestSetDB(t *testing.T) {
	dbname := "smartnet"
	if got := goMod2.SetDB(dbname); got != dbname {
		t.Errorf("SetDB(%q) = %q, а надо %q", dbname, got, dbname)
	}
}
