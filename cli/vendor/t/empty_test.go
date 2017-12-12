package t

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	expect := "yes"

	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}
