package percentagerounds

import "testing"

func TestGetPercentageRounds(t *testing.T) {
	for i := 0; i <= 100; i++ {
		out := GetPercentageRounds(i)
		t.Log(out)
	}
}

func TestGetPercentageRoundsSubstring(t *testing.T) {
	for i := 0; i <= 100;i++{
		out:=GetPercentageRoundsSubstring(i)
		t.Log(out)
	}
}
