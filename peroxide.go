package peroxide

type TestingT interface {
	Fatal(args ...interface{})
}

type Conn interface {
	Close()
}

type Listener interface {
	AcceptOne() <-chan *Conn
	Close()
}