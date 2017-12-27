package commands

import (
	"fmt"
)

// WriteCommandType enumeration of Hostlink write command types
type WriteCommandType string

const (
	WR WriteCommandType = "WR"
	WL                  = "WL"
	WH                  = "WH"
	WD                  = "WD"
	WJ                  = "WJ"
	WC                  = "WC"
	WG                  = "WG"
)

// IsValidWriteCommand checks if the given string is valid WriteCommandType
func IsValidWriteCommand(t string) bool {
	return string(WC) == t ||
		string(WD) == t ||
		string(WG) == t ||
		string(WH) == t ||
		string(WJ) == t ||
		string(WL) == t ||
		string(WR) == t
}

// WriteCommand holds Hostlink write command data
type WriteCommand struct {
	CommandBase
	Type WriteCommandType
	Data string
}

// NewWriteCommand creates a new Hostlink WriteCommand
func NewWriteCommand(t WriteCommandType, module int, address int, length int, data string) *WriteCommand {
	return &WriteCommand{
		CommandBase: CommandBase{module, address, length},
		Type:        t,
		Data:        data,
	}
}

// ToString creates to command string that can be send to the Hostlink device
func (c *WriteCommand) ToString() string {
	data := c.createData(c.Data, c.Length)
	result := fmt.Sprintf(`@%02d%s%04d%s`, c.Module, c.Type, c.Address, data)
	result += c.CalculateFCS(result)
	result += terminator
	return result
}
