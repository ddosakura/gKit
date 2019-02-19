package switcher

import (
	"errors"
	"net/http"

	"github.com/rakyll/statik/fs"
)

// Loader id the `init` in statik.go
type Loader func()

var (
	// Errors
	ErrNsNotFound = errors.New("Namespace doesn't loaded")

	// Data
	nss = make(map[string]Loader)
)

// Register binding the loader to the namespace
func Register(ns string, loader Loader) {
	nss[ns] = loader
}

// New build the target fs by namespace
func New(ns string) (http.FileSystem, error) {
	fn := nss[ns]
	if fn == nil {
		return nil, ErrNsNotFound
	}
	fn()
	return fs.New()
}
