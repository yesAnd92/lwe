package utils

import (
	"fmt"
	"net"
)

// PortAvailable whether the port can be used
func PortAvailable(port int) bool {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}

	defer listener.Close()
	return true

}
