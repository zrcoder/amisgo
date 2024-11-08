package comp

import (
	"testing"
)

func Test_getInnerApiID(t *testing.T) {
	for i := int32(0); i < 10; i++ {
		x := getInnerApiID()
		if x != i {
			t.Errorf("got %d, want %d", x, i)
		}
	}
}
