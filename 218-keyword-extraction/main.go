package main

import (
	"fmt"
	"net/http"
)


func main() {
	url := `https://www.stanventures.com/blog/multiple-h1-tags/`
	getSite(url)
}

func getSite(url string) (*http.Response, error) {
	res, err := http.Get(url)
	// data, _ := io.ReadAll(res.Body)
	h1s := make([]string, 20)
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			h1s = append(h1s, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(res.Body)
	fmt.Println(h1s)

	return res, err
}
