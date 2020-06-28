#QuickStart
```bash
go run main.go
```
Try following cmd
```bash
curl -L -X GET 'localhost:8080/uppercase' \
-H 'Content-Type: text/plain' \
--data-raw '{}'
```
Then you will get a response which header is including x-* key value pairs