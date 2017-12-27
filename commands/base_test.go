package commands

import "testing"

func TestCalculateFCS(t *testing.T) {
	expectation := "55"

	cmd := NewReadCommand(WD, 0, 1000, 20)
	fcs := cmd.CalculateFCS("@00RD100020")

	if fcs != expectation {
		t.Fatal("FCS calculation failed")
	}
}

func TestCheckFCS(t *testing.T) {
	cmdString := `@00RD10002055*`

	cmd := NewReadCommand(WD, 0, 0, 0)
	result := cmd.CheckFCS(cmdString)

	if !result {
		t.Fatal("Check FCS failed")
	}
}
