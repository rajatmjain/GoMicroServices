package details

import (
	"log"
	"net"
	"os"
)

func GetHostName() (string, error) {
	hostName, error := os.Hostname()
	return hostName, error
}

func GetIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, err
}
