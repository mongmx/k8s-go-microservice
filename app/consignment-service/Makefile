build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/mongmx/mymicro/app/consignment-service \
	  proto/consignment/consignment.proto

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo
	docker build -t consignment-service .

run: 
	docker run -p 50051:50051 --name consignment-service -d consignment-service
