


# Payment API.
Documentation for Payment API
  

## Informations

### Version

1.0.0

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /users | [create user](#create-user) |  |
| GET | /users/{id} | [get user](#get-user) |  |
| GET | /users | [list user](#list-user) |  |
  


## Paths

### <span id="create-user"></span> create user (*createUser*)

```
POST /users
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [CreateUserRequest](#create-user-request) | `models.CreateUserRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-user-200) | OK | createUserResponse response structure after success create user |  | [schema](#create-user-200-schema) |
| [400](#create-user-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#create-user-400-schema) |

#### Responses


##### <span id="create-user-200"></span> 200 - createUserResponse response structure after success create user
Status: OK

###### <span id="create-user-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="create-user-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="create-user-400-schema"></span> Schema
   
  

any

### <span id="get-user"></span> get user (*getUser*)

```
GET /users/{id}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | int32 (formatted integer) | `int32` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-user-200) | OK |  |  | [schema](#get-user-200-schema) |
| [400](#get-user-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#get-user-400-schema) |

#### Responses


##### <span id="get-user-200"></span> 200
Status: OK

###### <span id="get-user-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="get-user-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="get-user-400-schema"></span> Schema
   
  

any

### <span id="list-user"></span> list user (*listUser*)

```
GET /users
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-200) | OK |  |  | [schema](#list-user-200-schema) |
| [400](#list-user-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#list-user-400-schema) |

#### Responses


##### <span id="list-user-200"></span> 200
Status: OK

###### <span id="list-user-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="list-user-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="list-user-400-schema"></span> Schema
   
  

any

## Models

### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Balance | int64 (formatted integer)| `int64` |  | |  |  |
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| Currency | string| `string` |  | |  |  |
| ID | int32 (formatted integer)| `int32` |  | |  |  |
| Password | string| `string` |  | |  |  |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| Username | string| `string` |  | |  |  |



### <span id="create-user-request"></span> createUserRequest


> createUserRequest user request parameter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Currency | string| `string` |  | |  |  |
| Password | string| `string` |  | |  |  |
| Username | string| `string` |  | |  |  |


