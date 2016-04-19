
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=go_makefile

VERSION=$(shell git describe --always --long)
BUILD_TIME=$(shell date +%FT%T%z)

# Presumes we are using Viper/Cobra for CLI commands.  Place "version" and "buildDate" variable in your cmd/root.go file to enable populating of version flags
LDFLAGS=-ldflags "-X github.int.yammer.com/docker/go_makefile/cmd.version=${VERSION} -X github.int.yammer.com/docker/go_makefile/cmd.buildDate=${BUILD_TIME}"

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
