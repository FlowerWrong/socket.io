package parser

// https://github.com/socketio/socket.io-parser/blob/master/index.js
// https://github.com/socketio/socket.io-protocol

const (
	// Protocol version 4
	Protocol = 4
)

const (
	CONNECT = iota
	DISCONNECT
	EVENT
	ACK
	ERROR
	BINARY_EVENT
	BINARY_ACK
)

type Encoder struct{}
type Decoder struct{}
