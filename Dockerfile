FROM golang:1.16.1 as build
WORKDIR /
COPY main.go go.mod /
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .

FROM golang:1.16.1 as bin
COPY --from=build /api-client-gen /go/bin/

ENTRYPOINT [ "/go/bin/api-client-gen" ]