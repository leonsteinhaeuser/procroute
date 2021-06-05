package main

import (
	"fmt"
	"time"
)

type ExampleLogger struct{}

func (e *ExampleLogger) buildLogerEntry(prefix, format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\t"+prefix+"\t"+format+"\n", v...)
}

func (e *ExampleLogger) Trace(format string, v ...interface{}) {
	e.buildLogerEntry("TRACE", format, v...)
}

func (e *ExampleLogger) Debug(format string, v ...interface{}) {
	e.buildLogerEntry("DEBUG", format, v...)
}

func (e *ExampleLogger) Info(format string, v ...interface{}) {
	e.buildLogerEntry("INFO", format, v...)
}

func (e *ExampleLogger) Warn(format string, v ...interface{}) {
	e.buildLogerEntry("WARN", format, v...)
}

func (e *ExampleLogger) Error(format string, v ...interface{}) {
	e.buildLogerEntry("ERROR", format, v...)
}

func (e *ExampleLogger) Fatal(format string, v ...interface{}) {
	e.buildLogerEntry("FATAL", format, v...)
}
