# Simple API in Go | Domain Driven Development


```
git clone https://github.com/biganashvili/landcraft.git .
cd ./landcraft
go mod tidy
go run ./cmd/api/main.go
```
outputs preinited lands:
```
land 0: ee519cf4-1b87-45f5-8fb1-5e194b542a29
land 1: dd75f3af-1779-4749-83f4-b437a1386d95
land 2: 6bcc8a2e-2d12-44dc-9f13-183590108a32
```

# create user 
request:
```
curl --location 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--data '{
    "name":"sergi"
}'
```
response:
```
{
    "data": {
        "id": "2435f0f6-8b2a-4f9c-bbeb-15137eb0c366"
    },
    "error": null
}
```

# List all users:
request:
```
curl --location 'localhost:8080/user/list'
```
response:
```
{
    "data": {
        "users": [
            "2435f0f6-8b2a-4f9c-bbeb-15137eb0c366"
        ]
    },
    "error": null
}
```

# Create order:
request:
```
curl --location 'localhost:8080/order' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":"2435f0f6-8b2a-4f9c-bbeb-15137eb0c366",
    "land_id":"dd75f3af-1779-4749-83f4-b437a1386d95"
}'
```
response:
```
{
    "data": {
        "id": "cdc08933-fbb2-4841-88b7-6cf4a258ea7c"
    },
    "error": null
}
```

# List all orders of user
request:
```
curl --location --request GET 'localhost:8080/order/list' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":"2435f0f6-8b2a-4f9c-bbeb-15137eb0c366"
}'
```
response:
```
{
    "data": {
        "orders": [
            "0e37005f-0c43-4052-86bd-21da8ce8297d",
            "cdc08933-fbb2-4841-88b7-6cf4a258ea7c"
        ]
    },
    "error": null
}
```
