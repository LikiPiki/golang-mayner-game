# Dockerfile

FROM golang:1.7.5
RUN go get github.com/mattn/go-sqlite3 && go get gopkg.in/telegram-bot-api.v4
WORKDIR /server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /root/
COPY --from=0 /server/ .
CMD ["./app"]
