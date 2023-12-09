# Login LLD/API Contract

## Description
The login api is used to login the user in the password manager.

## LLD
![Login Low Level Diagram](../assets/LoginLld.png)


## Request

### Path
| **Field** | **Value**                             |
|-----------|-----------                            |
| Base Url  | http://localhost:8080/password-manager|
| Path      |    /user/login                        |
| Headers   | Content-Type: application/json        |

### Request Body
| **Field** | **Description**             | **Valid Values**                                                                 |
|-----------|-----------------------------|----------------------------------------------------------------------------------|
| email     | username for user to login  | valid email                                                                      |
| password  | password for user to login  | Contains lowercase, uppercase, special character, digits and minimum length of 8 |


### Sample Request
```
curl --location 'http://localhost:8080/password-manager/user/login' \
--header 'Content-Type: application/json' \
--data '{
    "email" : "email@gmail.com",
    "password" : "abc123@Abc"
}'
```

## Response

### Response Body
| **Field**          | **Description**                    |
|--------------------|------------------------------------|
| status             | status of login  - SUCCESS/FAILED  |
| userId             |    user's Id assigned by service   |
| email              |    user's email                    |
| error.Code         |                                    |
| error.Description  |                                    |

### Sample Response
```
{
    "status" : "SUCCESS"/"FAILIURE",
    "userId" : "e27d273a-8f9b-11ee-b9d1-0242ac120002",
    "email" : "email@gmail.com",
    "error" : {
        "code" : "",
        "description" : ""
    }
}
```


