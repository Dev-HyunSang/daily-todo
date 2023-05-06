# daily-todo

## Table of Content
- [daily-todo](#daily-todo)
  - [Table of Content](#table-of-content)
  - [REST API Docs](#rest-api-docs)
    - [POST `/api/user/join`](#post-apiuserjoin)
      - [Request](#request)
      - [Response](#response)
    - [POST `/api/user/login`](#post-apiuserlogin)
      - [Request](#request-1)
      - [Response](#response-1)
    - [POST `/api/todos/create`](#post-apitodoscreate)
      - [Request](#request-2)
      - [Response](#response-2)
    - [POST `/api/todos/list`](#post-apitodoslist)
      - [Request](#request-3)
      - [Response](#response-3)

## REST API Docs

### POST `/api/user/join`
#### Request
```json
{
    "email": "me@hyunsang.dev",
    "password": "password",
    "nickname": "HyunSang Park"
}
```

#### Response
```json
{
    "meta_data": {
        "is_success": true,
        "status_code": 200,
        "message": "어서와요! 성공적으로 회원가입이 완료되었습니다!"
    },
    "data": {
        "id": 1,
        "user_uuid": "e104b644-3254-45b7-a613-f5a5dac9ff0d",
        "email": "me@hyunsang.dev",
        "password": "$2a$10$Mlz6Q7BXsfqJijkYvfxNzu0sA2Hre3pj4lU4l8j1P1sojxo15AI4C",
        "nickname": "HyunSang Park",
        "created_at": "2023-05-02T15:55:50.168361+09:00",
        "updated_at": "2023-05-02T15:55:50.168361+09:00"
    },
    "responsed_at": "2023-05-02T15:55:50.182113+09:00"
}
```

### POST `/api/user/login`
#### Request
```json
{
    "email": "",
    "password": ""
}
```

#### Response
```json
{
    "meta_data": {
        "is_success": true,
        "status_code": 200,
        "message": "어서와요! 성공적으로 로그인이 완료되었습니다!"
    },
    "data": {
        "user_uuid": "e72bcb78-9bab-4360-a96f-8c4033c86bb7",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImIxMzJjYzVlLWNjZjYtNDQ2MC04YTE2LTM1MzU2NGZhMGQ4YSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY4MzM1ODM1NywidXNlcl91dWlkIjoiZTcyYmNiNzgtOWJhYi00MzYwLWE5NmYtOGM0MDMzYzg2YmI3In0.WX3xC5toOa_7JxF9z_0vBM1JUu3MTbaxfQyqS3eOSec",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM5NjIyNTcsInJlZnJlc2hfdXVpZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCIsInVzZXJfdXVpZCI6ImU3MmJjYjc4LTliYWItNDM2MC1hOTZmLThjNDAzM2M4NmJiNyJ9.RrbKQNlFGP9Dabrjuj6BkoNJZ6UbD8tNYnav6GX1Xww",
        "access_uuid": "b132cc5e-ccf6-4460-8a16-353564fa0d8a",
        "refresh_uuid": "00000000-0000-0000-0000-000000000000",
        "at_expires": 1683358357,
        "rt_expires": 1683962257
    },
    "responsed_at": "2023-05-06T16:17:37.489233+09:00"
}
```

### POST `/api/todos/create`
#### Request
- `Bearer Token`를 사용합니다.
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImIxMzJjYzVlLWNjZjYtNDQ2MC04YTE2LTM1MzU2NGZhMGQ4YSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY4MzM1ODM1NywidXNlcl91dWlkIjoiZTcyYmNiNzgtOWJhYi00MzYwLWE5NmYtOGM0MDMzYzg2YmI3In0.WX3xC5toOa_7JxF9z_0vBM1JUu3MTbaxfQyqS3eOSec
```

```json
{
    "context": "Hello, World!"
}
```

#### Response
```json
{
    "meta_data": {
        "is_success": true,
        "status_code": 200,
        "message": "사용자 분의 소중한 할일이 성공적으로 등록되었어요!"
    },
    "data": {
        "id": 1,
        "todo_uuid": "6d3c65a0-fe8a-46b9-930f-4434d4adffc9",
        "user_uuid": "e72bcb78-9bab-4360-a96f-8c4033c86bb7",
        "context": "Hello, World!",
        "created_at": "2023-05-06T16:19:53.009259+09:00",
        "updated_at": "2023-05-06T16:19:53.009259+09:00"
    },
    "responsed_at": "2023-05-06T16:19:53.021696+09:00"
}
```

### POST `/api/todos/list`
#### Request
- `Bearer Token`를 사용합니다.
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImIxMzJjYzVlLWNjZjYtNDQ2MC04YTE2LTM1MzU2NGZhMGQ4YSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY4MzM1ODM1NywidXNlcl91dWlkIjoiZTcyYmNiNzgtOWJhYi00MzYwLWE5NmYtOGM0MDMzYzg2YmI3In0.WX3xC5toOa_7JxF9z_0vBM1JUu3MTbaxfQyqS3eOSec
```
#### Response
```json
{
    "meta_data": {
        "is_success": true,
        "status_code": 200,
        "message": "성공적으로 기록된 할 일들을 불러왔어요!"
    },
    "data": [
        {
            "id": 1,
            "todo_uuid": "6d3c65a0-fe8a-46b9-930f-4434d4adffc9",
            "user_uuid": "e72bcb78-9bab-4360-a96f-8c4033c86bb7",
            "context": "Hello, World!",
            "created_at": "2023-05-06T07:19:53Z",
            "updated_at": "2023-05-06T07:19:53Z"
        }
    ],
    "responsed_at": "2023-05-06T16:26:50.310999+09:00"
}
```