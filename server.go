package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jose-manaloto/joseapp/graph"
	"github.com/jose-manaloto/joseapp/graph/generated"
	"github.com/jose-manaloto/joseapp/graph/model"
)

var db *gorm.DB;
var defaultPort = "8000"

func initDB() {
    var err error
    dataSourceName := "root:root@tcp(localhost:3306)/?parseTime=True"
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
    db.AutoMigrate(&model.House{}, &model.Issue{})
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    initDB()
    http.Handle("/", handler.Playground("GraphQL playground", "/query"))
    http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
        DB: db,
    }})))

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
