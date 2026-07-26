package internal

const ToolDirArg string = "tool-dir"
const configFilename string = "config.yaml"
const dbFilename string = "articles.db"
const articlesSchema string = `
CREATE TABLE IF NOT EXISTS articles (
	"id" 			INTEGER 	PRIMARY KEY AUTOINCREMENT,
	"feed_url" 		TEXT		NOT NULL,
	"guid" 			TEXT		NOT NULL,
	"title" 		TEXT		NULL,
	"description"	TEXT		NULL,
	"link" 			TEXT		NULL,
	"author" 		TEXT		NULL,
	"pub_date" 		TEXT		NULL,
	"read" 			INTEGER		NOT NULL DEFAULT 0,
	"created_at"	DATETIME	NOT NULL DEFAULT CURRENT_TIMESTAMP,
	UNIQUE("feed_url", "guid")
);`
const articleInsertQuery = `
INSERT OR IGNORE INTO "articles" (
	"feed_url",
	"guid",
	"title",
	"description",
	"link",
	"author",
	"pub_date"
) VALUES (?, ?, ?, ?, ?, ?, ?)
`
