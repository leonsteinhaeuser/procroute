package procroute

// Parser provides the interface that must be implemented to marshal and unmarshal the data sent during http request and http responses.
type Parser interface {
	// Unmarshal parses the encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an error.
	//
	// Example:
	//  func (m *JsonParser) Unmarshal(data []byte, v interface{}) error {
	//  	return json.Unmarshal(data, v)
	//  }
	Unmarshal(data []byte, v interface{}) error
	// Marshal returns the encoded data as byte slice.
	//
	// Exaple:
	//  func (m *JsonParser) Marshal(v interface{}) ([]byte, error) {
	//  	return json.Marshal(v)
	//  }
	Marshal(v interface{}) ([]byte, error)
	// MimeType returns the associated mime type in string representation.
	// A list of available MIME types can be found at: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	//
	// Example:
	//  func (m *JsonParser) MimeType() string {
	//  	return "application/json"
	//  }
	MimeType() string
}
