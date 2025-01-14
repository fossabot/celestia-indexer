package types

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestHex_UnmarshalJSON(t *testing.T) {
	type test struct {
		Field Hex `json:"field"`
	}

	tests := []struct {
		name    string
		json    []byte
		want    []byte
		wantErr bool
	}{
		{
			name: "test 1",
			json: []byte(`{"field": "deadbeaf"}`),
			want: []byte{0xde, 0xad, 0xbe, 0xaf},
		}, {
			name: "test 2",
			json: []byte(`{"field": "DEADBEAF"}`),
			want: []byte{0xde, 0xad, 0xbe, 0xaf},
		}, {
			name:    "test 3",
			json:    []byte(`{"field": "DEADBEAFG"}`),
			wantErr: true,
		}, {
			name: "test 4",
			json: []byte(`{"field": null}`),
			want: nil,
		}, {
			name:    "test 5",
			json:    []byte(`{"field": 1234}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var val test
			err := json.Unmarshal(tt.json, &val)
			require.Equal(t, err != nil, tt.wantErr)
			if err == nil {
				require.Equal(t, tt.want, []byte(val.Field))
			}
		})
	}
}

func TestHex_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		h       Hex
		want    []byte
		wantErr bool
	}{
		{
			name: "test 1",
			h:    Hex([]byte{0xde, 0xad, 0xbe, 0xaf}),
			want: []byte(`"DEADBEAF"`),
		}, {
			name: "test 2",
			h:    Hex(nil),
			want: []byte("null"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.MarshalJSON()
			require.Equal(t, err != nil, tt.wantErr)
			if err == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
