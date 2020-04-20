package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func main() {
	sh = shell.NewShell("https://ipfs.infura.io:5001")

	fmt.Println("MMMMMMMMMMMMMMWX0kdlldk0XWMMMMMMMMMMMMMM\nMMMMMMMMMMWNKOxdoddooddodxOKNWMMMMMMMMMM\nMMMMMMMWX0kxdoddddddddddddodxk0XWMMMMMMM\nMMMWNKOxdoddddddddddddddddddddodxOKNWMMM\nMWKkdooddddddddddddddddddddddddddoodkKWM\nMNxcllloddddddddddddddddddddddddolllcxNM\nMNxooollllooddddddddddddddddoolllloooxNM\nMNxoddddoolllloodddddddddollllooddddoxNM\nMNxoddddddddolllllllllllllloddddddddoxNM\nMNxodddddddddddolcccccclodddddddddddoxNM\nMNxoddddddddddddolccccloddddddddddddoxNM\nMNxoddddddddddddddlcclddddddddddddddoxNM\nMNxoddddddddddddddoccoddddddddddddddoxNM\nMNxlodddddddddddddoccodddddddddddddolxNM\nMNklloddddddddddddoccoddddddddddddollkNM\nMMWX0kddddddddddddoccoddddddddddddk0XNMM\nMMMMMWN0OxdoddddddoccoddddddodxOKNWMMMMM\nMMMMMMMMWNXOkdodddoccodddodkOXNMMMMMMMMM\nMMMMMMMMMMMMWX0OxolccloxO0XWMMMMMMMMMMMM\nMMMMMMMMMMMMMMMWN0dlld0NWMMMMMMMMMMMMMMM\n")

	fmt.Println("### ######  #######  #####  \n #  #     # #       #     # \n #  #     # #       #       \n #  ######  #####    #####  \n #  #       #             # \n #  #       #       #     # \n### #       #        #####  \n")

	fmt.Println("###########################\n   Welcome to IPLD-CRUD!\n###########################\n")
	fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!\n")

	keyValueMap := make(map[string]interface{})

	fmt.Println("Enter value for the key field: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputKey := scanner.Text()

	fmt.Println("Enter value for the value field: ")
	scanner.Scan()
	inputValue := scanner.Text()

	keyValueMap[inputKey] = inputValue

	entryJSON, err := json.Marshal(keyValueMap)
	if err != nil {
		fmt.Printf("There was an error marshalling your data %v", err)
	}

	jsonStr := string(entryJSON)
	fmt.Println("The JSON object of your key-value entry is:")
	fmt.Println(jsonStr)

	cid, err := sh.DagPut(entryJSON, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Println("-------\nOUTPUT\n-------")
	fmt.Printf("Successfully WRITE added %s. Here is the IPLD Explorer link: https://explore.ipld.io/#/explore/%s\n", string(cid+"\n"), string(cid+"\n"))

	fmt.Printf("READ: Value for key \"%s\" is: ", inputKey)

	res, err := GetDag(cid, inputKey)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response is: %v", res)
}

// GetDag reads a DAG with a CID and returns the corresponding value
func GetDag(ref, key string) (out interface{}, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}
