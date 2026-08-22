# SESSION_STATE.md

## Текущий этап

Service Layer

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
- transport validation поля `product`: пустая строка и строка только из пробелов возвращают `400 Bad Request`;
- Bruno assertions для `400` при пустом `product`, пробелах и некорректном JSON;
- различие transport validation и business validation;
- начальная связка `main -> handler -> OrderService` через передачу зависимости в `createOrderHandler`.

## В процессе

Service Layer: первая бизнес-проверка товара `"unavailable"`.

## Текущая задача

Перед продолжением обучения завершить техническую правку после merge PR №10:

- создать отдельный PR из коммита `6fdab38` (`minimal changes`), в котором только форматирование `service/order_service.go`;
- дождаться зелёной проверки lint в новом PR; старый failed run для коммита `70526b3` не перезапускать;
- затем заменить передачу `err.Error()` клиенту на безопасный текст `product cannot be ordered`;
- исправить Bruno assertion в `Create unavailable order`: `expression` должен быть `res.status`, значение — `400`.

## Следующая тема

Продолжить Service Layer: почему handler не должен знать бизнес-правила и как service возвращает ошибки.

Во время следующего занятия начать с нового PR для форматирования, затем закончить первый сценарий business validation.
