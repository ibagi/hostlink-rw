package hostlink

import (
	"bufio"
	"errors"
	"net"

	"github.com/rudhyd/hostlink-rw/commands"
)

var (
	ErrFCSValidationFail error = errors.New("hostlink: FCS checksum failed")
)

//Hostlink Holds hostlink connection datas
type Hostlink struct {
	Address string
	socket  net.Conn
}

//New Creates a new hostlink connection struct
func NewConnection(address string) *Hostlink {
	return &Hostlink{Address: address}
}

//Open Creates the underlying TCP connection
func (c *Hostlink) Open() error {
	conn, err := net.Dial("tcp", c.Address)

	if err != nil {
		return err
	}

	c.socket = conn
	return nil
}

//Close closes the underlying TCP connection
func (c *Hostlink) Close() {
	c.socket.Close()
}

//Read sends out an ReadCommand then reads the response
func (c *Hostlink) Read(command *commands.ReadCommand) (string, error) {
	cmd := command.ToString()
	resp, err := c.send(cmd)

	if err != nil {
		return "", err
	}

	if !command.CheckFCS(resp) {
		return "", ErrFCSValidationFail
	}

	return parseReadResponse(resp)
}

//Write sends out an WriteCommand then reads the response
func (c *Hostlink) Write(command *commands.WriteCommand) (bool, error) {
	cmd := command.ToString()
	resp, err := c.send(cmd)

	if err != nil {
		return false, err
	}

	if !command.CheckFCS(resp) {
		return false, ErrFCSValidationFail
	}

	return parseWriteResponse(resp)
}

func (c *Hostlink) send(cmd string) (string, error) {
	_, err := c.socket.Write([]byte(cmd))

	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(c.socket)
	resp, err := reader.ReadString('*')

	if err != nil {
		return "", err
	}

	return resp, nil
}
