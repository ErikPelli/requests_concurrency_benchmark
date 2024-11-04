package client

type echoRequest interface {
	EchoRequest() error
}
