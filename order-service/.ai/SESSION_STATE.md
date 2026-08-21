# SESSION_STATE.md

## Текущий этап

GET handlers

## Изучено

- запуск HTTP-сервера через `http.ListenAndServe`;
- маршрутизация через `http.ServeMux` и method pattern;
- handler с сигнатурой `func(http.ResponseWriter, *http.Request)`;
- path parameter через `r.PathValue`;
- query parameter через `r.URL.Query().Get`;
- преобразование строковых параметров через `strconv.Atoi` и `strconv.ParseBool`;
- `200 OK`, `204 No Content`, `400 Bad Request`, `405 Method Not Allowed`;
- отправка response и заголовка `Content-Type`.

## В процессе

Закрепление GET endpoints и различия между path и query parameters.

## Текущая задача

Добавить `GET /health`:

- создать `healthHandler`;
- вернуть `204 No Content` без body;
- зарегистрировать маршрут на `ServeMux`;
- проверить `GET` и `POST` через `curl`.

## Следующая тема

`http.Handler` и `http.HandlerFunc`: зачем они нужны и как `ServeMux` вызывает handler.

Во время следующего занятия продолжить именно с `GET /health`.
