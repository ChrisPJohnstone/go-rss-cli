# RSS CLI

An RSS Reader for the CLI

CLI powered by [cobra](https://github.com/spf13/cobra)

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
