package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jose-manaloto/joseapp/graph"
	"github.com/jose-manaloto/joseapp/graph/generated"
	"github.com/jose-manaloto/joseapp/graph/model"
)

func initDB() {
    var err error
    dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
    db, err = gorm.Open("mysql", dataSourceName)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }

    db.LogMode(true)

    // Create the database. This is a one-time step.
    // Comment out if running multiple times - You may see an error otherwise
    db.Exec("CREATE DATABASE test_db")
    db.Exec("USE test_db")

    // Migration to create tables for Order and Item schema
    db.AutoMigrate(&models.House{}, &models.Issue{})
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    initDB()
    http.Handle("/", handler.Playground("GraphQL playground", "/query"))
    http.Handle("/query", handler.GraphQL(go_house_issues_api.NewExecutableSchema(go_house_issues_api.Config{Resolvers: &go_house_issues_api.Resolver{
        DB: db,
    }})))

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
