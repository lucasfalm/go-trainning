package builderv2

type ServerOptsFn func(opts *ServerOpts)

type ServerOpts struct {
	id  string
	tls bool
}

type Server struct {
	ServerOpts
}

func DefaultOpts() *ServerOpts {
	return &ServerOpts{
		id:  "default",
		tls: false,
	}
}

// NOTE: we do not need to return the ServerOptsFn
// because the WithTLS has the same signature of it
func WithTLS(opts *ServerOpts) {
	opts.tls = true
}

// NOTE: we need to return the ServerOptsFn because
// the WithID signature is different than the ServerOptsFn
func WithID(id string) ServerOptsFn {
	return ServerOptsFn(func(opts *ServerOpts) {
		opts.id = id
	})
}

// NOTE: it accepts zero or many ServerOptsFn
func NewServer(opts ...ServerOptsFn) *Server {
	o := DefaultOpts()

	for _, optsFn := range opts {
		optsFn(o)
	}

	s := Server{ServerOpts: *o}

	return &s
}
