### GOLANG-TELEGRAM-MAYNER BOT

## Описание

Простой бот, симулирующий майнинг криптовалюты! Поиграть *@pidwinbot*

Играйте присылайте pull-requests!

## Dependency
```
go get github.com/jinzhu/gorm
go get gopkg.in/telegram-bot-api.v4
```

## Docker
Лучше всего использовать в докер контейнере с компиляцией в один бинарник
```Dockerfile
FROM golang:1.7.5
RUN go get github.com/mattn/go-sqlite3 && go get gopkg.in/telegram-bot-api.v4
WORKDIR /server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch  
WORKDIR /root/
COPY --from=0 /server/app .
COPY --from=0 /server/data.db .
CMD ["./app"]  
```

## Разработчек 
[likipiki](https://github.com/LikiPiki), [Kolya Raketa](https://github.com/kolyaraketa)🚀
