package domain

type Context interface {
	Session() *Session
	Logger() Logger
}
