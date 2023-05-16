# Daily DSA
daily gives one question of dsa to solve 

## Basic usage 
run cmd/dailydsa.go

+ auth required URI without auth
    + user profile uri 
        ```bash
        curl http://localhost:8080/user/profile
        ```
    + user history
        ```bash
        curl http://localhost:8080/user/history
        ```
+ auth required URI with auth
    + user profile uri 
        ```bash
        curl --cookie "auth=anything4now" http://localhost:8080/user/profile
        ```
    + user history
        ```bash
        curl --cookie "auth=anything4now" http://localhost:8080/user/history
        ```

+ daily questions (no auth required)
    ```bash
    curl http://localhost:8080/daily
    ```