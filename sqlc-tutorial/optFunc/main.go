package main
// https://www.youtube.com/watch?v=MDy7JQN5MN4
import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	Host   string
	DBurl  string
	DBname string
	tls    bool
}

func defaultOpts() Opts {
	return Opts{
		Host:   "localhost",
		DBurl:  "postgres://demouser:demouser@localhost:5432/tutorial?sslmode=disable",
		DBname: "tutorial",
		tls:    false,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withTls)
	fmt.Println("%+v\n", s)
}
