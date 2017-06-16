go-ask [![][travis-badge]][travis-link] [![][godoc-badge]][godoc-link]
======

Just a few functions for helping questions and user's input.

`go get` it as `github.com/b4b4r07/go-ask`, import it as `"github.com/b4b4r07/go-ask"`, use it as `ask`.

See [godoc][godoc-link] for complete documentation.

## Example

```console
$ go get github.com/b4b4r07/go-ask
```

```go
package main

import (
	"fmt"

	"github.com/b4b4r07/go-ask"
)

func main() {
	if ask.NewQ().Confirm("Do you want to remove this?") {
		fmt.Println("done")
	}
}
```

## License

MIT

## Author

b4b4r07

[travis-badge]: https://travis-ci.org/b4b4r07/go-ask.svg?branch=master
[travis-link]: https://travis-ci.org/b4b4r07/go-ask

[godoc-badge]: https://godoc.org/github.com/b4b4r07/go-ask?status.svg
[godoc-link]: http://godoc.org/github.com/b4b4r07/go-ask
