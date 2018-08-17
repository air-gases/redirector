package redirector

import (
	"strings"

	"github.com/sheng/air"
)

// WWW2NonWWWGas is used to redirect www requests to non-www.
func WWW2NonWWWGas(next air.Handler) air.Handler {
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

// NonWWW2WWWGas is used to redirect non-www requests to www.
func NonWWW2WWWGas(next air.Handler) air.Handler {
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
