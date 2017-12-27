package commands

import (
	"strings"
	"testing"
)

func TestWriteCommandBuildingWorks(t *testing.T) {
	cmd := NewWriteCommand(WR, 0, 1000, 20, "USER").ToString()

	if !strings.Contains(cmd, "WR") {
		t.Fatal("Command has invalid prefix")
	}
}
