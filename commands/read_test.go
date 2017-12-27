package commands

import (
	"strings"
	"testing"
)

func TestReadCommandBuildingWorks(t *testing.T) {
	cmd := NewReadCommand(WD, 0, 1000, 20).ToString()

	if !strings.Contains(cmd, "WD") {
		t.Fatal("Command has invalid prefix")
	}
}
