# Makefile
# lijiaocn, 2019-05-05 16:56
#

VERSION=1.0.0
COMPILE=$(shell date -u "+%Y-%m-%d/%H:%M:%S")

all: build
build:
	go build -ldflags "-X github.com/lijiaocn/golib/version.VERSION=${VERSION} -X github.com/lijiaocn/golib/version.COMPILE=${COMPILE}"
