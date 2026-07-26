package internal

import "database/sql"

var (
	Verbose bool
	ToolDir string
	Cfg     Config
	Db      *sql.DB
)
