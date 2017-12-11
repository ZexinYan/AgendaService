package server

import (
	"args"
	"fmt"

	"github.com/urfave/negroni"
)

func serveURL() string {
	return fmt.Sprintf("%s:%d", *args.Host, *args.Port)
}

// Start server
func Start() {
	n := negroni.Classic()
	n.UseHandler(router())
	n.Run(serveURL())
}
