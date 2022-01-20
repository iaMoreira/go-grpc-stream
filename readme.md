# Aplicação Go com GRPC

## execução

Subindo o cliente e o servidor Go com Grpc:

```shell
docker-compose up --build
```

Executando o cliente:

````shell
docker exec -it go-grpc-client  go run cmd/client/client.go
````

## Alterando o código

```go
// cmd/client/client.go 

func main () {
	...
    
    // Liberar as execuções dos métodos para chamar o servidor 
    // 1- AddUser(client)
	// 2- AddUserVerbose(client)
	// 3- AddUsers(client)
	// 4- AddUserStreamBoth(client)
}
```


## Alterando os arquivos protos

Acrescentar ao `Dockerfile` as seguintes linhas:

```docker
RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN apt-get update && apt-get install golang-goprotobuf-dev -y

RUN protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```
