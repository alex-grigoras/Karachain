package main

import (

"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)

const(
VideoTable = "Videos"
RatingTable = "Ratings"
UserTable = "Users"
SingerTable = "Singer"
ProducerTable = "Producer"
ContractTable = "Contract"
)
type Karachain struct{}

func main() {

	err := shim.Start(new(Karachain))
	if err != nil {
		fmt.Println("Error starting Karachain: %s", err)
	}
}

func (t *Karachain) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	return nil, nil
}

func (t *Karachain) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Karachain) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	shim.LogInfo("Hello " + args[0] + "!")
	return nil, nil
}
