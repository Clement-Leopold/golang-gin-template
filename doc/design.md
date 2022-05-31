# Design Details:

## Tech Stack：
    
1. gin: web-framework
2. DB: postgresSQL
3. use domain driven to form the basic layout of projects.

## package structure
- domain:
  - user:
    - http/routers: specify the rest api details.
    - method: methods against with the user domain 
    - repository: DB functions related to user
- common: 
  - code.go: defines ErrorCode, Response Format
  - db.go :init DB Connections
- config.json: defines several parameters for use.
- main.go: the entrypoint
- schema.sql: the db sql file.


## api
### Get Users
#### request:
- URL:
Get http://localhost:10010/v1/users/?limit=5&offset=0
- query:
  - limit: 
    - int16 
    - not required
  - offset: 
    - int16 
    - not required

#### response:
```json
[
    {
        "id": "2a79e84c-8a33-49f6-9de7-dead84aef404",
        "name": "第二",
        "dob": "1990-07-20",
        "address": "shenzhen",
        "description": "第二个用户",
        "CreatedAt": "2022-05-31T16:08:40.072875Z"
    },
    {
        "id": "5b4b803d-12da-4625-8f09-da1a3a99e907",
        "name": "第三",
        "dob": "1990-07-20",
        "address": "shenzhen",
        "description": "第二个用户",
        "CreatedAt": "2022-05-31T16:18:34.176547Z"
    }
]

```
### Get User By Id
#### request:
- URL:
Get http://localhost:10010/v1/users/:id

- Param:
  - id: 
    - string 
    - required 
    - userId fetched in GetUsers

#### response:
```json
[
    {
        "id": "2a79e84c-8a33-49f6-9de7-dead84aef404",
        "name": "第二",
        "dob": "1990-07-20",
        "address": "shenzhen",
        "description": "第二个用户",
        "CreatedAt": "2022-05-31T16:08:40.072875Z"
    }
]

```
### Create
#### request:
- URL:
POST http://localhost:10010/v1/users

```json
    {
    "name": "第三",
    "address": "shenzhen",
    "description": "第二个用户",
    "dob": "1990-07-20"
    }

```
### Delete
#### request:
- URL:
Delete http://localhost:10010/v1/users/:id

- Param:
  - id: 
    - string
    - required 
    - userId fetched in GetUsers

```json
    {
    "name": "第三",
    "address": "shenzhen",
    "description": "第二个用户",
    "dob": "1990-07-20"
    }

```
### Update
#### request:
- URL:
PUT http://localhost:10010/v1/users/:id

- Param:
  - id: 
    - string
    - required 
    - userId fetched in GetUsers

```json
    {
    "name": "第三",
    "address": "shenzhen",
    "description": "修改",
    "dob": "1990-07-20"
    }
```


## how to start

```
go build
./user-management

```


  
