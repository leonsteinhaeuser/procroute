package procroute

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewRouteSet(t *testing.T) {
	type args struct {
		basePath string
		parser   Parser
	}
	tests := []struct {
		name string
		args args
		want *RouteSet
	}{
		{
			name: "test_empty_data",
			args: args{
				basePath: "",
				parser:   nil,
			},
			want: &RouteSet{
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
		},
		{
			name: "test_non_empty_data",
			args: args{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			want: &RouteSet{
				parser:   &exampleParser{},
				router:   nil,
				basePath: "/api",
				routeSet: nil,
				logger:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouteSet(tt.args.basePath, tt.args.parser); !reflect.DeepEqual(got.basePath, tt.want.basePath) {
				t.Errorf("NewRouteSet() basePath = %v, want %v", got.basePath, tt.want.basePath)
			}

			if got := NewRouteSet(tt.args.basePath, tt.args.parser); !reflect.DeepEqual(got.parser, tt.want.parser) {
				t.Errorf("NewRouteSet() parser = %v, want %v", got.parser, tt.want.parser)
			}
		})
	}
}

func TestRouteSet_withRouter(t *testing.T) {
	rt := mux.NewRouter()

	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteSet
	}{
		{
			name: "test_empty_data",
			fields: fields{
				parser:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				rt: nil,
			},
			want: &RouteSet{
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
		},
		{
			name: "test_non_empty_data",
			fields: fields{
				parser:   &exampleParser{},
				basePath: "/api",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				rt: rt,
			},
			want: &RouteSet{
				parser:   &exampleParser{},
				router:   rt,
				basePath: "/api",
				routeSet: nil,
				logger:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if got := rm.withRouter(tt.args.rt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteSet.withRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteSet_withLogger(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		logger Loggable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteSet
	}{
		{
			name: "test_empty_data",
			fields: fields{
				parser:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				logger: nil,
			},
			want: &RouteSet{
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
		},
		{
			name: "test_non_empty_data",
			fields: fields{
				parser:   &exampleParser{},
				basePath: "/api",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				logger: &exampleLogger{},
			},
			want: &RouteSet{
				parser:   &exampleParser{},
				router:   nil,
				basePath: "/api",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if got := rm.withLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteSet.withLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteSet_withRouterBasePath(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		basePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteSet
	}{
		{
			name: "test_empty_data",
			fields: fields{
				parser:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				basePath: "",
			},
			want: &RouteSet{
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
		},
		{
			name: "test_non_empty_data",
			fields: fields{
				parser:   &exampleParser{},
				basePath: "/troll",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			args: args{
				basePath: "/api",
			},
			want: &RouteSet{
				parser:   &exampleParser{},
				router:   nil,
				basePath: "/api/troll",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if got := rm.withRouterBasePath(tt.args.basePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteSet.withBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteSet_AddRoutes(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		routes []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RouteSet
	}{
		{
			name: "test_empty_data",
			fields: fields{
				parser:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				routes: []interface{}{},
			},
			want: &RouteSet{
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
		},
		{
			name: "test_non_empty_data",
			fields: fields{
				parser:   &exampleParser{},
				basePath: "/troll",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			args: args{
				routes: []interface{}{
					&getExample{},
				},
			},
			want: &RouteSet{
				parser:   &exampleParser{},
				router:   nil,
				basePath: "/troll",
				routeSet: []interface{}{
					&getExample{},
				},
				logger: &exampleLogger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if got := rm.AddRoutes(tt.args.routes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteSet.AddRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteSet_buildPath(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		uriPath []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test_empty_data",
			fields: fields{
				parser:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				uriPath: []string{},
			},
			want: "",
		},
		{
			name: "test_non_empty_data_1",
			fields: fields{
				parser:   nil,
				basePath: "",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				uriPath: []string{
					"/test",
				},
			},
			want: "/test",
		},
		{
			name: "test_non_empty_data_2",
			fields: fields{
				parser:   nil,
				basePath: "/api",
				routeSet: nil,
				logger:   nil,
			},
			args: args{
				uriPath: []string{
					"/test",
				},
			},
			want: "/api/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if got := rm.buildPath(tt.args.uriPath...); got != tt.want {
				t.Errorf("RouteSet.buildPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteSet_build(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "empty_data",
			fields: fields{
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   nil,
				basePath: "",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_3",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_4",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: nil,
				logger:   &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_getExample_data",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&getExample{},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_getExample_error",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&getExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},

		{
			name: "non_empty_data_getAllExample_data",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&getAllExample{},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_getAllExample_error",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&getAllExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},

		{
			name: "non_empty_data_postExample_data",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&postExample{},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_postExample_error",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&postExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},

		{
			name: "non_empty_data_updateExample_data",
			fields: fields{
				parser: &exampleParser{},

				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&updateExample{},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_updateExample_error",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&updateExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},

		{
			name: "non_empty_data_deleteExample_data",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&deleteExample{},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_deleteExample_error",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&deleteExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},

		{
			name: "non_empty_data_fullExample_data",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&fullExample{},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "non_empty_data_fullExample_error",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&fullExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
		{
			name: "with_error_parser",
			fields: fields{
				parser:   &exampleParserError{},
				router:   mux.NewRouter(),
				basePath: "/api",
				routeSet: []interface{}{
					&fullExample{err: &HttpError{Status: 500, ErrorCode: "x923a1", Message: "something went wrong"}},
				},
				logger: &exampleLogger{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.build(); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_registerPostRoute(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt PostRoute
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
				parser:   nil,
				router:   nil,
				routeSet: []interface{}{},
				logger:   nil,
			},
			args: args{
				rt: nil,
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &postExample{},
			},
			wantErr: false,
		},
		{
			name: "test_non_nil_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &fullExample{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.registerPostRoute(tt.args.rt); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.registerPostRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_definePostRoute(t *testing.T) {
	type fields struct {
		basePath string
		parser   Parser
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		rt PostRoute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil_data",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &postExample{},
			},
		},
		{
			name: "with_error_return",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &postExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil_data_with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &postExample{},
			},
		},
		{
			name: "with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &postExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "with_error",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &fullExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil",
			fields: fields{
				basePath: "/",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &postExample{},
			},
		},
		{
			name: "nil_error_parser",
			fields: fields{
				basePath: "/",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("POST", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &postExample{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := NewRouteSet(tt.fields.basePath, tt.fields.parser)

			rm.definePostRoute(tt.args.w, tt.args.r, tt.args.rt)
		})
	}
}

func TestRouteSet_registerGetRoute(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt      GetRoute
		getPath string
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
				parser:   nil,
				router:   nil,
				basePath: "",
				routeSet: []interface{}{},
				logger:   nil,
			},
			args: args{
				rt:      nil,
				getPath: "",
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				basePath: "",
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt:      &fullExample{},
				getPath: "",
			},
			wantErr: false,
		},
		{
			name: "test_non_nil_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &fullExample{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.registerGetRoute(tt.args.rt); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.registerGetRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_defineGetRoute(t *testing.T) {
	type fields struct {
		basePath string
		parser   Parser
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		rt GetRoute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil_data",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getExample{},
			},
		},
		{
			name: "with_error_return",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil_data_with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getExample{},
			},
		},
		{
			name: "with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "with_error",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &fullExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil",
			fields: fields{
				basePath: "/",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getExample{},
			},
		},
		{
			name: "nil_error_parser",
			fields: fields{
				basePath: "/",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getExample{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := NewRouteSet(tt.fields.basePath, tt.fields.parser)

			rm.defineGetRoute(tt.args.w, tt.args.r, tt.args.rt)
		})
	}
}

func TestRouteSet_registerGetAllRoute(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt GetAllRoute
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
				parser:   nil,
				router:   nil,
				routeSet: []interface{}{},
				logger:   nil,
			},
			args: args{
				rt: nil,
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &getAllExample{},
			},
			wantErr: false,
		},
		{
			name: "test_non_nil_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &fullExample{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.registerGetAllRoute(tt.args.rt); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.registerGetAllRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_defineGetAllRoute(t *testing.T) {
	type fields struct {
		basePath string
		parser   Parser
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		rt GetAllRoute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil_data",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getAllExample{},
			},
		},
		{
			name: "with_error_return",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getAllExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil_data_with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getAllExample{},
			},
		},
		{
			name: "with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getAllExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "with_error",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &fullExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil",
			fields: fields{
				basePath: "/",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getAllExample{},
			},
		},
		{
			name: "nil_error_parser",
			fields: fields{
				basePath: "/",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &getAllExample{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := NewRouteSet(tt.fields.basePath, tt.fields.parser)

			rm.defineGetAllRoute(tt.args.w, tt.args.r, tt.args.rt)
		})
	}
}

func TestRouteSet_registerUpdateRoute(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt UpdateRoute
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
				parser:   nil,
				router:   nil,
				routeSet: []interface{}{},
				logger:   nil,
			},
			args: args{
				rt: nil,
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &updateExample{},
			},
			wantErr: false,
		},
		{
			name: "test_non_nil_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &fullExample{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.registerUpdateRoute(tt.args.rt); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.registerUpdateRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_defineUpdateRoute(t *testing.T) {
	type fields struct {
		basePath string
		parser   Parser
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		rt UpdateRoute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil_data",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("PUT", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &updateExample{},
			},
		},
		{
			name: "with_error_return",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &updateExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil_data_with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("PUT", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &updateExample{},
			},
		},
		{
			name: "with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &updateExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "with_error",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &fullExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil",
			fields: fields{
				basePath: "/",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/{id}", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &updateExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "",
						Message:   "oops",
					},
				},
			},
		},
		{
			name: "nil_error_parser",
			fields: fields{
				basePath: "/",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/{id}", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &updateExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "",
						Message:   "oops",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := NewRouteSet(tt.fields.basePath, tt.fields.parser)

			rm.defineUpdateRoute(tt.args.w, tt.args.r, tt.args.rt)
		})
	}
}

func TestRouteSet_registerDeleteRoute(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt DeleteRoute
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
				parser:   nil,
				router:   nil,
				routeSet: []interface{}{},
				logger:   nil,
			},
			args: args{
				rt: nil,
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &deleteExample{},
			},
			wantErr: false,
		},
		{
			name: "test_non_nil_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &fullExample{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.registerDeleteRoute(tt.args.rt); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.registerDeleteRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_defineDeleteRoute(t *testing.T) {
	type fields struct {
		basePath string
		parser   Parser
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		rt DeleteRoute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil_data",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &deleteExample{},
			},
		},
		{
			name: "with_error_return",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &deleteExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil_data_with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &deleteExample{},
			},
		},
		{
			name: "with_error_parser",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &deleteExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "with_error",
			fields: fields{
				basePath: "/api",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &fullExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil",
			fields: fields{
				basePath: "/",
				parser:   &exampleParser{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &deleteExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
		{
			name: "nil_error_parser",
			fields: fields{
				basePath: "/",
				parser:   &exampleParserError{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/", strings.NewReader(`{"name": "sample2", "value": 2}`)),
				rt: &deleteExample{
					err: &HttpError{
						Status:    500,
						ErrorCode: "asd",
						Message:   "expected error during test",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := NewRouteSet(tt.fields.basePath, tt.fields.parser)

			rm.defineDeleteRoute(tt.args.w, tt.args.r, tt.args.rt)
		})
	}
}

func TestRouteSet_registerRawRoute(t *testing.T) {
	type fields struct {
		parser   Parser
		router   *mux.Router
		basePath string
		routeSet []interface{}
		logger   Loggable
	}
	type args struct {
		rt RawRoute
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
				parser:   nil,
				router:   nil,
				routeSet: []interface{}{},
				logger:   nil,
			},
			args: args{
				rt: nil,
			},
			wantErr: true,
		},
		{
			name: "test_non_nil_data_1",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &rawExample{},
			},
			wantErr: false,
		},
		{
			name: "test_non_nil_data_2",
			fields: fields{
				parser:   &exampleParser{},
				router:   mux.NewRouter(),
				routeSet: []interface{}{},
				logger:   &exampleLogger{},
			},
			args: args{
				rt: &fullExample{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser:   tt.fields.parser,
				router:   tt.fields.router,
				basePath: tt.fields.basePath,
				routeSet: tt.fields.routeSet,
				logger:   tt.fields.logger,
			}
			if err := rm.registerRawRoute(tt.args.rt); (err != nil) != tt.wantErr {
				t.Errorf("RouteSet.registerRawRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouteSet_unmarshal(t *testing.T) {
	type fields struct {
		parser Parser
	}
	type args struct {
		bts []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      interface{}
		wantError bool
	}{
		{
			name: "with_error",
			fields: fields{
				parser: &exampleParserError{},
			},
			args: args{
				bts: []byte(`{"name": "sample2", "value": 2}`),
			},
			want:      nil,
			wantError: true,
		},
		{
			name: "with_non_error",
			fields: fields{
				parser: &exampleParser{},
			},
			args: args{
				bts: []byte(`{"name": "sample2", "value": 2}`),
			},
			want: map[string]interface{}{
				"name":  "sample2",
				"value": float64(2),
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser: tt.fields.parser,
			}

			got, err := rm.unmarshal(tt.args.bts)
			if (err != nil) != tt.wantError {
				t.Errorf("RouteSet.unmarshal() received error = %+#v, want error = %+#v", err, tt.wantError)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteSet.unmarshal() is not equal \ngot = %+#v\nwant = %+#v", got, tt.want)
			}
		})
	}
}

func TestRouteSet_marshal(t *testing.T) {
	type fields struct {
		parser Parser
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []byte
		wantError bool
	}{
		{
			name: "with_error",
			fields: fields{
				parser: &exampleParserError{},
			},
			args: args{
				data: data{
					Name:  "sample",
					Value: 1,
				},
			},
			want:      nil,
			wantError: true,
		},
		{
			name: "with_error",
			fields: fields{
				parser: &exampleParser{},
			},
			args: args{
				data: data{
					Name:  "sample",
					Value: 1,
				},
			},
			want:      []byte(`{"Name":"sample","Value":1}`),
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &RouteSet{
				parser: tt.fields.parser,
			}

			got, httperr := rm.marshal(tt.args.data)
			if (httperr != nil) != tt.wantError {
				t.Errorf("RouteSet.marshal() received error got = %v, want %v", httperr, tt.wantError)
			}

			if !reflect.DeepEqual(string(got), string(tt.want)) {
				t.Errorf("RouteSet.marshal() is not equal \ngot = %+#v\nwant %+#v", string(got), string(tt.want))
			}
		})
	}
}

func TestRouteSet_doHttpOp(t *testing.T) {
	type args struct {
		routeController interface{}
		r               *http.Request
	}
	tests := []struct {
		name      string
		args      args
		want      interface{}
		wantError bool
	}{
		{
			name: "with_data",
			args: args{
				routeController: &fullExample{},
				r: func() *http.Request {
					jsonBytes := []byte(`{"foo": "bar"}`)

					request, _ := http.NewRequest(
						"get",
						"http://localhost:8080/api/example",
						bytes.NewReader(jsonBytes),
					)

					return request
				}(),
			},
			want: map[string]interface{}{
				"foo": "bar",
			},
			wantError: false,
		},
		{
			name: "with_empty_data",
			args: args{
				routeController: &fullExample{},
				r: func() *http.Request {
					jsonBytes := []byte(``)

					request, _ := http.NewRequest(
						"get",
						"http://localhost:8080/api/example",
						bytes.NewReader(jsonBytes),
					)

					return request
				}(),
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "with_wrong_data",
			args: args{
				routeController: &fullExample{},
				r: func() *http.Request {
					jsonBytes := []byte(`{`)

					request, _ := http.NewRequest(
						"get",
						"http://localhost:8080/api/example",
						bytes.NewReader(jsonBytes),
					)

					return request
				}(),
			},
			want:      nil,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RouteSet{
				parser: &exampleParser{},
			}

			got, err := rs.doHttpOp(tt.args.routeController, tt.args.r)
			if (err != nil) != tt.wantError {
				t.Errorf("RouteSet.doHttpOp() received error \ngot = %+#v\nwant %+#v", err, tt.wantError)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteSet.doHttpOp() is not equal\ngot = %+#v\nwant %+#v", got, tt.want)
			}
		})
	}
}
