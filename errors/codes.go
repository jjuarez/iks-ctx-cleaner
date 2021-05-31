package errors

type Code int

const (
	ReadError      Code = 1001
	UnmarshalError Code = 1002
	MarshalError   Code = 1003
)
