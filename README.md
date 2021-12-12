# url-shortener

## Как запустить

1. Вводим команду для создания docker образов:
```bash
docker-compose build
```
2. Как только образы будут собраны, запускаем контейнеры командой::
```bash
docker-compose up -d
```
3. Для приостановки docker контейнеров используем команду:
```bash
docker-compose down
```

### По умолчанию используется postgres. Чтоб сменить хранилище, надо в файлe docker-compose.yml заменить команду "./app -store postgres" на "./app -store in-memory"
##
##
##
## Пример

1. Создаем короткую ссылку.
```
$ curl --request POST 'localhost:8080/' --header 'Content-Type: application/json' --data-raw '{"url": "https://yandex.ru/search/?lr=28&clid=2270456&win=515&text=golang&src=suggest_B"}'

{
    "shortened_url":"http://localhost:8080/PZD4n80Aep"
}
```

2. Открываем любой браузер и вводим в адресную строку браузера ранее созданную короткую ссылку) 


## Test
```
go test ./... -cover
```