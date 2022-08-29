# RECORDS API

- [GitHub](https://github.com/jafarsirojov/mongoDB) - GitHub repository.

## Frameworks
- [UberFX micro framework](https://godoc.org/go.uber.org/fx) - DI framework.
- [gorilla/mux](https://github.com/gorilla/mux) - Gorilla MUX.
- [go-cache](github.com/patrickmn/go-cache) - In-memory key:value cache.
- [UberZAP logging](https://godoc.org/go.uber.org/zap) - Blazing fast, structured, leveled logging in Go.


## API

### Get all records
{api_address}:7777/api/record/v1/all

Method: GET

Responses:

```
{
    "code": 200,
    "message": "Success",
    "payload": [
        {
            "id": "6307bedf586634221590e75c",
            "name": "mongoDB",
            "text": "wiuefbwwoofw vwoiecwcops jcwsnc sw iopwncinwec we c",
            "status": 0,
            "createdAt": "2022-08-27T20:41:49.67Z",
            "updatedAt": "2022-08-27T21:02:34.857Z"
        },
        {
            "id": "630a82b979944d946423078f",
            "name": "name1",
            "text": "gyedcenwc3456uwuwcnw",
            "status": 2,
            "createdAt": "2022-08-27T20:46:49.67Z",
            "updatedAt": "2022-08-27T20:46:49.67Z"
        }
    ]
}
```

```
{
    "code": 400,
    "message": "BadRequest",
    "payload": null
}
```

```
{
    "code": 404,
    "message": "NotFound",
    "payload": null
}
```

```
{
    "code": 500,
    "message": "InternalErr",
    "payload": null
}
```




### Update record
{api_address}:7777/api/record/v1/update/{id}

Method: PUT

Request:

```
{
    "status": 1,
    "name": "newName",
    "text": "new text"
}
```

Responses:

```
{
    "code": 200,
    "message": "Success",
    "payload": null
}
```

```
{
    "code": 400,
    "message": "BadRequest",
    "payload": null
}
```

```
{
    "code": 404,
    "message": "NotFound",
    "payload": null
}
```

```
{
    "code": 500,
    "message": "InternalErr",
    "payload": null
}
```




### Delete record
{api_address}:7777/api/record/v1/delete/{id}

Method: DELETE

Responses:

```
{
    "code": 200,
    "message": "Success",
    "payload": null
}
```

```
{
    "code": 400,
    "message": "BadRequest",
    "payload": null
}
```

```
{
    "code": 500,
    "message": "InternalErr",
    "payload": null
}
```
