package data

// Fetcher is the common interface of a fetcher service.
type Fetcher interface {
	Run() error
	Stop() error
}
