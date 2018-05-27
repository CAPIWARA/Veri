package main

import (
	"Veri/Server/eth"
	"Veri/Server/gql"
	"log"
	"net/http"
	"runtime"

	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	eth.InitEth()
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gql.Schema,
	})

	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
	// res, _ = eth.Contract.ConsultarTamanho(nil)
	// fmt.Println("tamanho do contrato:", fmt.Sprintf("%s", res))

	// email, _ := eth.Contract.ConsultarEmail(nil, "222")
	// fmt.Println("finding by email:", fmt.Sprintf("%s", email))

}
