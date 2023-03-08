FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/red/ && CGO_ENABLED=0 go build -o /red .


FROM scratch

WORKDIR /data/red
WORKDIR /app

COPY --from=builder /red /app/

ENTRYPOINT ["./red"]