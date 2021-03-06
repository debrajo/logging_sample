/*
*  copied from  https://github.com/hyperledger/fabric-samples/tree/master/chaincode/sacc 
*  and modified to performing logging .
*
*/

package main

import (
    "fmt"
    "os"
    "github.com/op/go-logging"	
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

//var logger = shim.NewLogger("debrajo")
var logger = logging.MustGetLogger("debrajo")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
    fmt.Printf("logging_sample.init started")
	
    logger.Critical("logging_sample.init critical")
    logger.Warning("logging_sample.init warning")
    logger.Info("logging_sample.init info")
    logger.Debug("logging_sample.init debug")
	
    // Get the args from the transaction proposal
    args := stub.GetStringArgs()
    if len(args) != 2 {
            return shim.Error("Incorrect arguments. Expecting a key and a value")
    }

    // Set up any variables or assets here by calling stub.PutState()

    // We store the key and the value on the ledger
    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
    }
    fmt.Printf("logging_sample.init ended")
    return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    fmt.Printf("logging_sample.invoke started")

    logger.Critical("logging_sample.invoke critical")
    logger.Warning("logging_sample.invoke warning")
    logger.Info("logging_sample.invoke info")
    logger.Debug("logging_sample.invoke debug")

	
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()

    var result string
    var err error
    if fn == "set" {
            result, err = set(stub, args)
    } else { // assume 'get' even if fn is nil
            result, err = get(stub, args)
    }
    if err != nil {
            return shim.Error(err.Error())
    }

    // Return the result as success payload
    fmt.Printf("logging_sample.invoke ended")
    return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }

    value, err := stub.GetState(args[0])
    if err != nil {
            return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
            return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	fmt.Printf("logging_sample.go.main started")
    
	
	// Set up  op/go-logging the logging format
	format := logging.MustStringFormatter("%{time:15:04:05.000} [%{module}] %{shortfile} %{level:.4s} : %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.DEBUG, "main")


	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical
	//logger.SetLevel(shim.LogDebug)
	//logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	//logLevel, _ := shim.LogLevel("DEBUG")
	//logger.Info(logLevel)
	//shim.SetLoggingLevel(logLevel)
	//logger.Info(logger.IsEnabledFor(logLevel))
    
	err := shim.Start(new(SimpleAsset))
	if  err != nil {
            fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
    	}
	
	// set the shim logging level to Info
	//shim.SetLoggingLevel(shim.LogInfo)
	
	
	logger.Info("debrajo:logging_sample.go.main ended")
	fmt.Printf("logging_sample.go.main ended")
}
