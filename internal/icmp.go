package internal

import (
	"encoding/binary"
	"errors"
)

/*
	  | Type | Code | Checksum | Payload |
		------------------------------------
		| 1    | 1    | 2        | 4-1472  |
*/
type ICMPPacket struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	Payload  []byte
}

func (ip *ICMPPacket) MarshalBinary() ([]byte, error) {
	if ip.Type < 0 || ip.Type > 7 {
		return nil, errors.New("Type must be between 0 and 7")
	}

	/*
		buf is a byte slice with a length of 4 bytes + the length of the payload
	*/
	buf := make([]byte, 4+len(ip.Payload))
	buf[0] = ip.Type
	buf[1] = ip.Code
	binary.BigEndian.PutUint16(buf[2:4], ip.Checksum)
	copy(buf[4:], ip.Payload)

	return buf, nil
}

func (ip *ICMPPacket) UnmarshalBinary(data []byte) error {
	/*
		An ICMP packet must be at least 4 bytes long
	*/
	if len(data) < 4 {
		return errors.New("data is too short to be an ICMP packet")
	}

	ip.Type = data[0]
	ip.Code = data[1]
	ip.Checksum = binary.BigEndian.Uint16(data[2:4])
	ip.Payload = data[4:]

	return nil
}
