package feed

import "github.com/ChrisPJohnstone/go-rss-cli/internal"

func hasFeed(url string) bool {
	for _, f := range internal.Cfg.Feeds {
		if f == url {
			return true
		}
	}
	return false
}
