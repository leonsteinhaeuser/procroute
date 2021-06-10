package procroute

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpError_write(t *testing.T) {
	type fields struct {
		Status    int
		ErrorCode string
		Message   string
	}
	type args struct {
		contentType string
		parser      Parser
		w           http.ResponseWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_nil_data",
			fields: fields{
				Status:    0,
				ErrorCode: "",
				Message:   "",
			},
			args: args{
				contentType: "",
				parser:      &exampleParser{},
				w:           nil,
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				Status:    500,
				ErrorCode: "0x11111jnsd",
				Message:   "non nil data",
			},
			args: args{
				contentType: "application/json",
				parser:      &exampleParser{},
				w:           httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HttpError{
				Status:    tt.fields.Status,
				ErrorCode: tt.fields.ErrorCode,
				Message:   tt.fields.Message,
			}
			if err := h.write(tt.args.contentType, tt.args.parser, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("HttpError.write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
