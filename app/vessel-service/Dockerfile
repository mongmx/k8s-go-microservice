FROM golang:1.10.0 as builder

WORKDIR /go/src/github.com/mongmx/mymicro/app/vessel-service

COPY . .

# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init && dep ensure
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest
LABEL maintainer="mongmx@gmail.com"

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/mongmx/mymicro/app/vessel-service/vessel-service .

CMD ["./vessel-service"]