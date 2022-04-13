gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client

clean:
	rm -rf server/pb/

run: 
	go run server/main.go

