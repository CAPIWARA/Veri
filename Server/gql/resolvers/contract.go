package resolvers

import (
	"Veri/Server/eth"
	"math/big"
	"time"

	"github.com/graphql-go/graphql"
)

//Contract is
type Contract struct {
	UUID     string    `json:"uuid"`
	Trajeto  string    `json:"trajeto"`
	Km       int       `json:"km"`
	Pontos   int       `json:"pontos"`
	UpdateAt time.Time `json:"updateAt"`
}

//GetContract is
func GetContract(params graphql.ResolveParams) (interface{}, error) {
	//var contract Contract
	uuid := params.Args["uuid"].(string)
	eth.Contract.ConsultByUuid(nil, uuid)
	//res, _ := eth.Contract.ConsultarEmail(nil, key)
	return nil, nil
}

//CreateTrajectory is
func CreateTrajectory(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)
	trajeto := params.Args["trajeto"].(string)
	km := big.NewInt(3)
	pontos := big.NewInt(3)
	var res Contract
	res = Contract{
		UUID:    uuid,
		Trajeto: trajeto,
		Km:      int.new(km),
		Pontos:  pontos,
	}

	eth.Contract.RegistryTrajectory(eth.Auth, pontos, trajeto, uuid, km)
	eth.Sim.Commit()

	return res, nil
}
