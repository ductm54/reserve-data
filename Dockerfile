# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/KyberNetwork/reserve-data

WORKDIR /go/src/github.com/KyberNetwork/reserve-data
RUN go get github.com/tools/godep
RUN go get github.com/ethereum/go-ethereum/common
RUN go get github.com/ethereum/go-ethereum/common/hexutil
RUN go get github.com/ethereum/go-ethereum/crypto/sha3
RUN go get github.com/getsentry/raven-go
RUN go get github.com/gin-contrib/sentry
RUN go get github.com/gin-gonic/gin
RUN go install github.com/KyberNetwork/reserve-data/http

ENTRYPOINT ["http", "http://192.168.24.184:5000"]

EXPOSE 8000