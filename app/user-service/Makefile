build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/mongmx/mymicro/app/user-service \
		proto/user/user.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

	# GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo