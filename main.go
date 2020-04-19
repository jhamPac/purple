package main

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func main() {
	sh = shell.NewShell("https://ipfs.infura.io:5001")

	fmt.Println("MMMMMMMMMMMMMMWX0kdlldk0XWMMMMMMMMMMMMMM\nMMMMMMMMMMWNKOxdoddooddodxOKNWMMMMMMMMMM\nMMMMMMMWX0kxdoddddddddddddodxk0XWMMMMMMM\nMMMWNKOxdoddddddddddddddddddddodxOKNWMMM\nMWKkdooddddddddddddddddddddddddddoodkKWM\nMNxcllloddddddddddddddddddddddddolllcxNM\nMNxooollllooddddddddddddddddoolllloooxNM\nMNxoddddoolllloodddddddddollllooddddoxNM\nMNxoddddddddolllllllllllllloddddddddoxNM\nMNxodddddddddddolcccccclodddddddddddoxNM\nMNxoddddddddddddolccccloddddddddddddoxNM\nMNxoddddddddddddddlcclddddddddddddddoxNM\nMNxoddddddddddddddoccoddddddddddddddoxNM\nMNxlodddddddddddddoccodddddddddddddolxNM\nMNklloddddddddddddoccoddddddddddddollkNM\nMMWX0kddddddddddddoccoddddddddddddk0XNMM\nMMMMMWN0OxdoddddddoccoddddddodxOKNWMMMMM\nMMMMMMMMWNXOkdodddoccodddodkOXNMMMMMMMMM\nMMMMMMMMMMMMWX0OxolccloxO0XWMMMMMMMMMMMM\nMMMMMMMMMMMMMMMWN0dlld0NWMMMMMMMMMMMMMMM\n")

	fmt.Println("### ######  #######  #####  \n #  #     # #       #     # \n #  #     # #       #       \n #  ######  #####    #####  \n #  #       #             # \n #  #       #       #     # \n### #       #        #####  \n")

	fmt.Println("###########################\n   Welcome to IPLD-CRUD!\n###########################\n")
	fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!\n")
}
