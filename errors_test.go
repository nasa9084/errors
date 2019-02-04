package errors

import (
	"io"
	"testing"
)

func TestWithError(t *testing.T) {
	tests := []struct {
		e1   error
		e2   error
		want string
	}{
		{io.EOF, io.ErrClosedPipe, "EOF/io: read/write on closed pipe"},
		{nil, io.EOF, "EOF"},
		{io.EOF, nil, "EOF"},
		{nil, nil, ""},
	}
	for _, tt := range tests {
		err := WithError(tt.e1, tt.e2)
		if err == nil {
			if tt.want == "" {
				continue
			}
			t.Errorf("want: %v, err: nil", tt.want)
		}
		got := err.Error()
		if got != tt.want {
			t.Errorf("WithError(%v, %v): got: %v, want: %v", tt.e1, tt.e2, got, tt.want)
		}
	}
}
