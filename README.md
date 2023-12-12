# About
Application to see how some stuff works between Go and Kafka

# Built With
- Go 1.21
- Docker
- Kafka

# How To Run
You must have `Make` to run the `Makefile` commands and `Docker`.

## 1- Install dependencies
```
make install
```

## 2 - Run Kafka locally
```
make deps/up
```
## 3 - Run API
```
make api/dev
```
## 4 - Run Consumer
```
make consumer/dev
```

# How it works
Make a POST request to send messages to Kafka.
```
curl --request POST \
  --url http://localhost:8080/users
```

The consumer will receive the messages sent.