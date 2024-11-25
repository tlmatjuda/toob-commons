package machine

import (
	"fmt"
	"net"
)

// GetLocalIP
// The net.InterfaceAddrs() function returns a list of the system’s network addresses.
// It loops through each address and checks if:
//   - The address is an IPv4 (ipNet.IP.To4() != nil).
//   - It’s not a loopback address (!ipNet.IP.IsLoopback()), which means it excludes 127.0.0.1.
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// Check if the address is an IP address and not a loopback
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil // Return the first non-loopback IPv4 address
		}
	}

	return "", fmt.Errorf("no IP address found")
}
