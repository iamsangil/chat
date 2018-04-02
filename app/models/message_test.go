package models

import (
	"testing"
)

func TestMessageType_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		m       MessageType
		want    []byte
		wantErr bool
	}{
		{
			"type connect",
			TypeConnect,
			[]byte("connect"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.m.MarshalJSON()
			t.Logf("%s", got)
		})
	}
}

func TestMessageType_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		m       *MessageType
		args    args
		wantErr bool
	}{
		{
			"unmarshal send",
			new(MessageType),
			args{
				data: []byte("send"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("MessageType.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%s", tt.m)
		})
	}
}
