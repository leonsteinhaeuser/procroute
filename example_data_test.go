package routemod

import (
	"encoding/json"
	"errors"
)

type exampleParser struct{}

func (e *exampleParser) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (e *exampleParser) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

type exampleParserError struct{}

func (e *exampleParserError) Unmarshal(data []byte, v interface{}) error {
	return errors.New("exampleParserError")
}

func (e *exampleParserError) Marshal(v interface{}) ([]byte, error) {
	return nil, errors.New("exampleParserError")
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

	err *HttpError
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

type getAllExample struct {
	data []data

	err *HttpError
}

func (g *getAllExample) GetAll() ([]interface{}, *HttpError) {
	return []interface{}{g.data}, g.err
}

func (g *getAllExample) Type() interface{} {
	return g.data
}

func (f *getAllExample) SetUrlParams(val map[string]string) {
	return
}

type postExample struct {
	data *data

	err *HttpError
}

func (p *postExample) Post() *HttpError {
	return p.err
}

func (p *postExample) Type() interface{} {
	return p.data
}

func (f *postExample) SetUrlParams(val map[string]string) {
	return
}

type updateExample struct {
	data *data

	err *HttpError
}

func (u *updateExample) Update() *HttpError {
	return u.err
}

func (u *updateExample) Type() interface{} {
	return u.data
}

func (f *updateExample) SetUrlParams(val map[string]string) {
	return
}

type deleteExample struct {
	data *data

	err *HttpError
}

func (u *deleteExample) Delete() *HttpError {
	return u.err
}

func (u *deleteExample) Type() interface{} {
	return u.data
}

func (f *deleteExample) SetUrlParams(val map[string]string) {
	return
}

type fullExample struct {
	data *data

	err *HttpError
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
