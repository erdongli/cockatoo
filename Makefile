genproto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/*.proto

clean:
	rm api/*.go 

bdg:
	go run cmd/bdg/bdg.go

client:
	go run cmd/client/client.go