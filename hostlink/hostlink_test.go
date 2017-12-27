package hostlink

import (
	"testing"
)

func TestOpenWillFail(t *testing.T) {
	link := NewConnection("localhost:8080")
	err := link.Open()

	if err == nil {
		t.Fatal("Fail is expected")
	}
}
