package internal

import (
	"reflect"
	"testing"
)

func TestEthernetFrame_MarshalBinary(t *testing.T) {
	type fields struct {
		DestinationMAC []byte
		SourceMAC      []byte
		EtherType      uint16
		Payload        []byte
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
				DestinationMAC: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				SourceMAC:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				EtherType:      0x0800,
				Payload:        []byte{0x00, 0x00, 0x00, 0x00},
			},
			want: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Destination MAC
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Source MAC
				0x08, 0x00, // EtherType
				0x00, 0x00, 0x00, 0x00, // Payload
			},
			wantErr: false,
		},
		// invalid test case of Destination MAC
		{
			name: "invalid test case of Destination MAC",
			fields: fields{
				DestinationMAC: []byte{0x00, 0x00, 0x00, 0x00, 0x00},
				SourceMAC:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
				EtherType:      0x0800,
				Payload:        []byte{0x00, 0x00, 0x00, 0x00},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ef := &EthernetFrame{
				DestinationMAC: tt.fields.DestinationMAC,
				SourceMAC:      tt.fields.SourceMAC,
				EtherType:      tt.fields.EtherType,
				Payload:        tt.fields.Payload,
			}
			got, err := ef.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("EthernetFrame.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EthernetFrame.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
