# svad

instant microservice proxy server.


```go

go get github.com/kcwebapply/svad

./svad ../config.toml # db setting file

```

- [proxy](#proxy)
- [add (add memo)](#add)
- [rm (delete memo)](#rm)
- [clear (delete all memo)](#clear)


### register
----
**POST `/register`** 

- header
  - `service_name` : service name binded with hosts.

- requestBody
```java
// post urls to register on service.
{
  "hosts":["https://jsonplaceholder.typicode.com"]
}
```

Then, you can register service with hosts. 

### services
----
**GET `/services`**

response:
```
{
    "api": [
        "https://jsonplaceholder.typicode.com",
        "https://jsonplaceholder-2.typicode.com"
    ],
    "yahoo": [
        "http://www.yahoo.co.jp"
    ]
}
```

you can get service list with its binded hosts.



### proxy

**`/svad/**`**

- header
  - service_name: designate service you want to send request to.
  - request_type: `request_type` only allows two values `all` or `rand`.
    - `all` diffuse http_request to all hosts binded with service.
    - `rand` select only 1 host for forwading request.


```shell
## rand type proxy
curl http://localhost:8888/svad/posts/1 -H "service_name:api" -H "request_type:rand"

--> proxying request to 
GET https://jsonplaceholder.typicode.com/posts/1
or 
GET https://jsonplaceholder.typicode.com/posts/1

## all type proxy
curl http://localhost:8888/svad/posts/1 -H "service_name:api" -H "request_type:all"

-->proxying request to 
GET https://jsonplaceholder.typicode.com/posts/1
GET https://jsonplaceholder-2.typicode.com/posts/1
```


