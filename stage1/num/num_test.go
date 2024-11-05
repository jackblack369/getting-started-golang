package num

import "testing"

func Test1(t *testing.T) {
	t.Log("Test1")
	const (
		// checker
		CHECK_TOPOLOGY int = iota
		CHECK_SSH_CONNECT
	)
	t.Log(CHECK_TOPOLOGY)
	t.Log(CHECK_SSH_CONNECT)
}
