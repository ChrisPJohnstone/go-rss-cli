# RSS CLI

An RSS Reader for the CLI

## Installation

> [!NOTE]
> This project uses [go-sqlite3](https://github.com/mattn/go-sqlite3), which requires CGO. On native platform builds this is enabled by default — no action needed. If you're cross-compiling, ensure `CGO_ENABLED=1` is set.

### Dependencies

- [gcc](https://gcc.gnu.org/) C Compiler
- [sqlite3](https://sqlite.org/index.html) Database Engine

### Build

```sh
go install github.com/ChrisPJohnstone/go-rss-cli@latest
```

## Usage

```sh
# Follow a new feed
rss feed add https://feeds.bbci.co.uk/news/rss.xml

# List followed feeds
rss feed ls

# Unfollow a feed
rss feed rm https://feeds.bbci.co.uk/news/rss.xml

# Pull latest articles from all followed feeds
rss pull
```

TODO:

- Add support for listing articles (something like `rss ls` shows a list of ID - title)
- Add support for getting articles (something like `rss show {id}` & output to stdout)
