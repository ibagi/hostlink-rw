package shell

import (
	"fmt"
	"os"

	"github.com/rudhyd/hostlink-rw/hostlink"
)

func (s *Shell) prompt() string {
	var prompt string

	if s.isOpened {
		prompt = fmt.Sprintf("%s:%s> ", s.name, s.link.Address)
	} else {
		prompt = s.name + ":> "
	}

	return prompt
}

func (s *Shell) open(args []string) {
	err := s.validator.ValidateOpen(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s.link = hostlink.NewConnection(args[1])
	err = s.link.Open()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		s.isOpened = true
		fmt.Println("Connection opened successfully...")
	}
}

func (s *Shell) read(args []string) {
	cmd, err := s.validator.ValidateRead(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := s.link.Read(cmd)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
}

func (s *Shell) write(args []string) {
	cmd, err := s.validator.ValidateWrite(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ok, err := s.link.Write(cmd)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ok)
}

func (s *Shell) close(args []string) {
	if s.isOpened {
		s.link.Close()
		s.isOpened = false
		fmt.Println("Connection closed...")
	}
}

func (s *Shell) exit(args []string) {
	s.close(args)
	fmt.Println("Exiting...")
	os.Exit(0)
}
