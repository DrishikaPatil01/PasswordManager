# Update Credentials LLD/API Contract

## Description
The update credentials api is used to update existing credentials for an authorized user in the password manager.

## LLD
![Delete Credentials Low Level Diagram](../assets/UpdateCredsLld.png)


## Request

### Path
| **Field** | **Value**                             |
|-----------|-----------                            |
| Base Url  | http://localhost:8080/password-manager|
| Path      |    /credentials?id=id                 |
| Headers   | Content-Type: application/json        |

### Request Body
| **Field**    | **Description**                           | **Valid Values**                                                                 |
|--------------|-------------------------------------------|----------------------------------------------------------------------------------|
| username     | username for credential to be changed     |                                                                                  |
| password     | password for credential to be changed     |                                                                                  |
| options      | options for credential to be changed      | List of json objects                                                             |


### Sample Request
```
curl --location --request PUT 'http://localhost:8080/password-manager/credentials?id=id' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer aksjfkfjdlkjfkdsjfkl' \
--data '{
    "username" : "email@gmail.com",
    "password" : "abc123@Abc",
    "options" : "[]"
}'
```

## Response

### Response Body
| **Field**             | **Description**                                |
|-----------------------|------------------------------------------------|
| status                | status of update credentials - SUCCESS/FAILED  |
| credential.id         | id of credentials assigned                     |
| credential.username   | username of credentials                        |
| credential.password   | password of credentials                        |
| credential.options    | options of credentials                         |
| error.Code            |                                                |
| error.Description     |                                                |

### Sample Response
```
{
    "status" : "SUCCESS"/"FAILIURE",
    "credential" : {
        "id" : "<id>",
        "username" : "<username>",
        "password" : "<password>",
        "options" : "[{""},{""},{""}]"
    }
    "error" : {
        "code" : "",
        "description" : ""
    }
}
```


