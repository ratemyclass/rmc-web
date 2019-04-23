package rmcdb

// RMCDB is an interface for the methods defined on the RateMyClass database
// so we can possibly switch if there's an issue with one or want it to be
// more powerful / complex. In other words, we're not bound to one database
// implementation and can customize as we please
type RMCDB interface {
	// SetupClient makes sure that whatever things need to be done before we use the datastore can be done before we need to use the datastore.
	SetupClient() error

	// empty for now
}
