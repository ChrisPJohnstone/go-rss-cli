# RSS CLI

An RSS Reader for the CLI

## Dependencies

> [!NOTE]
> This project uses [go-sqlite3](https://github.com/mattn/go-sqlite3), which requires CGO. On native platform builds this is enabled by default — no action needed. If you're cross-compiling, ensure `CGO_ENABLED=1` is set.

- [gcc](https://gcc.gnu.org/) C Compiler
- [sqlite3](https://sqlite.org/index.html) Database Engine

## Libraries

- [cobra](https://github.com/spf13/cobra) CLI Framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3#installation) Database

## Design Scribbles

- Need some kind of persistent store
  - List of feeds which you are subscribed to
  - List of articles which you have & haven't read
- Required functionality
  - Manage feeds
    - Should probably be human readable config (e.g. YAML) so easier to have in like dotfiles
    - Should have command wrappers
  - Pull new articles
    - Stored somewhere (sqlite3, avro, plain text?)
    - Refresh triggered by command
    - People can schedule with cron
  - Show list of pulled articles
  - Open article - This tool should probably just output raw text? Can have another tool manage rendering

Something like

```sh
$ rss feed add https://feeds.bbci.co.uk/news/rss.xml
$ rss pull
$ rss ls
...
$ rss read {article}
...
```
