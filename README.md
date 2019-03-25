go-enex
=======

[![GoDoc](https://godoc.org/github.com/macrat/go-enex?status.svg)](https://godoc.org/github.com/macrat/go-enex)

The parser for [the xml that exported by Evernote (.enex)](https://help.evernote.com/hc/en-us/articles/209005557).

## Example

``` go
package main

import (
	"os"
	
	"github.com/macrat/go-enex"
)

func main() {
	f, _ := os.Open("notebook.enex")

	data, err := enex.ParseFromReader(f)
	if err != nil {
		panic(err.Error())
	}

	for _, note := range data.Notes {
		fmt.Println(note.UpdatedAt, ": ", note.Title)
	}
}
```
