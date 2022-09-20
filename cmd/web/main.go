package main

import (
	"database/sql"
	"flag"
	"github.com/alexhiggins/go-api/internal/command"
	"github.com/alexhiggins/go-api/internal/query"
	"github.com/alexhiggins/go-api/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
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

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(cors.Default())

	logger, _ := zap.NewProduction()

	server := &server{
		config: cfg,
		logger: logger,
		router: router,
		commands: commands{
			createAuthor: command.NewCreateAuthorCommand(repository.NewAuthorRepository(db)),
		},
		queries: queries{
			fetchAuthor:     query.NewFetchAuthorQuery(repository.NewAuthorRepository(db)),
			fetchAllAuthors: query.NewFetchAllAuthorsQuery(repository.NewAuthorRepository(db)),
		},
	}

	if err := server.run(cfg.port); err != nil {
		panic(err)
	}
}
