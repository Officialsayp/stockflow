# SESSION_STATE.md

## Текущий этап

JSON

## Изучено

- запуск HTTP-сервера через `http.ListenAndServe`;
- маршрутизация через `http.ServeMux` и method pattern;
- handler с сигнатурой `func(http.ResponseWriter, *http.Request)`;
- path parameter через `r.PathValue`;
- query parameter через `r.URL.Query().Get`;
- преобразование строковых параметров через `strconv.Atoi` и `strconv.ParseBool`;
- `200 OK`, `204 No Content`, `400 Bad Request`, `405 Method Not Allowed`;
- отправка response и заголовка `Content-Type`.
- `GET /health` и статус `204 No Content` без body;
- разница между `mux.HandleFunc` и `mux.Handle` с `http.HandlerFunc`.
- `POST /orders` и статус `201 Created`;
- ручная проверка endpoints и assertions в Bruno.
- отправка текстового request body из Bruno;
- чтение текстового body через `r.Body` и `io.ReadAll`.
- JSON request body, `json.Decoder` и request DTO.

## В процессе

JSON response: response DTO, `Content-Type` и `json.Encoder`.

## Текущая задача

Вернуть JSON response из `POST /orders`:

- создать простой response DTO;
- установить `Content-Type: application/json`;
- отправить JSON response со статусом `201 Created`;
- проверить JSON response в Bruno.

## Следующая тема

transport validation JSON request.

Во время следующего занятия продолжить именно с JSON response для `POST /orders`.
