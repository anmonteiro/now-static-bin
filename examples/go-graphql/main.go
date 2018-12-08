package main

import (
        "log"
        "net/http"

        graphql "github.com/graph-gophers/graphql-go"
        "github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func main() {
        s := `
                schema {
                        query: Query
                }
                type Query {
                        hello: String!
                }
        `
        schema := graphql.MustParseSchema(s, &query{})
        http.Handle("/", &relay.Handler{Schema: schema})
        log.Fatal(http.ListenAndServe(":8080", nil))
}