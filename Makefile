SEMANTIC_VER=0.7.14+
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

test:
	# coverpkg means what test dependencies to include in coverage
	${GO_CMD} test -short -v -coverpkg=./... -coverprofile=profile.cov ./pkg/... \
		&& ${GO_CMD} tool cover -func=profile.cov

.PHONY: build
