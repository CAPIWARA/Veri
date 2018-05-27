package resolvers

import (
	"Veri/Server/eth"
	"fmt"
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

type Account struct {
	Balance    string `json:"balance"`
	Trajectory string `json:"trajectory"`
}

//GetContract is
func GetContract(params graphql.ResolveParams) (interface{}, error) {
	//var contract Contract
	uuid := params.Args["uuid"].(string)
	balance, err := eth.Contract.ConsultByUuid(nil, uuid)

	if err != nil {
		return nil, err
	}

	account := &Account{
		Balance:    fmt.Sprintf("%v", balance.Balance),
		Trajectory: fmt.Sprintf("%v", balance.Trajectory),
	}
	fmt.Printf("aCCOUNT: %#V \n", account)
	return account, nil
}

//CreateTrajectory is
func CreateTrajectory(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)
	trajeto := params.Args["trajeto"].(string)
	km := big.NewInt(3)
	pontos := big.NewInt(3)
	var res Contract
	res = Contract{
		UUID: uuid,
	}

	eth.Contract.RegistryTrajectory(eth.Auth, pontos, trajeto, uuid, km)
	eth.Sim.Commit()

	return res, nil
}
