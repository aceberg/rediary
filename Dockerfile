FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/rediary/ && CGO_ENABLED=0 go build -o /rediary .


FROM scratch

WORKDIR /data/rediary
WORKDIR /app

COPY --from=builder /rediary /app/

ENTRYPOINT ["./rediary"]