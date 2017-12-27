package commands

import (
	"fmt"
)

// ReadCommandType enumeration of Hostlink read command types
type ReadCommandType string

const (
	RR ReadCommandType = "RR"
	RL                 = "RL"
	RH                 = "RH"
	RD                 = "RD"
	RJ                 = "RJ"
	RC                 = "RC"
	RG                 = "RG"
)

// IsValidReadCommand checks if the given string is valid ReadCommandType
func IsValidReadCommand(t string) bool {
	return string(RC) == t ||
		string(RD) == t ||
		string(RG) == t ||
		string(RH) == t ||
		string(RJ) == t ||
		string(RL) == t ||
		string(RR) == t
}

// ReadCommand holds Hostlink read command data
type ReadCommand struct {
	CommandBase
	Type ReadCommandType
}

// NewReadCommand creates new Hostlink ReadCommand
func NewReadCommand(t ReadCommandType, module int, address int, length int) *ReadCommand {
	return &ReadCommand{
		CommandBase: CommandBase{module, address, length},
		Type:        t,
	}
}

// ToString creates to command string that can be send to the Hostlink device
func (c *ReadCommand) ToString() string {
	result := fmt.Sprintf(`@%02d%s%04d%04d`, c.Module, c.Type, c.Address, c.Length)
	result += c.CalculateFCS(result)
	result += terminator
	return result
}
