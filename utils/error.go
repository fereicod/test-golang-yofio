package utils

type GeneralError struct {
	Message string
}

func (ge GeneralError) Error() string {
	return ge.Message
}
