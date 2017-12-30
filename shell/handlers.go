package shell

import (
	"fmt"
	"os"

	"github.com/rudhyd/hostlink-rw/hostlink"
)

func (s *Shell) usage() {
	fmt.Printf("%s commands: \n", s.name)
	fmt.Println("- open: 	[host], open new connection")
	fmt.Println("- r*: 		[type] [module] [address] [length], send read-command when connection is opened")
	fmt.Println("- w*: 		[type] [module] [address] [length] [data], send write-command when connection is opened")
	fmt.Println("- close: 	close current connection")
	fmt.Println("- exit: 	exit current program")
}

func (s *Shell) checkStatus() bool {
	if !s.isOpened {
		fmt.Println("first you need to open a connection with the [open] command!")
		return false
	}

	return true
}

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

	s.close(args)
	s.link = hostlink.NewConnection(args[1])
	err = s.link.Open()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		s.isOpened = true
		fmt.Println("connection opened successfully...")
	}
}

func (s *Shell) read(args []string) {
	if !s.checkStatus() {
		return
	}

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
	if !s.checkStatus() {
		return
	}

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
		fmt.Println("connection closed...")
	}
}

func (s *Shell) exit(args []string) {
	s.close(args)
	fmt.Println("exiting...")
	os.Exit(0)
}
