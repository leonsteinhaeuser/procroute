package procroute

type Loggable interface {
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

// WithLogger provides an optional interface that is used to attach the logger to the route or middleware
type WithLogger interface {
	// WithLogger represents the function that must be implemented to inject the Loggable object into the concrete implementation.
	// In most cases the implementation is similar to:
	//
	// Example:
	//  func (m *MyType) WithLogger(lggbl Loggable) {
	//     m.logger = lggbl
	//  }
	WithLogger(Loggable)
}
