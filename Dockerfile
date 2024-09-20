FROM golang:1.22.4 AS build
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -v -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=build /go/src/app/app .
CMD ["./app"]
