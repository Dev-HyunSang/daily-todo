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
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6Ijg3YzMwZjJlLTg0YTItNDUzYS05NDM2LTliY2I0ZmVkZTI2OSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY4MzAxNDI1OSwidXNlcl91dWlkIjoiZTcyYmNiNzgtOWJhYi00MzYwLWE5NmYtOGM0MDMzYzg2YmI3In0.Ae4YvshiuEvJcIWewjb8qUA-Wg_KHnOzJnn1uhd92qw",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM2MTgxNTksInJlZnJlc2hfdXVpZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCIsInVzZXJfdXVpZCI6ImU3MmJjYjc4LTliYWItNDM2MC1hOTZmLThjNDAzM2M4NmJiNyJ9.S788UuZpPxD_KQLgiXRMZMKh_jsNDcTQgAguuA8nqJk",
        "access_uuid": "87c30f2e-84a2-453a-9436-9bcb4fede269",
        "refresh_uuid": "00000000-0000-0000-0000-000000000000",
        "at_expires": 1683014259,
        "rt_expires": 1683618159
    },
    "responsed_at": "2023-05-02T16:42:39.091634+09:00"
}
```