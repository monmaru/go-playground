package mockingwithif

import (
	"io"
	"strings"
	"testing"
)

// Doerのインターフェースをみたすダミーのstruct
type dummyDoer struct{}

func (d *dummyDoer) doGistsRequest(user string) (io.Reader, error) {
	return strings.NewReader(`[
			{"html_url": "https://gist.github.com/example1"},
			{"html_url": "https://gist.github.com/example2"}
		]`), nil
}

func TestGetGists2(t *testing.T) {
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")

	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
