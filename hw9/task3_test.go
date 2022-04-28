package hw9

import (
	"bytes"
	"testing"
)

func Test_printStrings(t *testing.T) {
	type args struct {
		objects []any
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name:    "Test #1",
			args:    args{objects: []any{Employee2{18}, "some string", 17, true, "87", "false"}},
			wantW:   "some string87false",
			wantErr: false,
		},
		{
			name:    "Test #2",
			args:    args{objects: []any{Customer{45}, 11, 17, true, "", false, ""}},
			wantW:   "",
			wantErr: false,
		},
		{
			name:    "Test #3",
			args:    args{objects: []any{}},
			wantW:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := printStrings(w, tt.args.objects...)
			if (err != nil) != tt.wantErr {
				t.Errorf("printStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("printStrings() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
