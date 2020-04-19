package main

shell "github.com/ipfs/go-ipfs-api"

var sh *shell.Shell

func main() {
	sh = shell.NewShell("https://ipfs.infura.io:5001")
}
