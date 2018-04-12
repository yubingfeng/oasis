/*
Copyright hill  2018 All Rights Reserved

*/


package main

import (
        "fmt"
        "strconv"

        "github.com/hyperledger/fabric/core/chaincode/shim"
        pb "github.com/hyperledger/fabric/protos/peer"

)

var logger =  shim.NewLogger("oasis")

//Oasis chain code

type OasisChaincode struct{

}

func  (t *OasisChaincode) Init(stub shim.ChaincodeInterface) pb.Response{

        logger.Info("####### Oasis Chaincode Init #######")
        

}
