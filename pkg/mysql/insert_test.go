package mysql

import "testing"

func TestGetPlaceholderN(t *testing.T) {
	s := GetPlaceholderN(3)
	if s != "?,?,?" {
		t.Error(s)
	}

	s = GetPlaceholderN(1)
	if s != "?" {
		t.Error(s)
	}
}
