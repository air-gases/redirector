package redirector

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/aofei/air"
)

// WWW2NonWWWGasConfig is a set of configurations for the `WWW2NonWWWGas()`.
type WWW2NonWWWGasConfig struct {
}

// WWW2NonWWWGas returns an `air.Gas` that is used to redirect www requests to
// non-www.
func WWW2NonWWWGas(w2nwgc WWW2NonWWWGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if strings.HasPrefix(req.Authority, "www.") {
				res.Status = http.StatusMovedPermanently
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
				res.Status = http.StatusMovedPermanently
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

// OneHostGasConfig is a set of configurations for the `OneHostGas()`.
type OneHostGasConfig struct {
	Host string
}

// OneHostGas returns an `air.Gas` that is used to ensure that there is only one
// host.
func OneHostGas(oagc OneHostGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			host := oagc.Host
			if host == "" {
				if len(req.Air.HostWhitelist) == 0 {
					return next(req, res)
				}

				host = req.Air.HostWhitelist[0]
			}

			hn, _, err := net.SplitHostPort(req.Authority)
			if err != nil {
				hn = req.Authority
			}

			if hn != host {
				res.Status = http.StatusMovedPermanently
				return res.Redirect(fmt.Sprintf(
					"%s://%s%s",
					req.Scheme,
					host,
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}
