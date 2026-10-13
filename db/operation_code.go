package db

type OperationCode int

const (
	Error OperationCode = iota
	Success
)

func (c OperationCode) String() string {
	switch c {
	case Success:
		return "Success"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}
