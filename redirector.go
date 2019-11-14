package redirector

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/aofei/air"
	"golang.org/x/net/idna"
)

// WWW2NonWWWGasConfig is a set of configurations for the `WWW2NonWWWGas`.
type WWW2NonWWWGasConfig struct {
	HTTPSEnforced bool

	Skippable func(*air.Request, *air.Response) bool
}

// WWW2NonWWWGas returns an `air.Gas` that is used to redirect www requests to
// non-www.
func WWW2NonWWWGas(w2nwgc WWW2NonWWWGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if w2nwgc.Skippable != nil &&
				w2nwgc.Skippable(req, res) {
				return next(req, res)
			}

			if strings.HasPrefix(
				strings.ToLower(req.Authority),
				"www.",
			) {
				res.Status = http.StatusMovedPermanently

				scheme := req.Scheme
				if w2nwgc.HTTPSEnforced {
					scheme = "https"
				}

				return res.Redirect(fmt.Sprintf(
					"%s://%s%s",
					scheme,
					req.Authority[4:],
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}

// NonWWW2WWWGasConfig is a set of configurations for the `NonWWW2WWWGas`.
type NonWWW2WWWGasConfig struct {
	HTTPSEnforced bool

	Skippable func(*air.Request, *air.Response) bool
}

// NonWWW2WWWGas returns an `air.Gas` that is used to redirect non-www requests
// to www.
func NonWWW2WWWGas(nw2wgc NonWWW2WWWGasConfig) air.Gas {
	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if nw2wgc.Skippable != nil &&
				nw2wgc.Skippable(req, res) {
				return next(req, res)
			}

			if !strings.HasPrefix(
				strings.ToLower(req.Authority),
				"www.",
			) {
				res.Status = http.StatusMovedPermanently

				scheme := req.Scheme
				if nw2wgc.HTTPSEnforced {
					scheme = "https"
				}

				return res.Redirect(fmt.Sprintf(
					"%s://www.%s%s",
					scheme,
					req.Authority,
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}

// OneHostGasConfig is a set of configurations for the `OneHostGas`.
type OneHostGasConfig struct {
	Host          string
	HTTPSEnforced bool

	Skippable func(*air.Request, *air.Response) bool
}

// OneHostGas returns an `air.Gas` that is used to ensure that there is only one
// host.
func OneHostGas(oagc OneHostGasConfig) air.Gas {
	if h, err := idna.Lookup.ToASCII(oagc.Host); err == nil {
		oagc.Host = h
	}

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if oagc.Skippable != nil && oagc.Skippable(req, res) {
				return next(req, res)
			}

			if oagc.Host == "" {
				return next(req, res)
			}

			h, _, _ := net.SplitHostPort(req.Authority)
			if h == "" {
				h = req.Authority
			}

			h, err := idna.Lookup.ToASCII(h)
			if err != nil {
				return err
			}

			if h != oagc.Host {
				res.Status = http.StatusMovedPermanently

				scheme := req.Scheme
				if oagc.HTTPSEnforced {
					scheme = "https"
				}

				return res.Redirect(fmt.Sprintf(
					"%s://%s%s",
					scheme,
					oagc.Host,
					req.Path,
				))
			}

			return next(req, res)
		}
	}
}
