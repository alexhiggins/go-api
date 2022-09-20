package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/alexhiggins/go-api/internal/author"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
)

type config struct {
	port string
	db   struct {
		dsn string
	}
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8080", "API api port")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres-dsn", "PostgresSQL DSN")
	flag.Parse()

	router := gin.Default()
	router.Use(cors.Default())

	logger, _ := zap.NewProduction()

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		_ = fmt.Errorf("unable to connect to database: %v", err)
		os.Exit(1)
	}

	defer func() {
		_ = db.Close()
	}()

	server := &server{
		logger: logger,
		router: router,
		commands: commands{
			createAuthor: author.NewCreateAuthorCommand(author.NewDbRepository(db)),
		},
		queries: queries{
			fetchAuthor:     author.NewFetchAuthorQuery(author.NewDbRepository(db)),
			fetchAllAuthors: author.NewFetchAllAuthorsQuery(author.NewDbRepository(db)),
		},
	}

	if err := server.run(cfg.port); err != nil {
		_ = fmt.Errorf("error starting up: %s", err)
		os.Exit(1)
	}
}
