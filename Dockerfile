FROM golang:latest AS gobuilder

WORKDIR /go/src/mongo-api

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



FROM alpine:edge as img_app

WORKDIR /mongo/app

COPY --from=gobuilder /go/src/mongo-api/ .

EXPOSE 8000

EXPOSE 5434

CMD ["./main"]


FROM postgres:9.6.18 as img_db




