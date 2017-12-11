package args

import (
	"github.com/ogier/pflag"
)

// Host ..
var Host *string

// Port ..
var Port *uint16

// DB is the path to the sqlite db file
var DB *string

func init() {
	Host = pflag.StringP("host", "h", "0.0.0.0", "Host of the server listening on")
	Port = pflag.Uint16P("port", "p", 8080, "Port of the server listening on")
	DB = pflag.StringP("db", "d", "agenda.db", "Path to the sqlite db file")
	pflag.Parse()
}
