# SESSION_STATE.md

## Текущий этап

Validation

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
- JSON response через response DTO и `json.Encoder`.

## В процессе

transport validation: корректный JSON и обязательное поле `product`.

## Текущая задача

Проверить обязательное поле `product` в `POST /orders`:

- принять JSON с пустым `product`;
- вернуть `400 Bad Request` до успешного response;
- создать Bruno request и assertion для этого сценария.

## Следующая тема

различие transport validation и business validation.

Во время следующего занятия продолжить именно с transport validation `product`.
