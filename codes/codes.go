package codes

import "fmt"

type Code int

const (
	ReadError      Code = 1001
	UnmarshalError Code = 1002
	MarshalError   Code = 1003
)

var Codes = [3]Code{
	ReadError,
	UnmarshalError,
	MarshalError,
}

func (o Code) String() string {
	return fmt.Sprintf("%d", o)
}

func (o Code) Message() string {
	switch o {
	case 1001:
		return "READERROR"
	case 1002:
		return "UNMARSHALERROR"
	case 1003:
		return "MARSHALERROR"
	}

	return "UNKNOWN"
}
