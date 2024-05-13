package internal

import (
	"encoding/binary"
	"errors"
	"net"
)

/*
| Destination MAC | Source MAC | EtherType | Payload       |
------------------------------------------------------------
| 6 bytes         | 6 bytes    | 2 bytes   | 46-1500 bytes |
*/
type EthernetFrame struct {
	DestinationMAC net.HardwareAddr
	SourceMAC      net.HardwareAddr
	EtherType      uint16
	Payload        []byte
}

func (ef *EthernetFrame) MarshalBinary() ([]byte, error) {
	/*
		Destination MAC address must be 6 bytes
	*/
	if len(ef.DestinationMAC) != 6 {
		return nil, errors.New("Destination MAC address must be 6 bytes")
	}

	/*
	  buf is a byte slice with a length of 14 bytes
	*/
	buf := make([]byte, 14+len(ef.Payload))
	copy(buf[0:6], ef.DestinationMAC)
	copy(buf[6:12], ef.SourceMAC)
	binary.BigEndian.PutUint16(buf[12:14], ef.EtherType)
	copy(buf[14:], ef.Payload)

	return buf, nil
}

func (ef *EthernetFrame) UnmarshalBinary(data []byte) error {
	/*
		An Ethernet frame must be at least 14 bytes long
	*/
	if len(data) < 14 {
		return errors.New("data is too short to be an Ethernet frame")
	}

	ef.DestinationMAC = net.HardwareAddr(data[0:6])
	ef.SourceMAC = net.HardwareAddr(data[6:12])
	ef.EtherType = binary.BigEndian.Uint16(data[12:14])
	ef.Payload = data[14:]

	return nil
}
