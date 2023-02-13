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