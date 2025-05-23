# DSN Generator

Generate correct connection strings for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

## Usage

```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/simonward87/dsn"
)

func main() {
	db, err := sql.Open("sqlite3", dsn.SQLite("db.sqlite3", dsn.Options{
		BusyTimeout: 10_000,
		ForeignKeys: true,
		JournalMode: dsn.JournalModeWal,
		Synchronous: dsn.SynchronousNormal,
	}))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: database stuff
}
```
