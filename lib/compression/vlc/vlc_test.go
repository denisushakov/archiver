package vlc

import (
	"reflect"
	"testing"

	"github.com/denisushakov/archiver/lib/compression/vlc/table"
	"github.com/denisushakov/archiver/lib/compression/vlc/table/shannon_fano"
)

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!ted",
			want: "001000100110100101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str, table.EncodingTable{}); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New(shannon_fano.NewGenerator())

			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name       string
		encodeData []byte
		want       string
	}{
		{
			name:       "base test",
			encodeData: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:       "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New(shannon_fano.NewGenerator())

			if got := decoder.Decode(tt.encodeData); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
