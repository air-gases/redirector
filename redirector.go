package redirector

import (
	"strings"

	"github.com/sheng/air"
)

// WWW2NonWWWGasConfig is a set of configurations for the `WWW2NonWWWGas()`.
type WWW2NonWWWGasConfig struct {
}

// WWW2NonWWWGas returns an `air.Gas` that is used to redirect www requests to
// non-www.
func WWW2NonWWWGas(w2nwgc WWW2NonWWWGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if strings.HasPrefix(req.URL.Host, "www.") {
				res.StatusCode = 301
				u := *req.URL
				u.Host = u.Host[4:]
				return res.Redirect(u.String())
			}

			return next(req, res)
		}
	}
}

// NonWWW2WWWGasConfig is a set of configurations for the `NonWWW2WWWGas()`.
type NonWWW2WWWGasConfig struct {
}

// NonWWW2WWWGas returns an `air.Gas` that is used to redirect non-www requests
// to www.
func NonWWW2WWWGas(nw2wgc WWW2NonWWWGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if !strings.HasPrefix(req.URL.Host, "www.") {
				res.StatusCode = 301
				u := *req.URL
				u.Host = "www." + u.Host
				return res.Redirect(u.String())
			}

			return next(req, res)
		}
	}
}
