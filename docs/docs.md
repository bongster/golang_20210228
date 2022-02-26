


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

###  me

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /me | [me](#me) |  |
| GET | /me/balance | [me blanace](#me-blanace) |  |
| GET | /me/transactions | [me transactions](#me-transactions) |  |
  


###  transactions

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /transactions | [list transactions](#list-transactions) |  |
  


###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /users | [list users](#list-users) |  |
  


## Paths

### <span id="list-transactions"></span> list transactions (*listTransactions*)

```
GET /transactions
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-transactions-200) | OK |  |  | [schema](#list-transactions-200-schema) |
| [default](#list-transactions-default) | |  |  | [schema](#list-transactions-default-schema) |

#### Responses


##### <span id="list-transactions-200"></span> 200
Status: OK

###### <span id="list-transactions-200-schema"></span> Schema
   
  

[ListTransactionsOKBody](#list-transactions-o-k-body)

##### <span id="list-transactions-default"></span> Default Response


###### <span id="list-transactions-default-schema"></span> Schema

  

[ListTransactionsDefaultBody](#list-transactions-default-body)

###### Inlined models

**<span id="list-transactions-default-body"></span> ListTransactionsDefaultBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | | An optional field name to which this validation applies |  |
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



**<span id="list-transactions-o-k-body"></span> ListTransactionsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



### <span id="list-users"></span> list users (*listUsers*)

```
GET /users
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-users-200) | OK |  |  | [schema](#list-users-200-schema) |
| [default](#list-users-default) | |  |  | [schema](#list-users-default-schema) |

#### Responses


##### <span id="list-users-200"></span> 200
Status: OK

###### <span id="list-users-200-schema"></span> Schema
   
  

[ListUsersOKBody](#list-users-o-k-body)

##### <span id="list-users-default"></span> Default Response


###### <span id="list-users-default-schema"></span> Schema

  

[ListUsersDefaultBody](#list-users-default-body)

###### Inlined models

**<span id="list-users-default-body"></span> ListUsersDefaultBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | | An optional field name to which this validation applies |  |
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



**<span id="list-users-o-k-body"></span> ListUsersOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



### <span id="me"></span> me (*me*)

```
GET /me
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#me-200) | OK |  |  | [schema](#me-200-schema) |
| [default](#me-default) | |  |  | [schema](#me-default-schema) |

#### Responses


##### <span id="me-200"></span> 200
Status: OK

###### <span id="me-200-schema"></span> Schema
   
  

[MeOKBody](#me-o-k-body)

##### <span id="me-default"></span> Default Response


###### <span id="me-default-schema"></span> Schema

  

[MeDefaultBody](#me-default-body)

###### Inlined models

**<span id="me-default-body"></span> MeDefaultBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | | An optional field name to which this validation applies |  |
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



**<span id="me-o-k-body"></span> MeOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



### <span id="me-blanace"></span> me blanace (*meBlanace*)

```
GET /me/balance
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#me-blanace-200) | OK |  |  | [schema](#me-blanace-200-schema) |
| [default](#me-blanace-default) | |  |  | [schema](#me-blanace-default-schema) |

#### Responses


##### <span id="me-blanace-200"></span> 200
Status: OK

###### <span id="me-blanace-200-schema"></span> Schema
   
  

[MeBlanaceOKBody](#me-blanace-o-k-body)

##### <span id="me-blanace-default"></span> Default Response


###### <span id="me-blanace-default-schema"></span> Schema

  

[MeBlanaceDefaultBody](#me-blanace-default-body)

###### Inlined models

**<span id="me-blanace-default-body"></span> MeBlanaceDefaultBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | | An optional field name to which this validation applies |  |
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



**<span id="me-blanace-o-k-body"></span> MeBlanaceOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



### <span id="me-transactions"></span> me transactions (*meTransactions*)

```
GET /me/transactions
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#me-transactions-200) | OK |  |  | [schema](#me-transactions-200-schema) |
| [default](#me-transactions-default) | |  |  | [schema](#me-transactions-default-schema) |

#### Responses


##### <span id="me-transactions-200"></span> 200
Status: OK

###### <span id="me-transactions-200-schema"></span> Schema
   
  

[MeTransactionsOKBody](#me-transactions-o-k-body)

##### <span id="me-transactions-default"></span> Default Response


###### <span id="me-transactions-default-schema"></span> Schema

  

[MeTransactionsDefaultBody](#me-transactions-default-body)

###### Inlined models

**<span id="me-transactions-default-body"></span> MeTransactionsDefaultBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | | An optional field name to which this validation applies |  |
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



**<span id="me-transactions-o-k-body"></span> MeTransactionsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` | ✓ | | The validation message | `Expected type int` |



## Models
