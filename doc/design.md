# Design Details:

## Tech Stack：
    
1. gin: web-framework
2. DB: postgresSQL
3. package structure:use domain driven to form the basic layout of projects.

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

## design details:
1.  use distance in each row of following relation, to decrease the time of calculation of distance.
2.  if update the address, use go routine to update the distance in following relation just among the following not follower. 

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
{
    "Code": "10000",
    "Message": "OK",
    "Data": [
        {
            "id": "2eea6585-7294-48eb-a0d6-d80d1cc565f8",
            "name": "第二",
            "dob": "1990-07-20",
            "address": "shenzhen",
            "x_coordinate": 3,
            "y_coordinate": 4,
            "description": "第二个用户"
        },
        {
            "id": "ed95af03-6e98-471e-8c9c-d8a8e069ff8c",
            "name": "第1",
            "dob": "1990-07-20",
            "address": "shenzhen",
            "x_coordinate": 30,
            "y_coordinate": 40,
            "description": "第二个用户"
        }
    ]
}

```
### Get User By Id
#### request:
- URL:
Get http://localhost:10010/v1/users/:id

- Param:
  - id: 
    - string 
    - required 
    - userId 

#### response:
```json
{
    "Code": "10000",
    "Message": "OK",
    "Data": {
        "id": "ed95af03-6e98-471e-8c9c-d8a8e069ff8c",
        "name": "第1",
        "dob": "1990-07-20",
        "address": "shenzhen",
        "x_coordinate": 1,
        "y_coordinate": 2,
        "description": "第二个用户"
    }
}

```
### Create
#### request:
- URL:
POST http://localhost:10010/v1/users

```json
{
    "Code": "10000",
    "Message": "OK",
    "Data": {
        "id": "2eea6585-7294-48eb-a0d6-d80d1cc565f8",
        "name": "第二",
        "dob": "1990-07-20",
        "address": "shenzhen",
        "x_coordinate": 3,
        "y_coordinate": 4,
        "description": "第二个用户"
    }
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
    - description: userId

```json
{
    "Code": "10000",
    "Message": "OK",
    "Data": "fb599e73-9833-4dc3-8b66-0739caff04b3"
}

```
### Update
Update user info, and use go routine to update the distance in the relation of following.
#### request:
- URL:
PUT http://localhost:10010/v1/users/:id

- Param:
  - id: 
    - string
    - required 
    - description: userId fetched in GetUsers

```json
   {
    "Code": "10000",
    "Message": "OK",
    "Data": {
        "id": "ed95af03-6e98-471e-8c9c-d8a8e069ff8c",
        "name": "第1",
        "dob": "1990-07-20",
        "address": "shenzhen",
        "x_coordinate": 30,
        "y_coordinate": 40,
        "description": "第二个用户"
    }
}
```

### Following(关注)
#### request:
- URL:
POST http://localhost:10010/v1/users/:id/followings/:following_id?following=1

- Param:
  - id: 
    - string
    - required 
    - description: userId
  - following_id
    - string
    - required
    - description: userId
- Query:
  - following:  
    - 1: following
    - 0: cancel following

```json
   {
    "Code": "10000",
    "Message": "OK",
    "Data": null
}
```

### Following list(list some of the following persons.)
#### request:
- URL:
GET http://localhost:10010/v1/users/:id/followings?following=1

- Param:
  - id: 
    - string
    - required 
    - description: userId
- Query:
  - following:  
    - 1: following(list following, 列出关注)
    - 0: follower(list follower. 列出粉丝)
  - limit: 
    - int16 
    - not required
  - offset: 
    - int16 
    - not required
  - 

```json
  {
    "Code": "10000",
    "Message": "OK",
    "Data": [
        {
            "id": "ed95af03-6e98-471e-8c9c-d8a8e069ff8c",
            "name": "第1"
        }
    ]
} 
```
### find nearest friend
list the nearest friend among the following list by name.
assume the name is unique.
#### request:
- URL:
GET http://localhost:10010/v1/users/nearest-following/:name

- Param:
  - name: 
    - string
    - required 
    - description: username


```json
 {
    "Code": "10000",
    "Message": "OK",
    "Data": {
        "id": "2eea6585-7294-48eb-a0d6-d80d1cc565f8",
        "name": "第二"
    }
} 
```
## how to start

```
go build
./user-management

```


  
