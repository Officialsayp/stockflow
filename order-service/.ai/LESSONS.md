# LESSONS

## 2026-08-21

### Тема

GET handlers на стандартном `net/http`.

### Освоено

- запуск сервера, `ServeMux` и регистрация маршрута `GET /orders/{id}`;
- извлечение `id` через `r.PathValue`, преобразование в число и transport validation;
- response с корректными HTTP status codes;
- различие между path parameter (`/orders/42`) и query parameter (`?details=true`);
- обработка необязательного параметра `details` как значения `true` или `false`.

### Важные выводы

- `fmt.Printf` пишет в терминал сервера, а не клиенту; для response нужен `http.ResponseWriter`.
- В handler нельзя вызывать `log.Fatal` для ошибки одного request: это останавливает весь server.
- Все значения из URL сначала являются строками; их нужно явно преобразовывать и проверять.
- Отсутствующий необязательный query parameter — нормальная ситуация, не обязательно ошибка.
- В одном Go package не может быть две функции `main`; независимые приложения должны находиться в разных директориях.

### Следующее занятие

Завершить `GET /health` с `204 No Content`, затем разобрать `http.Handler` и `http.HandlerFunc`.

## 2026-08-22

### Завершено

- `GET /health` с `204 No Content` и без body;
- регистрация handler через `mux.Handle` и `http.HandlerFunc`.

### Важный вывод

`mux.HandleFunc` принимает обычную функцию, а `mux.Handle` ожидает `http.Handler`. `http.HandlerFunc` преобразует функцию в handler.

### Следующее занятие

Базовый `POST /orders`: method pattern и `201 Created` до разбора request body и JSON.

## 2026-08-22 — POST и Bruno

### Завершено

- базовый `POST /orders` с `201 Created`;
- проверка правильного method через `405 Method Not Allowed`;
- Bruno collection в проекте с двумя request и assertions на `201` и `405`.

### Важный вывод

Один Bruno request должен проверять один ожидаемый сценарий: нельзя одновременно ожидать `201` и `405` от одного response.

### Следующее занятие

Разобрать `r.Body` на простом текстовом POST request, затем перейти к JSON и DTO.

## 2026-08-22 — текстовый request body

### Завершено

- отправка text body из Bruno;
- чтение `r.Body` через `io.ReadAll`;
- преобразование байтов request body в строку и отправка response `received: keyboard`.

### Важный вывод

`r.Body` содержит входящие данные request. `io.ReadAll` удобно использовать в учебном примере, но в production body нужно ограничивать по размеру, чтобы клиент не мог отправить чрезмерно большой payload.

### Следующее занятие

JSON request body и DTO для `POST /orders`.

## 2026-08-22 — JSON request DTO

### Завершено

- DTO `createOrderRequest` для поля `product`;
- декодирование JSON из `r.Body` через `json.Decoder`;
- `400 Bad Request` при ошибке разбора JSON;
- проверка JSON request через Bruno.

### Важный вывод

JSON сначала должен быть декодирован в Go-struct. После этого handler работает с понятным полем `req.Product`, а не с исходным текстом JSON.

### Следующее занятие

JSON response: response DTO, `Content-Type: application/json` и `json.Encoder`.
