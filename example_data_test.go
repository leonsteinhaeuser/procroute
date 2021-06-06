package routemod

import (
	"encoding/json"
	"errors"
	"net/http"
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
	data *data

	err    *HttpError
	logger Loggable
}

func (g *getExample) Get() (interface{}, *HttpError) {
	return g.data, g.err
}

func (g *getExample) Type() interface{} {
	return g.data
}

func (f *getExample) SetUrlParams(val map[string]string) {
	return
}

func (f *getExample) WithLogger(lggbl Loggable) {
	f.logger = lggbl
}

type getAllExample struct {
	data []data

	err    *HttpError
	logger Loggable
}

func (g *getAllExample) GetAll() ([]interface{}, *HttpError) {
	return []interface{}{g.data}, g.err
}

func (g *getAllExample) Type() interface{} {
	return g.data
}

func (g *getAllExample) SetUrlParams(val map[string]string) {
	return
}

func (g *getAllExample) WithLogger(lggbl Loggable) {
	g.logger = lggbl
}

type postExample struct {
	data *data

	err    *HttpError
	logger Loggable
}

func (p *postExample) Post() *HttpError {
	return p.err
}

func (p *postExample) Type() interface{} {
	return p.data
}

func (p *postExample) SetUrlParams(val map[string]string) {
	return
}

func (p *postExample) WithLogger(lggbl Loggable) {
	p.logger = lggbl
}

type updateExample struct {
	data *data

	err    *HttpError
	logger Loggable
}

func (u *updateExample) Update() *HttpError {
	return u.err
}

func (u *updateExample) Type() interface{} {
	return u.data
}

func (u *updateExample) SetUrlParams(val map[string]string) {
	return
}

func (u *updateExample) WithLogger(lggbl Loggable) {
	u.logger = lggbl
}

type deleteExample struct {
	data *data

	err    *HttpError
	logger Loggable
}

func (d *deleteExample) Delete() *HttpError {
	return d.err
}

func (d *deleteExample) Type() interface{} {
	return d.data
}

func (d *deleteExample) SetUrlParams(val map[string]string) {
	return
}

func (d *deleteExample) WithLogger(lggbl Loggable) {
	d.logger = lggbl
}

type fullExample struct {
	data *data

	err    *HttpError
	logger Loggable
}

func (f *fullExample) Get() (interface{}, *HttpError) {
	return f.data, f.err
}

func (f *fullExample) GetRoutePath() string {
	return "/{id}"
}

func (f *fullExample) GetAll() ([]interface{}, *HttpError) {
	return []interface{}{f.data}, f.err
}

func (f *fullExample) GetAllRoutePath() string {
	return "/all"
}

func (f *fullExample) Post() *HttpError {
	return f.err
}

func (f *fullExample) PostRoutePath() string {
	return "/all"
}

func (f *fullExample) Update() *HttpError {
	return f.err
}

func (f *fullExample) UpdateRoutePath() string {
	return "/all"
}

func (f *fullExample) Delete() *HttpError {
	return f.err
}

func (f *fullExample) DeleteRoutePath() string {
	return "/all"
}

func (f *fullExample) SetUrlParams(val map[string]string) {
	return
}

func (f *fullExample) Type() interface{} {
	return f.data
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
