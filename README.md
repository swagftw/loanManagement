# Loan Management Service API

This API is Loan Management microservice written in Golang using Go-kit.

> - It uses http as transport layer
> - JSON data is communicated between client and api
> - MongoDB is used as datastore
> - Email is sent to user using SMTP server, after loan status changes
> - Pagination is allowed, result limiting is allowed

**Postman Collection for api**

> [Loan API.postman_collection.json](https://github.com/swagftw/loanManagement/blob/main/Loan_API.postman_collection.json)



**Testing API locally**

Prerequisites

> - Go 1.15
> - MongoDB ^latest version

Changes

> - Change database `(dbString)` string in `main.go`
> - Default port is `8080` you can use any valid port of your choice by changing it in `main.go`
> - Sending mail `email_service.go` `google SMTP server`

Change sender id

> - Add App password for that Google account

And you are good to go

## Endpoints

**Create a loan**

> `POST` '/loans'
> 
> `BODY`
> 
>```json
>{
>    "customerName":"User Name",
>    "phoneNo":"1234567890",
>    "email":"user@example.com",
>    "loanAmount":123456,
>    "creditScore":800
>    }
>```

**Get loans**

> `GET` '/loans'
>
> `QUERY PARAMETERS`
> 
> - status [new,cancelled,approve,reject]
> - page [1,...] {for pagination of loans}
> - limit [1,...] {loans per page}
> - loanAmountGreater [0,...] {filter loan with Loan Amount}

**Get loan with id**

> `GET` '/loan/{id}' 

**Update loan status to approve or reject**

> `PATCH` '/loans/{id}'
> 
>`BODY` 
>```json
>{"status":"approve"}
>```

**Cancel loan**

> `DELETE` '/loans/{id}'
> 