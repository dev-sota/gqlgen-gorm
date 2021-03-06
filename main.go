package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	dataloader "github.com/dev-sota/gqlgen-gorm/dataloader/graph"
	"github.com/dev-sota/gqlgen-gorm/graph"
	"github.com/dev-sota/gqlgen-gorm/graph/generated"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dsn  = "root:@tcp(127.0.0.1:3306)/gqlgen?charset=utf8&parseTime=True&loc=Local"
	port = "8080"
)

func main() {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle(
		"/query",
		dataloader.Middleware(
			db,
			handler.NewDefaultServer(
				generated.NewExecutableSchema(
					generated.Config{
						Resolvers: &graph.Resolver{
							DB: db,
						},
					},
				),
			),
		),
	)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
