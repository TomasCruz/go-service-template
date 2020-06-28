package presenter

// ErrStruct contains errorMessage
type ErrStruct struct {
	Msg string `json:"errorMessage" example:"A horrible, terrible, absolutely awful error"`
}

// MsgStruct contains a message
type MsgStruct struct {
	Msg string `json:"message" example:"A message"`
}
