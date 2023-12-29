# redis-go
In this repo I will document all I know about how to use redis proffessionaly in Golang applications

# To start redis using Docker 
- pull the image from docker hub
```shell
docker pull redis 
```
- run the containe and expose a port
```shell
docker run --name redisgo -p 6379:6379 redis
```
- to access the shell inside the running instance 
```shell
docker exec -it redisgo redis-cli
```

# How to `SET strings` into redis
```powershell
SET key value \
[Options to specify when the value will expire] \
[option to specify to set this key if only its not set before or set it only if it does already exists] \
[option to ask redis to return the old value stored in this key]
``` 

Examples:
- we can get the previous stored value before we override it
```powershell
127.0.0.1:6379> SET username fady
OK
127.0.0.1:6379> SET username marwan GET
"fady"
127.0.0.1:6379>
```      

- we can set the key only if this key exists before
- we can do this using the `XX` option 
```powershell
127.0.0.1:6379> GET age
(nil)
127.0.0.1:6379> SET age 25 XX
(nil)
127.0.0.1:6379> SET age 25
OK
127.0.0.1:6379> SET age 24 XX
OK
```

- now we can specify that we need to set the key if its doesn't exists in our database before
- we can do this using the `NX` option

```powershell
127.0.0.1:6379> GET age
"24" # so the age exists 
127.0.0.1:6379> SET age 23 NX
(nil) # we couldn't set it
127.0.0.1:6379>
```