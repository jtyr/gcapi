FROM golang AS builder

COPY . /app

WORKDIR /app

RUN go mod vendor
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X github.com/jtyr/gcapi/pkg/version.Version=$(git describe --tags --exact-match HEAD 2>/dev/null || git rev-parse --short HEAD 2>/dev/null || echo -n unknown)" -o /gcapi-cli ./main.go


FROM alpine

COPY --from=builder /gcapi-cli /bin/

ENTRYPOINT ["/bin/gcapi-cli"]
