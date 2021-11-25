FROM golang:1.17-stretch as builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -o sensitive_check main.go

FROM scratch

ARG run_env=prod

ENV APP_RUN_ENV=$run_env

WORKDIR /app

COPY --from=builder /build/sensitive_check /app/

EXPOSE 5001 5002

CMD ["/app/sensitive_check"]

