package internal

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/ChrisPJohnstone/go-rss-client/v2"
	_ "github.com/mattn/go-sqlite3"
)

type Article struct {
	ID          int
	FeedURL     string
	GUID        string
	Title       string
	Description string
	Link        string
	Author      string
	PubDate     string
	Read        bool
	CreatedAt   string
}

func dbPath() string {
	return filepath.Join(ToolDir, dbFilename)
}

func InitDB() error {
	path := dbPath()
	ensureDir(ToolDir)
	if Verbose {
		log.Println("Connecting to DB ", path)
	}
	var err error
	Db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	if Verbose {
		log.Println("Ensuring artciles table exists")
	}
	_, err = Db.Exec(articlesSchema)
	return err
}

func InsertArticle(feedURL string, item rssclient.Item) error {
	_, err := Db.Exec(
		articleInsertQuery,
		feedURL,
		item.ID,
		item.Title,
		item.Description,
		item.Link,
		item.Author,
		item.PubDate,
	)
	return err
}
