package domain

type Session struct {
	UserId          string
	ClientIds       []string
	TraceId         string
	InternalRequest string
}
