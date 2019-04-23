package rmcserver

import (
	"crypto/x509"

	"github.com/ratemyclass/rmc-web/rmcdb"
)

// RMCServer does things like handle requests, interact with the database, basically integrate all the parts of the ratemyclass server
type RMCServer struct {
	RMCDatastore rmcdb.RMCDB

	certificate *x509.Certificate
}

// InitServer creates a server with options
func InitServer(db rmcdb.RMCDB) (server *RMCServer) {

	server = &RMCServer{
		RMCDatastore: db,
	}

	return
}
