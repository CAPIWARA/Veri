package eth

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var key, _ = crypto.GenerateKey()

//Auth is var
var Auth = bind.NewKeyedTransactor(key)

//Sim start a testnet dev
var Sim *backends.SimulatedBackend

//Contract document register
var Contract *DocumentRegister

//InitEth config eth settings
func InitEth() {
	alloc := make(core.GenesisAlloc)
	alloc[Auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	Sim = backends.NewSimulatedBackend(alloc)

	addr, _, contract, err := DeployDocumentRegister(Auth, Sim)
	fmt.Printf("contract deployed to %s\n", addr.String())
	if err != nil {
		log.Fatalf("could not deploy contract; %v ", err)
	}
	log.Printf("contract contains %v documents\n", contract)
	Contract = contract
	Sim.Commit()

	res, _ := Contract.ConsultByUuid(nil, "333")
	fmt.Println("tamanho do contrato:", fmt.Sprintf("%s", res))

	Contract.RegistryTrajectory(Auth, big.NewInt(2), "18288283.33323.13-12313.333.3", "333", big.NewInt(3))
	Sim.Commit()
}
