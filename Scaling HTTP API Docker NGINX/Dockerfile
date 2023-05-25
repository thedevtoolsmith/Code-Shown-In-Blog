FROM golang:1.19-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o /exe

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /exe /app
ENTRYPOINT /app
LABEL Name=simpleapi Version=0.0.1
EXPOSE 11011
