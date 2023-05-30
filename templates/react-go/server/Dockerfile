FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/server

FROM scratch
WORKDIR /out
COPY --from=builder /app/server /out/server
COPY --from=builder /app/public/ /out/public/
EXPOSE 8080
ENTRYPOINT ["/out/server"]