package id

import "testing"

func TestGenHashID(t *testing.T) {
	id := 999999
	hasid := NumberToHashID(int64(id), "123456789")
	number := HashIDToNumber(hasid, "123456789")
	t.Log(hasid, number)
	if number != int64(id) {
		t.Errorf("hashid error")
	}
}
