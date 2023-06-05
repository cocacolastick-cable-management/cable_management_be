package sse

type IServerSentEventProvider interface {
	SendMessage(clientId, message string)
}
