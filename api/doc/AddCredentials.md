# Add Credentials LLD/API Contract

## Description
The add credentials api is used to add credentials for an authorized user in the password manager.

## LLD
![Add Credentials Low Level Diagram](../assets/AddCredsLld.png)


## Request

### Path
| **Field** | **Value**                             |
|-----------|-----------                            |
| Base Url  | http://localhost:8080/password-manager|
| Path      |    /credentials                       |
| Headers   | Content-Type: application/json        |
| Headers   |Authorization: Bearer aksjfkfjd        |


### Request Body
| **Field**    | **Description**             | **Valid Values**                                                                 |
|--------------|-----------------------------|----------------------------------------------------------------------------------|
| username     | username for credential     |                                                                                  |
| password     | password for credential     |                                                                                  |
| options      | options for credential      | List of json objects                                                             |


### Sample Request
```
curl --location 'http://localhost:8080/password-manager/credentials' \
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
| **Field**             | **Description**                             |
|-----------------------|---------------------------------------------|
| status                | status of add credentials - SUCCESS/FAILED  |
| credential.id         | id of credentials assigned                  |
| credential.username   | username of credentials added               |
| credential.password   | password of credentials added               |
| credential.options    | options of credentials added                |
| error.Code            |                                             |
| error.Description     |                                             |

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


