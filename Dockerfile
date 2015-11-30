# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

RUN go get github.com/google/go-github/github
RUN go get golang.org/x/oauth2
RUN go get github.com/lib/pq

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/awsmsrc/gophergalasubs

# Copy the local package files to the container's workspace.
WORKDIR /go/src/github.com/awsmsrc/gophergalasubs

RUN go build .

CMD ["./gophergalasubs"]

