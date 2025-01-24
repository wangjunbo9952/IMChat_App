FROM golang:1.20-alpine AS builder

LABEL authors="WangJunBo"

WORKDIR /IMChat_App

COPY go.mod go.sum ./
RUN go mod download

COPY docker .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o IMChat_App .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /IMChat_App/IMChat_App .

EXPOSE 9090

CMD ["./IMChat_App"]
