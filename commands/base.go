package commands

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

const (
	terminator string = "*\r"
	wordSize   int    = 4
)

// CommandBase base structure for Hostlink commands
type CommandBase struct {
	Module, Address, Length int
}

// CalculateFCS calculates command checksum
func (c *CommandBase) CalculateFCS(cmd string) string {
	checksum := int64(64)

	for _, c := range cmd[1:] {
		checksum = checksum ^ int64(c)
	}

	return strconv.FormatInt(checksum, 16)
}

// CheckFCS checks if the calculated checksum of the given command equals with current FCS
func (c *CommandBase) CheckFCS(cmd string) bool {
	expected := cmd[len(cmd)-3 : len(cmd)-1]
	fcs := cmd[:len(cmd)-3]

	return expected == c.CalculateFCS(fcs)
}

func (c *CommandBase) createData(s string, length int) string {
	var result string

	if value, err := strconv.Atoi(s); err == nil {
		result = fmt.Sprintf("%x", value)
	} else {
		result = hex.EncodeToString([]byte(s))
	}

	for len(result) < length*wordSize {
		result += "0"
	}

	return result
}
