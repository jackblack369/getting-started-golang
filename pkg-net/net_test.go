package pkg_net

import (
	"fmt"
	"net"
	"testing"
)

// Checks if a port is available on the host machine
func isPortAvailable(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return false // Port is occupied
	}
	listener.Close()
	return true
}

func TestIsPortAvailable(t *testing.T) {
	res := isPortAvailable(4001)
	fmt.Println("res is ", res)
}
