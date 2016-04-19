package makefile


import (
	"path/filepath"
	"os"
	"log"
	"html/template"
	"bytes"
	"os/exec"
)

var makeFileTemplate = `
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY={{ .Binary }}

VERSION=$(shell git describe --always --long)
BUILD_TIME=$(shell date +%FT%T%z)

# Presumes we are using Viper/Cobra for CLI commands.  Place "version" and "buildDate" variable in your cmd/root.go file to enable populating of version flags
LDFLAGS=-ldflags "-X {{ .GoPackage }}/cmd.version=${VERSION} -X {{ .GoPackage }}/cmd.buildDate=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)


$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY}

.PHONY: linux
linux: $(SOURCES)
	env GOOS=linux GOARCH=amd64  go build ${LDFLAGS} -o ${BINARY}

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
`

func BuildMakefile(binaryName string) {
	myPath, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
		return
	}
	myPackage := filepath.Base(myPath)
	myPackageRepo := filepath.Base(filepath.Dir(myPath))
	myPackageRepoHost := filepath.Base(filepath.Dir(filepath.Dir(myPath)))
	myGoPackage := myPackageRepoHost + "/" + myPackageRepo + "/" + myPackage

	if len(binaryName) == 0 {
		log.Printf("Binary name not given, using whats in the path: %v", myPackage)
		binaryName = myPackage
	}

	makeFile, err := os.Create("Makefile")
	if err != nil {
		log.Printf("Could not open Makefile for writing: %v", err)
		return
	}
	defer makeFile.Close()

	t, err := template.New("Makefile").Parse(makeFileTemplate)
	if err != nil {
		log.Printf("Could not open Template: %v", err)
		return
	}

	templateData := bytes.NewBufferString("")

	err = t.Execute(templateData, struct { Binary string
	                                       GoPackage string }{ binaryName, myGoPackage })

	if err != nil {
		log.Printf("Could not execute Template: %v", err)
		return
	}

	_, err = makeFile.Write(templateData.Bytes())
	if err != nil {
		log.Printf("Could not write to Makefile: %v", err)
		return
	}
}

var CobraTemplate = `
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var (
	version string
	buildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of {{ .Binary }}",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("{{ .Binary }}: %s - %s\n", version, buildDate)
	},
}
`
func InitCobra(binaryName string) {
	if cmdExists, _ := exists("cmd/"); !cmdExists {
		cmd := exec.Command("cobra", "init")
		err := cmd.Run()
		if err != nil {
			log.Printf("Could not init cobra: %v", err)
			return
		}
	}
	myPath, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
		return
	}
	myPackage := filepath.Base(myPath)

	if len(binaryName) == 0 {
		log.Printf("Binary name not given, using whats in the path: %v", myPackage)
		binaryName = myPackage
	}

	makeFile, err := os.Create("cmd/version.go")
	if err != nil {
		log.Printf("Could not open cmd/version.go for writing: %v", err)
		return
	}
	defer makeFile.Close()

	t, err := template.New("cmd/version.go").Parse(CobraTemplate)
	if err != nil {
		log.Printf("Could not open Template: %v", err)
		return
	}

	templateData := bytes.NewBufferString("")

	err = t.Execute(templateData, struct { Binary string }{ binaryName })

	if err != nil {
		log.Printf("Could not execute Template: %v", err)
		return
	}

	_, err = makeFile.Write(templateData.Bytes())
	if err != nil {
		log.Printf("Could not write to cmd.version.go: %v", err)
		return
	}

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}