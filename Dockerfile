FROM golang:1.17

WORKDIR /go/grpc

COPY . .

# RUN export PATH="$PATH:$(go env GOPATH)/bin"

# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# RUN apt-get update && apt-get install golang-goprotobuf-dev -y

# RUN protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

EXPOSE 50051

CMD [ "go", "run", "cmd/server/server.go"]