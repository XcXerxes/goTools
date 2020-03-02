package interfaces

import "testing"

func TestLikeInterface(t *testing.T) {
	b := new(B)
	b.A.Parent = b
	b.Run()
}
