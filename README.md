# Redirector

A useful gas that used to redirect unintended requests for the web applications
built using [Air](https://github.com/sheng/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/redirector
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.8.

## Usage

The following application will redirect all www requests to non-www:

```go
package main

import (
	"github.com/air-gases/redirector"
	"github.com/sheng/air"
)

func main() {
	air.Gases = []air.Gas{redirector.WWW2NonWWWGas}
	air.GET("/", func(req *air.Request, res *air.Response) error {
		return res.String("Absolutely non-www.")
	})
	air.Serve()
}
```

The following application will redirect all non-www requests to www:

```go
package main

import (
	"github.com/air-gases/redirector"
	"github.com/sheng/air"
)

func main() {
	air.Gases = []air.Gas{redirector.NonWWW2WWWGas}
	air.GET("/", func(req *air.Request, res *air.Response) error {
		return res.String("Absolutely www.")
	})
	air.Serve()
}
```

## Community

If you want to discuss this gas, or ask questions about it, simply post
questions or ideas [here](https://github.com/air-gases/redirector/issues).

## Contributing

If you want to help build this gas, simply follow
[this](https://github.com/air-gases/redirector/wiki/Contributing) to send pull
requests [here](https://github.com/air-gases/redirector/pulls).

## License

This gas is licensed under the Unlicense.

License can be found [here](LICENSE).
