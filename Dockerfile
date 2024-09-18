FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o crm.ecommerce.com ./cmd/server

FROM scratch

COPY ./config /config

EXPOSE 8080

COPY --from=builder /build/crm.ecommerce.com /


ENTRYPOINT [ "/crm.ecommerce.com", "config/local.yaml"]