package main

import (
	"github.com/rudhyd/hostlink-rw/shell"
)

func main() {
	hShell := shell.NewShell("hostlink-rw")
	hShell.Execute()
}
