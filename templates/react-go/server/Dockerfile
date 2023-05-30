FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/server

FROM scratch
WORKDIR /out
COPY --from=builder /app/server /out/server
COPY --from=builder /app/public/ /out/public/
# In case the app calls out to external services
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["/out/server"]