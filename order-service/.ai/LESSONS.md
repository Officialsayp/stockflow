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
