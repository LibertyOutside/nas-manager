package settings

import "testing"

func TestInitClient(t *testing.T) {
	err := InitClient()
	if err != nil {
		t.Fail()
	}
}
