package redirector

import (
	"fmt"
	"strings"

	"github.com/aofei/air"
)

// HTTP2HTTPSGasConfig is a set of configurations for the `HTTP2HTTPSGas()`.
type HTTP2HTTPSGasConfig struct {
}

// HTTP2HTTPSGas returns an `air.Gas` that is used to redirect HTTP requests to
// HTTPS.
func HTTP2HTTPSGas(h2hsgc HTTP2HTTPSGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if req.Scheme == "http" {
				res.Status = 301
				return res.Redirect(fmt.Sprintf(
					"https://%s%s",
					req.Authority,
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}

// HTTPS2HTTPGasConfig is a set of configurations for the `HTTPS2HTTPGas()`.
type HTTPS2HTTPGasConfig struct {
}

// HTTPS2HTTPGas returns an `air.Gas` that is used to redirect HTTPS requests
// to HTTP.
func HTTPS2HTTPGas(hs2h HTTP2HTTPSGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if req.Scheme == "https" {
				res.Status = 301
				return res.Redirect(fmt.Sprintf(
					"http://%s%s",
					req.Authority,
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}

// WWW2NonWWWGasConfig is a set of configurations for the `WWW2NonWWWGas()`.
type WWW2NonWWWGasConfig struct {
}

// WWW2NonWWWGas returns an `air.Gas` that is used to redirect www requests to
// non-www.
func WWW2NonWWWGas(w2nwgc WWW2NonWWWGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if strings.HasPrefix(req.Authority, "www.") {
				res.Status = 301
				return res.Redirect(fmt.Sprintf(
					"%s://%s%s",
					req.Scheme,
					req.Authority[4:],
					req.Path,
				))
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
			if !strings.HasPrefix(req.Authority, "www.") {
				res.Status = 301
				return res.Redirect(fmt.Sprintf(
					"%s://www.%s%s",
					req.Scheme,
					req.Authority,
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}
