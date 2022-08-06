FROM golang:1.17 AS base

WORKDIR Solanum
ENV CGO_ENABLED=0

COPY go.* ./

COPY cmd ./cmd
COPY broadcast ./broadcast
COPY config ./config
COPY redAlert ./redAlert

RUN go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o /go/bin/Solanum ./cmd/server

########
FROM alpine:3.14.4 AS final

WORKDIR Solanum

COPY --from=base /go/bin/Solanum .

ENTRYPOINT ["./Solanum"]
