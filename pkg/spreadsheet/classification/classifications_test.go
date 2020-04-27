package classification

import "testing"

func TestContains(t *testing.T) {
	res := contains(transport, "deb.automatico -  cgmp- sem parar /sp*-")
	if !res {
		t.Error("should contain")
	}
}
