[Orders Book](https://github.com/XO490/ordersBook)
---
### Prototype of a hotel room booking service

---

#### Build:
```bash
git clone https://github.com/XO490/ordersBook
cd ordersBook
make build
cd bin/
./ordersBook-l64
```

```bash
./ordersBook-l64 -h           
Usage of ./bin/ordersBook-l64:
  -conf string
        path to config file (default "config.json")
```

default config: config.json
```json
{
  "bind_addr": ":8080",
  "room_types": ["econom", "standart", "lux"]
}
```

---

### Handlers

make order: `/order`
```bash
curl "http://127.0.0.1:8081/order?room=lux&email=qwe@qwe.qwe&from=2023-02-14&to=2023-02-18"
```
success:
```json
{
  "ok": true,
  "result": {
    "http_status": 201,
    "description": "The order has been placed"
  }
}
```
fail:
```json
{
  "ok": false,
  "result": {
    "http_status": 409,
    "description": "Date booking conflict"
  }
}
```
```json
{
  "ok": false,
  "result": {
    "http_status": 400,
    "description": "No available rooms"
  }
}
```
```json
{
  "ok": false,
  "result": {
    "http_status": 400,
    "description": "Invalid date format: from"
  }
}
```
```json
{
  "ok": false,
  "result": {
    "http_status": 400,
    "description": "Required parameters are missing: email, room, from, to"
  }
}
```

get orders: `/orders`
```bash
curl "http://127.0.0.1:8081/orders?email=qwe@qwe.qwe"
```
succes:
```json
{
  "ok": true,
  "result": {
    "http_status": 200,
    "description": "Request successful",
    "orders": [
      {
        "room": "lux",
        "user_email": "qwe@qwe.qwe",
        "from": "2023-02-14",
        "to": "2023-02-18"
      },
      {
        "room": "standart",
        "user_email": "qwe@qwe.qwe",
        "from": "2023-02-19",
        "to": "2023-02-20"
      }
    ]
  }
}
```
fail:
```json
{
  "ok": false,
  "result": {
    "http_status": 400,
    "description": "Required parameters are missing: email"
  }
}
```

---

## Logging

```bash
2023/02/14 11:27:06 [Info]: Orders Book | Server: v0.0.1, api-v1 
2023/02/14 11:27:06 [Info]: Server started at :8081 
2023/02/14 11:27:08 [Info]: 127.0.0.1:8081 GET /order?room=lux&email=qwe@qwe.qwe&from=2023-02-14&to=2023-02-18 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 11:27:26 [Info]: 127.0.0.1:8081 GET /orders?email=qwe@qwe.qwe 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 11:28:01 [Info]: 127.0.0.1:8081 GET /order?room=standart&email=qwe@qwe.qwe&from=2023-02-14&to=2023-02-18 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 11:28:14 [Info]: 127.0.0.1:8081 GET /order?room=standart&email=qwe@qwe.qwe&from=2023-02-15&to=2023-02-18 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 11:28:32 [Info]: 127.0.0.1:8081 GET /order?room=standart&email=qwe@qwe.qwe&from=2023-02-18&to=2023-02-20 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 11:28:58 [Info]: 127.0.0.1:8081 GET /order?room=standart&email=qwe@qwe.qwe&from=2023-02-19&to=2023-02-20 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 11:29:16 [Info]: 127.0.0.1:8081 GET /orders?email=qwe@qwe.qwe 127.0.0.1:37650 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
```

