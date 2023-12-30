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

# To delete key-value pair from redis (Whatever the type)
```powershell
DEL [key]
```

# How Redis range queries can solve some performence issues in your application 
- Assume we have this table structure in our relational database, and we have millions of orders
- required queries are : <br>
    1. Fetch one row (order) to view one or more fields of this order (select shipping_addr / select *)
    2. Update one or more properties of a specific order 
    3. Fetch all props of multiple orders 
    4. Insert multiple orders (Rows)
    <br>
![Alt text](problem_solved_by_redis_Range_Queries.png)

- Solution : <br>
    -   Since that we have specific values for each field, for example we will have 20 meals so the possible values in the orderd_meal column will be 1 of 20 values, and same for the side_Dish.
    - we will create some mapping between the known range of values to some numeric values or characters or whatever is reasonable for our usecase
    ![](mapping.png)
    - After we got this encoded table (mapping), we can store some light weight table in redis so we can access and manipulate our data more fast.
    ![Alt text](store-mapping-in-redis.png)
    - Now lets solve the queries in redis-style .. <br>
        1. Fetch one row (order) to view one or more fields of this order (select shipping_addr / select *)
        ```powershell
        GETRANGE "order:1" 0 0 # this will returns the characters from 0 to 0 which is the first character, and now in our application we can map the returned response to the actuall values

        GETRANGE "order:2" 0 1 # return all fields
        ```
        2. Update one or more properties of a specific order 
        ```powershell
        SETRANGE "order:1" 0 B # This will start override the value from index zero, and replcae only one character because we specified only one character, by updating from A to B so we updated the meal from Chicken to Steak

        GETRANGE "order:2" 0 BU # update all fields (meal and side_dish)
        ```
        3. Fetch all props of multiple orders 
        ```powershell
        MGET "order:1" "order:2" "order:3" # returl all fields of these three items
        ```
        4. Insert multiple orders (Rows)
        ```powershell
        MSET "order:3" AU "order:4" BX
        ```
