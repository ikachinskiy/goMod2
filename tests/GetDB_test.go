package sngDB

import (
	sngDB "smartnet.ru/sngDB"
	"testing"
)

func TestGetDB(t *testing.T) {
	want := "GetDB - работает"
	if got := sngDB.GetDB(); got != want {
		t.Errorf("GetDB() = %q, а надо %q", got, want)
	}
}
