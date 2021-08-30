package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/dev-sota/gqlgen-gorm/graph"
	"github.com/dev-sota/gqlgen-gorm/graph/generated"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn  = "root:@tcp(127.0.0.1:3306)/gqlgen?charset=utf8&parseTime=True&loc=Local"
	port = "8080"
)

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	srv := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
