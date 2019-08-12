package wolgo

import "net"

// WakeOnLan turns on the computer with specific mac address by a network message.
func WakeOnLan(mac string) error {
	addr, err := net.ParseMAC(mac)
	if err != nil {
		return err
	}

	packet := make([]byte, 6, 102)
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}

	for i := 0; i < 16; i++ {
		packet = append(packet, addr...)
	}

	conn, err := net.Dial("udp", "255.255.255.255:7")
	if err != nil {
		return err
	}

	_, err = conn.Write(packet)
	return err
}
