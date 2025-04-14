package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	pl := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	pl(u.Scheme)

	pl(u.User)
	pl(u.User.Username())
	p, _ := u.User.Password()
	pl(p)

	pl(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	pl(host)
	pl(port)

	pl(u.Path)
	pl(u.Fragment)

	pl(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	pl(m)
	pl(m["k"][0])

}
