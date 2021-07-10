package procroute

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type exampleParser struct{}

func (e *exampleParser) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (e *exampleParser) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (e *exampleParser) MimeType() string {
	return "application/json"
}

type exampleParserError struct{}

func (e *exampleParserError) Unmarshal(data []byte, v interface{}) error {
	return errors.New("exampleParserError")
}

func (e *exampleParserError) Marshal(v interface{}) ([]byte, error) {
	return nil, errors.New("exampleParserError")
}

func (e *exampleParserError) MimeType() string {
	return ""
}

type exampleLogger struct{}

func (e *exampleLogger) Trace(format string, v ...interface{}) {}
func (e *exampleLogger) Debug(format string, v ...interface{}) {}
func (e *exampleLogger) Info(format string, v ...interface{})  {}
func (e *exampleLogger) Warn(format string, v ...interface{})  {}
func (e *exampleLogger) Error(format string, v ...interface{}) {}
func (e *exampleLogger) Fatal(format string, v ...interface{}) {}

type data struct {
	Name  string
	Value int
}

type getExample struct {
	err    *HttpError
	logger Loggable
}

func (g *getExample) Get(requestData interface{}) (interface{}, *HttpError) {
	return requestData, g.err
}

func (f *getExample) SetUrlParams(val map[string]string) {
	return
}

func (g *getExample) SetQueryParams(args url.Values) {
	return
}

func (g *getExample) WithLogger(lggbl Loggable) {
	g.logger = lggbl
}

type getAllExample struct {
	err    *HttpError
	logger Loggable
}

func (g *getAllExample) GetAll(requestData interface{}) ([]interface{}, *HttpError) {
	return []interface{}{requestData}, g.err
}

func (g *getAllExample) SetUrlParams(val map[string]string) {
	return
}

func (g *getAllExample) SetQueryParams(args url.Values) {
	return
}

func (g *getAllExample) WithLogger(lggbl Loggable) {
	g.logger = lggbl
}

type postExample struct {
	err    *HttpError
	logger Loggable
}

func (p *postExample) Post(requestData interface{}) *HttpError {
	return p.err
}

func (p *postExample) SetUrlParams(val map[string]string) {
	return
}

func (p *postExample) SetQueryParams(args url.Values) {
	return
}

func (p *postExample) WithLogger(lggbl Loggable) {
	p.logger = lggbl
}

type updateExample struct {
	err    *HttpError
	logger Loggable
}

func (u *updateExample) Update(requestData interface{}) *HttpError {
	return u.err
}

func (u *updateExample) SetUrlParams(val map[string]string) {
	return
}

func (u *updateExample) SetQueryParams(args url.Values) {
	return
}

func (u *updateExample) WithLogger(lggbl Loggable) {
	u.logger = lggbl
}

type deleteExample struct {
	err    *HttpError
	logger Loggable
}

func (d *deleteExample) Delete(requestData interface{}) *HttpError {
	return d.err
}

func (d *deleteExample) SetUrlParams(val map[string]string) {
	return
}

func (d *deleteExample) SetQueryParams(args url.Values) {
	return
}

func (d *deleteExample) WithLogger(lggbl Loggable) {
	d.logger = lggbl
}

type rawExample struct {
	err    *HttpError
	logger Loggable
}

func (d *rawExample) Raw(w http.ResponseWriter, r *http.Request) {
	d.logger.Info("raw hit")
	return
}

func (d *rawExample) HttpMethods() []string {
	return []string{"GET", "OPTIONS"}
}

func (d *rawExample) RawRoutePath() string {
	return "/raw"
}

func (d *rawExample) WithLogger(lggbl Loggable) {
	d.logger = lggbl
}

type fullExample struct {
	err    *HttpError
	logger Loggable
}

func (f *fullExample) Get(requestData interface{}) (interface{}, *HttpError) {
	return requestData, f.err
}

func (f *fullExample) GetRoutePath() string {
	return "/{id}"
}

func (f *fullExample) GetAll(requestData interface{}) ([]interface{}, *HttpError) {
	return []interface{}{requestData}, f.err
}

func (f *fullExample) GetAllRoutePath() string {
	return "/all"
}

func (f *fullExample) Post(requestData interface{}) *HttpError {
	return f.err
}

func (f *fullExample) PostRoutePath() string {
	return "/all"
}

func (f *fullExample) Update(requestData interface{}) *HttpError {
	return f.err
}

func (f *fullExample) UpdateRoutePath() string {
	return "/all"
}

func (f *fullExample) Delete(requestData interface{}) *HttpError {
	return f.err
}

func (f *fullExample) DeleteRoutePath() string {
	return "/all"
}

func (f *fullExample) Raw(w http.ResponseWriter, r *http.Request) {
	f.logger.Info("raw hit")
	return
}

func (f *fullExample) HttpMethods() []string {
	return []string{"GET", "OPTIONS"}
}

func (f *fullExample) RawRoutePath() string {
	return "/raw"
}

func (f *fullExample) SetUrlParams(val map[string]string) {
	return
}

func (f *fullExample) SetQueryParams(args url.Values) {
	return
}

func (f *fullExample) WithLogger(lggbl Loggable) {
	f.logger = lggbl
}

type exampleMiddleware struct {
	logger Loggable
}

func (e *exampleMiddleware) Middleware(h http.Handler) http.Handler {
	return h
}

func (e *exampleMiddleware) WithLogger(lggbl Loggable) {
	e.logger = lggbl
}
