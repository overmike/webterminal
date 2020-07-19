package terminal

import (
	"os"
	"testing"
)

func Test_action(t *testing.T) {
	type args struct {
		req  *SessionRequest
		ptmx *os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := action(tt.args.req, tt.args.ptmx); (err != nil) != tt.wantErr {
				t.Errorf("action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
