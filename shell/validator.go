package shell

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rudhyd/hostlink-rw/commands"
)

var (
	ErrOpenCommand         error = errors.New("open: One argument [address] expected")
	ErrReadCommandArgs     error = errors.New("read: Four argument expected [type] [module] [address] [length]")
	ErrReadCommandModule   error = errors.New("read: param: [module], expected integer")
	ErrReadCommandAddress  error = errors.New("read: param: [address], expected integer")
	ErrReadCommandLength   error = errors.New("read: param: [length], expected integer")
	ErrWriteCommandArgs    error = errors.New("write: Five argument expected [type] [module] [address] [length] [data]")
	ErrWriteCommandModule  error = errors.New("write: param: [module], expected integer")
	ErrWriteCommandAddress error = errors.New("write: param: [address], expected integer")
	ErrWriteCommandLength  error = errors.New("write: param: [length], expected integer")
)

type ShellValidator struct{}

func (v *ShellValidator) ValidateOpen(args []string) error {
	if len(args) != 2 {
		return ErrOpenCommand
	}

	return nil
}

func (v *ShellValidator) ValidateRead(args []string) (*commands.ReadCommand, error) {
	if len(args) != 4 {
		return nil, ErrReadCommandArgs
	}

	cmdType := commands.ReadCommandType(strings.ToUpper(args[0]))

	module, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, ErrReadCommandModule
	}

	address, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, ErrReadCommandAddress
	}

	length, err := strconv.Atoi(args[3])
	if err != nil {
		return nil, ErrReadCommandLength
	}

	return commands.NewReadCommand(cmdType, module, address, length), nil
}

func (v *ShellValidator) ValidateWrite(args []string) (*commands.WriteCommand, error) {
	if len(args) != 5 {
		return nil, ErrWriteCommandArgs
	}

	cmdType := commands.WriteCommandType(strings.ToUpper(args[0]))

	module, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, ErrReadCommandModule
	}

	address, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, ErrReadCommandAddress
	}

	length, err := strconv.Atoi(args[3])
	if err != nil {
		return nil, ErrReadCommandLength
	}

	data := args[4]

	return commands.NewWriteCommand(cmdType, module, address, length, data), nil
}
