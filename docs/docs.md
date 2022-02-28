


# classification Payment API.
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

###  entries

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /entries/{id} | [get entry](#get-entry) |  |
| GET | /entries | [list entry](#list-entry) |  |
  


###  transfers

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /transfers | [create transfer](#create-transfer) |  |
  


###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /register | [create user](#create-user) |  |
| GET | /users/{id} | [get user](#get-user) |  |
| GET | /users | [list user](#list-user) |  |
| GET | /login | [login user](#login-user) |  |
  


## Paths

### <span id="create-transfer"></span> create transfer (*createTransfer*)

```
POST /transfers
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [CreateTransferRequest](#create-transfer-request) | `models.CreateTransferRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-transfer-200) | OK | createTransferResponse response structure after success create user |  | [schema](#create-transfer-200-schema) |
| [400](#create-transfer-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#create-transfer-400-schema) |

#### Responses


##### <span id="create-transfer-200"></span> 200 - createTransferResponse response structure after success create user
Status: OK

###### <span id="create-transfer-200-schema"></span> Schema
   
  

[TransferTXResult](#transfer-t-x-result)

##### <span id="create-transfer-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="create-transfer-400-schema"></span> Schema
   
  

any

### <span id="create-user"></span> create user (*createUser*)

```
POST /register
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
   
  

[User](#user)

##### <span id="create-user-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="create-user-400-schema"></span> Schema
   
  

any

### <span id="get-entry"></span> get entry (*getEntry*)

```
GET /entries/{id}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | int32 (formatted integer) | `int32` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-entry-200) | OK |  |  | [schema](#get-entry-200-schema) |
| [400](#get-entry-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#get-entry-400-schema) |

#### Responses


##### <span id="get-entry-200"></span> 200
Status: OK

###### <span id="get-entry-200-schema"></span> Schema
   
  

[Entry](#entry)

##### <span id="get-entry-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="get-entry-400-schema"></span> Schema
   
  

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

### <span id="list-entry"></span> list entry (*listEntry*)

```
GET /entries
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-entry-200) | OK | listEntyResponse return response type from /entries API |  | [schema](#list-entry-200-schema) |
| [400](#list-entry-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#list-entry-400-schema) |

#### Responses


##### <span id="list-entry-200"></span> 200 - listEntyResponse return response type from /entries API
Status: OK

###### <span id="list-entry-200-schema"></span> Schema
   
  

[][Entry](#entry)

##### <span id="list-entry-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="list-entry-400-schema"></span> Schema
   
  

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

### <span id="login-user"></span> login user (*loginUser*)

```
GET /login
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [LoginUserRequest](#login-user-request) | `models.LoginUserRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#login-user-200) | OK |  | ✓ | [schema](#login-user-200-schema) |
| [400](#login-user-400) | Bad Request | badRequestResponse response structure after invalid input from body |  | [schema](#login-user-400-schema) |

#### Responses


##### <span id="login-user-200"></span> 200
Status: OK

###### <span id="login-user-200-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| access_token | string | `string` |  |  |  |

##### <span id="login-user-400"></span> 400 - badRequestResponse response structure after invalid input from body
Status: Bad Request

###### <span id="login-user-400-schema"></span> Schema
   
  

any

## Models

### <span id="entry"></span> Entry


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Amount | int64 (formatted integer)| `int64` |  | |  |  |
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| Currency | string| `string` |  | |  |  |
| ID | int32 (formatted integer)| `int32` |  | |  |  |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| UserID | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="transfer"></span> Transfer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Amount | int64 (formatted integer)| `int64` |  | |  |  |
| CreatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| FromUserID | int32 (formatted integer)| `int32` |  | |  |  |
| ID | int32 (formatted integer)| `int32` |  | |  |  |
| Status | string| `string` |  | |  |  |
| ToUserID | int32 (formatted integer)| `int32` |  | |  |  |
| UpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |



### <span id="transfer-t-x-result"></span> TransferTXResult


> TransferTXResult is the result of the transfer transaction
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| from_user | [User](#user)| `User` |  | |  |  |
| to_user | [User](#user)| `User` |  | |  |  |
| transfer | [Transfer](#transfer)| `Transfer` |  | |  |  |



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



### <span id="create-transfer-request"></span> createTransferRequest


> createTransferRequest user request parameter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Amount | int64 (formatted integer)| `int64` |  | |  |  |
| FromUserID | int32 (formatted integer)| `int32` |  | |  |  |
| ToUserID | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="create-user-request"></span> createUserRequest


> createUserRequest user request parameter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Currency | string| `string` |  | |  |  |
| Password | string| `string` |  | |  |  |
| Username | string| `string` |  | |  |  |



### <span id="login-user-request"></span> loginUserRequest


> loginUserRequest user login request parameter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Password | string| `string` |  | |  |  |
| Username | string| `string` |  | |  |  |


