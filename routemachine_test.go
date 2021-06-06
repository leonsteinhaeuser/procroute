package routemod

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewRouteMachine(t *testing.T) {
	type args struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	tests := []struct {
		name string
		args args
		want *RouteMachine
	}{
		{
			name: "non_nil_values",
			args: args{
				addr:     "127.0.0.1",
				port:     8080,
				basePath: "/api",
				loggable: &exampleLogger{},
			},
			want: &RouteMachine{
				server: &http.Server{
					Addr: "127.0.0.1:8080",
				},
				basePath: "/api",
				logger:   &exampleLogger{},
			},
		},
		{
			name: "nil_values",
			args: args{
				addr:     "",
				port:     0,
				basePath: "",
				loggable: nil,
			},
			want: &RouteMachine{
				server: &http.Server{
					Addr: ":0",
				},
				basePath: "",
				logger:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouteMachine(tt.args.addr, tt.args.port, tt.args.basePath, tt.args.loggable); !reflect.DeepEqual(got.server.Addr, tt.want.server.Addr) {
				t.Errorf("NewRouteMachine() .server.Addr = %v, want %v", got.server.Addr, tt.want.server.Addr)
			}

			if got := NewRouteMachine(tt.args.addr, tt.args.port, tt.args.basePath, tt.args.loggable); !reflect.DeepEqual(got.basePath, tt.want.basePath) {
				t.Errorf("NewRouteMachine() .basePath = %v, want %v", got.basePath, tt.want.basePath)
			}
		})
	}
}

func TestRouteMachine_AddRouteSet(t *testing.T) {
	type fields struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	type args struct {
		routeSet *RouteSet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: nil,
			},
			args: args{
				routeSet: nil,
			},
			wantErr: true,
		},
		{
			name: "non_nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: &exampleLogger{},
			},
			args: args{
				routeSet: NewRouteSet("/api", &exampleParser{}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rtm := NewRouteMachine(tt.fields.addr, tt.fields.port, tt.fields.basePath, tt.fields.loggable)
			if err := rtm.AddRouteSet(tt.args.routeSet); (err != nil) != tt.wantErr {
				t.Errorf("RouteMachine.AddRouteSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteMachine_SetReadTimeout(t *testing.T) {
	type fields struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteMachine
	}{
		{
			name: "nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: nil,
			},
			args: args{
				timeout: time.Second * 5,
			},
			want: &RouteMachine{
				server: &http.Server{
					ReadTimeout: time.Second * 5,
				},
			},
		},
		{
			name: "non_nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: &exampleLogger{},
			},
			args: args{
				timeout: time.Second * 5,
			},
			want: &RouteMachine{
				server: &http.Server{
					ReadTimeout: time.Second * 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rtm := NewRouteMachine(tt.fields.addr, tt.fields.port, tt.fields.basePath, tt.fields.loggable)
			if got := rtm.SetReadTimeout(tt.args.timeout); !reflect.DeepEqual(got.server.ReadTimeout, tt.want.server.ReadTimeout) {
				t.Errorf("RouteMachine.SetReadTimeout() = %v, want %v", got.server.ReadTimeout, tt.want.server.ReadTimeout)
			}
		})
	}
}

func TestRouteMachine_SetReadHeaderTimeout(t *testing.T) {
	type fields struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteMachine
	}{
		{
			name: "nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: nil,
			},
			args: args{
				timeout: time.Second * 5,
			},
			want: &RouteMachine{
				server: &http.Server{
					ReadHeaderTimeout: time.Second * 5,
				},
			},
		},
		{
			name: "non_nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: &exampleLogger{},
			},
			args: args{
				timeout: time.Second * 5,
			},
			want: &RouteMachine{
				server: &http.Server{
					ReadHeaderTimeout: time.Second * 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rtm := NewRouteMachine(tt.fields.addr, tt.fields.port, tt.fields.basePath, tt.fields.loggable)
			if got := rtm.SetReadHeaderTimeout(tt.args.timeout); !reflect.DeepEqual(got.server.ReadHeaderTimeout, tt.want.server.ReadHeaderTimeout) {
				t.Errorf("RouteMachine.SetReadHeaderTimeout() = %v, want %v", got.server.ReadHeaderTimeout, tt.want.server.ReadHeaderTimeout)
			}
		})
	}
}

func TestRouteMachine_SetIdleTimeout(t *testing.T) {
	type fields struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteMachine
	}{
		{
			name: "nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: nil,
			},
			args: args{
				timeout: time.Second * 5,
			},
			want: &RouteMachine{
				server: &http.Server{
					IdleTimeout: time.Second * 5,
				},
			},
		},
		{
			name: "non_nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: &exampleLogger{},
			},
			args: args{
				timeout: time.Second * 5,
			},
			want: &RouteMachine{
				server: &http.Server{
					IdleTimeout: time.Second * 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rtm := NewRouteMachine(tt.fields.addr, tt.fields.port, tt.fields.basePath, tt.fields.loggable)
			if got := rtm.SetIdleTimeout(tt.args.timeout); !reflect.DeepEqual(got.server.IdleTimeout, tt.want.server.IdleTimeout) {
				t.Errorf("RouteMachine.SetIdleTimeout() = %v, want %v", got.server.IdleTimeout, tt.want.server.IdleTimeout)
			}
		})
	}
}

func TestRouteMachine_Start(t *testing.T) {
	type fields struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	type args struct {
		routeMachine *RouteMachine
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil_values",
			args: args{
				routeMachine: NewRouteMachine("127.0.0.1", 7654, "/api", &exampleLogger{}),
			},
			wantErr: true,
		},
		{
			name: "non_nil_values",
			args: args{
				routeMachine: func() *RouteMachine {
					rm := NewRouteMachine("127.0.0.1", 7654, "/api", &exampleLogger{})
					if err := rm.AddRouteSet(NewRouteSet("/sample", &exampleParser{}).AddRoutes(&getExample{})); err != nil {
						panic(err)
					}
					return rm
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.routeMachine.Start(); (err != nil) != tt.wantErr {
				t.Errorf("RouteMachine.Start() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := tt.args.routeMachine.Stop(); err != nil {
				t.Errorf("RouteMachine.Start() stop error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteMachine_Stop(t *testing.T) {
	type fields struct {
		addr     string
		port     uint16
		basePath string
		loggable Loggable
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: nil,
			},
			wantErr: false,
		},
		{
			name: "non_nil_values",
			fields: fields{
				addr:     "",
				port:     0,
				basePath: "/",
				loggable: &exampleLogger{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rtm := NewRouteMachine(tt.fields.addr, tt.fields.port, tt.fields.basePath, tt.fields.loggable)
			if err := rtm.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("RouteMachine.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
