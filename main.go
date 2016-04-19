package main

import (
"fmt"
"github.int.yammer.com/docker/go_makefile/cmd"
"os"
)

func main() {
if err := cmd.RootCmd.Execute(); err != nil {
fmt.Println(err)
os.Exit(-1)
}
}
