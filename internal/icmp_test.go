package internal

import (
	"reflect"
	"testing"
)

func TestICMPPacket_MarshalBinary(t *testing.T) {
	type fields struct {
		Type     uint8
		Code     uint8
		Checksum uint16
		Payload  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// valid test case
		{
			name: "valid test case",
			fields: fields{
				Type:     0x01,
				Code:     0x00,
				Checksum: 0x0000,
				Payload:  []byte{0x00, 0x00, 0x00, 0x00},
			},
			want: []byte{
				0x01, 0x00, // Type, Code
				0x00, 0x00, // Checksum
				0x00, 0x00, 0x00, 0x00, // Payload
			},
			wantErr: false,
		},
		// invalid test case of Type
		{
			name: "invalid test case of Type",
			fields: fields{
				Type:     0x08,
				Code:     0x00,
				Checksum: 0x0000,
				Payload:  []byte{0x00, 0x00, 0x00, 0x00},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &ICMPPacket{
				Type:     tt.fields.Type,
				Code:     tt.fields.Code,
				Checksum: tt.fields.Checksum,
				Payload:  tt.fields.Payload,
			}
			got, err := ip.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("ICMPPacket.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ICMPPacket.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
