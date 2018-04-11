build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/mongmx/mymicro/app/vessel-service \
		proto/vessel/vessel.proto

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo
	# docker build -t vessel-service .
	docker build -t asia.gcr.io/micro-umx/vessel-service:latest .
	docker push asia.gcr.io/micro-umx/vessel-service:latest

run:
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		vessel-service

push-image:
	docker build -t asia.gcr.io/micro-umx/vessel-service:latest .
	docker push asia.gcr.io/micro-umx/vessel-service:latest

deploy:
	# gen deployment yaml file from template
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml