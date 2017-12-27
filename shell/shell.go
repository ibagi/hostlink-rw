package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rudhyd/hostlink-rw/commands"
	"github.com/rudhyd/hostlink-rw/hostlink"
)

// Shell interactive hostlink shell
type Shell struct {
	name      string
	isOpened  bool
	reader    *bufio.Reader
	link      *hostlink.Hostlink
	validator *ShellValidator
}

// NewShell creates new interactive hostlink shell
func NewShell(name string) *Shell {
	return &Shell{
		name:      name,
		reader:    bufio.NewReader(os.Stdin),
		validator: &ShellValidator{},
	}
}

// Execute starts interactive shell execution
func (s *Shell) Execute() {
	for {
		fmt.Print(s.prompt())
		cmd, err := s.reader.ReadString('\n')

		if err != nil {
			panic(err.Error())
		}

		args := strings.Fields(cmd)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "open":
			s.open(args)
		case "close":
			s.close(args)
		case "exit":
			s.exit(args)
		default:
			switch {
			case commands.IsValidReadCommand(strings.ToUpper(args[0])):
				s.read(args)
			case commands.IsValidWriteCommand(strings.ToUpper(args[0])):
				s.write(args)
			default:
				fmt.Println("Unknown command")
			}
		}
	}
}
