package minedash

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func Test_parseJsonBody(t *testing.T) {
	type args struct {
		res  string
		dest interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "With empty body",
			args: args{
				res:  "",
				dest: &EntityAllResponse{},
			},
			wantErr: true,
		},
		{
			name: "With braces only body",
			args: args{
				res:  "{}",
				dest: &EntityAllResponse{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &http.Response{Body: io.NopCloser(bytes.NewBufferString(tt.args.res))}
			if err := parseJsonBody(res, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("parseJsonBody() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%v", tt.args.dest)
		})
	}
}
