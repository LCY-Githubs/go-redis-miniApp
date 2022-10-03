package util

type Responses interface {
	SetCode(string)
	SetData(interface{})
	SetMsg(string)
	SetStatus(bool)
	// Clone() Responses
}
