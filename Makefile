SEMANTIC_VER=0.3.17+
BUILD_VER=$(shell git describe --always --long)
PRE_RELEASE_VER=alpha

VERSION=${SEMANTIC_VER}${BUILD_VER}-${PRE_RELEASE_VER}
NAME=io-sdk-golang
PKG=github.com/clearchanneloutdoor/${NAME}
BLD_DST=./
BLD_FLGS=-v -a -tags netgo -ldflags "\
	-w -extldflags \"-static\" \
	-X ${PKG}/pkg/api.Version=$(VERSION) \
	"
BNRY_NM=io
CGO_ENABLED=0
DST=${BLD_DST}${BNRY_NM}
GO_CMD=go

build: download
	${GO_CMD} build ${BLD_FLGS} -o ${DST} ./cmd/...

download:
	${GO_CMD} mod download

.PHONY: build
