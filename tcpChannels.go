package tcpChannels

import "net"

type Receiver[T any] struct {
	Channel    chan T
	listenAddr string
	listener net.Listener
}

func NewReceiver[T any](listenAddr string) (*Receiver[T], error) {

	recv := &Receiver[T]{
		listenAddr: listenAddr,
		Channel:    make(chan T),
	}

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err;
	}

	recv.listener = listener;

	return recv, nil
}