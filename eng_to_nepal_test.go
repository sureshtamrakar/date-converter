package convert

import "testing"

func TestEngToNep(t *testing.T) {
	//expect := 10
	err, val := EngToNep(2022, 02, 23)
	if err != nil || val["date"] != "11" {
		t.Errorf("Something wrong")
	}
}
