FROM golang:1.10.0 as builder

WORKDIR /go/src/github.com/mongmx/mymicro/app/user-service

COPY . .

# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init && dep ensure
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest
LABEL maintainer="mongmx@gmail.com"

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/mongmx/mymicro/app/user-service/user-service /go/src/github.com/mongmx/mymicro/app/user-service/wait-for ./

CMD ["./wait-for", "database:5432", "--", "./user-service"]