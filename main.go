package main

import (
"fmt"
"github.com/SirAlvarex/go_makefile/cmd"
"os"
)

func main() {
if err := cmd.RootCmd.Execute(); err != nil {
fmt.Println(err)
os.Exit(-1)
}
}
