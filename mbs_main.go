
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type loan struct {
	loanAccNo int
	loanAmt int
	loanRate int
	loanTerm int
	propertyNo string
	borrowerName string
	borrowerBSN string
	borrowerCreditRating int
}

type tranche struct {
	trancheId int
	trancheRating int
	trancheRate int	
	loans []loan //array of loans

}

type trancheRatingList struct {
	list [3]string //harcoded list of values
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var loan A1, A2 //loans 
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	//int i: = 1
	//for i<=2
	//{
	//	i = i + 1	
	//}
	
	A1 = loan{loanAccNo: 101, loanAmt: 100000, loanRate: 5, loanTerm: 10, propertyNo "HSR1001", borrowerName "R. Anndurai", borrowerBSN: "000001",	borrowerCreditRating: 8}
	A2 = loan{loanAccNo: 102, loanAmt: 200000, loanRate: 4, loanTerm: 10, propertyNo "HSR1002", borrowerName "R. Murthy", borrowerBSN: "000002",	borrowerCreditRating: 5}
	
	fmt.Printf("loanA1No = %d, loanA2No = %d\n", A1.loanAccNo, A2.loanAccNo)

	// Write the state to the ledger
	err = stub.PutState(A1, []byte(strconv.Itoa(A1.loanAccNo)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(A2, []byte(strconv.Itoa(A2.loanAccNo)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke code for the MBS blockchain
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateTranche" {
		// Deletes an entity from its state
		return t.CreateTranche(stub, args)
	}

	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) CreateTranche(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var A, B, C, Aval, Bval, Cval, retstr string    // 
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0] //Loan1
	B := args[1] //Loan2
	C := args[2] //Loan3
	
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval = string(Avalbytes))
	out := strings.split(Aval, "|")
	rating := strconv.Atoi(out[3])
	if rating >=7 then retstr = "Sr. Secured"
	
	jsonResp := "{\"Loan\":\"" + A + "\",\"Tranche\":\"" + retstr + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


	
