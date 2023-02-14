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

---

## Logging

```bash
2023/02/14 10:39:04 [Info]: Orders Book | Server: v0.0.1, api-v1 
2023/02/14 10:39:04 [Info]: Server started at :8081 
2023/02/14 10:39:15 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-14] room:[4] to:[2023-02-18]] 127.0.0.1:53984 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 10:39:24 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-14] room:[lux] to:[2023-02-18]] 127.0.0.1:53984 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 10:42:20 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-14] room:[lux] to:[2023-02-18]] 127.0.0.1:46572 curl/7.87.0 
2023/02/14 10:45:03 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-14] room:[lux] to:[2023-02-18]] 127.0.0.1:51274 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 10:45:39 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-14] room:[luxx] to:[2023-02-18]] 127.0.0.1:51274 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 10:47:10 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-14] room:[lux] to:[2023-02-183]] 127.0.0.1:51274 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 10:47:22 [Info]: 127.0.0.1:8081 GET map[email:[qwe@qwe.qwe] from:[2023-02-144] room:[lux] to:[2023-02-18]] 127.0.0.1:51274 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
2023/02/14 10:47:53 [Info]: 127.0.0.1:8081 GET map[email:[] from:[2023-02-14] room:[lux] to:[2023-02-18]] 127.0.0.1:51274 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 
```

