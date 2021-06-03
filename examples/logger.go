package main

import (
	"fmt"
	"time"
)

type ExampleLogger struct{}

func (e *ExampleLogger) Trace(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tTRACE\t"+format+"\n", v...)
}

func (e *ExampleLogger) Debug(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tDEBUG\t"+format+"\n", v...)
}

func (e *ExampleLogger) Info(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tINFO\t"+format+"\n", v...)
}

func (e *ExampleLogger) Warn(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tWARN\t"+format+"\n", v...)
}

func (e *ExampleLogger) Error(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tERROR\t"+format+"\n", v...)
}

func (e *ExampleLogger) Fatal(format string, v ...interface{}) {
	t := time.Now().Format(time.RFC3339)
	fmt.Printf(t+"\tFATAL\t"+format+"\n", v...)
}
