// +build test

package wol

import (
	"bytes"
	"log"
	"net"
	"testing"
	"time"
)

const mac = "12:34:56:78:90:ab"

func TestWakeOnLan(t *testing.T) {
	addr, err := net.ParseMAC(mac)
	if err != nil {
		t.Fatal(err)
	}

	packet := make([]byte, 6, 102)
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}

	for i := 0; i < 16; i++ {
		packet = append(packet, addr...)
	}

	go func() {
		time.Sleep(time.Millisecond)
		if err := WakeOnLan(mac); err != nil {
			log.Println(err)
		}
	}()

	conn, err := net.ListenPacket("udp", ":7777")
	if err != nil {
		t.Fatal(err)
	}

	if err := conn.SetDeadline(time.Now().Add(time.Second)); err != nil {
		t.Error(err)
	}

	buf := make([]byte, 102)
	if _, _, err := conn.ReadFrom(buf); err != nil {
		t.Error(err)
	}

	if !bytes.Equal(packet, buf) {
		t.Errorf("%v %v", packet, buf)
	}
}
