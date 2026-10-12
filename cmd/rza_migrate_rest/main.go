package main

import (
	"os"

	internalcmd "rza_migrate_db/internal/cmd"
)

func main() {
	os.Exit(internalcmd.RunREST(os.Args))
}
