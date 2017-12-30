package hostlink

import (
	"fmt"
	"testing"

	"github.com/rudhyd/hostlink-rw/commands"
)

func TestOpenWillFail(t *testing.T) {
	link := NewConnection("localhost:8080")
	err := link.Open()

	if err == nil {
		t.Fatal("Fail is expected")
	}
}

func TestRead(t *testing.T) {
	link := NewConnection("10.150.10.5:4001")
	err := link.Open()

	if err != nil {
		panic(err.Error())
	}

	cmd := commands.NewReadCommand(commands.RD, 0, 1280, 20)
	data, err := link.Read(cmd)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(data)
}
