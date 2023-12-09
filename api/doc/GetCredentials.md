# Get Credentials LLD/API Contract

## Description
The get credentials api is used to get credentials for an authorized user in the password manager. Credentials can be fetched based on id or title, if neither are provided all credentials are fetched.

## LLD
![Get Credentials Low Level Diagram](../assets/GetCredsLld.png)


## Request

### Path
| **Field**                                       | **Value**                             |
|-------------------------------------------------|-------------------------------------  |
| Base Url                                        | http://localhost:8080/password-manager|
| Path to get all credentials                     |    /credentials                       |
| Path to get a single credential based on id     |    /credentials?id=id                 |
| Path to get credentials based on title          |    /credentials?title=title           |
| Headers                                         | Content-Type: application/json        |
| Headers                                         |Authorization: Bearer aksjfkfjd        |


### Request Body
N/A


### Sample Request
```
curl --location 'http://localhost:8080/password-manager/credentials' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer aksjfkfjdlkjfkdsjfkl' \
```

## Response

### Response Body
| **Field**             | **Description**                             |
|-----------------------|---------------------------------------------|
| status                | status of get credentials - SUCCESS/FAILED  |
| credential.id         | id of credentials assigned                  |
| credential.username   | username of credentials                     |
| credential.password   | password of credentials                     |
| credential.options    | options of credentials                      |
| error.Code            |                                             |
| error.Description     |                                             |

### Sample Response
```
{
    "status" : "SUCCESS"/"FAILIURE",
    "credentials" : {[
        "credential" : {
            "id" : "<id>",
            "username" : "<username>",
            "password" : "<password>",
            "options" : "[{""},{""},{""}]"
    }, "credential" : {
            "id" : "<id>",
            "username" : "<username>",
            "password" : "<password>",
            "options" : "[{""},{""},{""}]"
    }
    ]},
    "error" : {
        "code" : "",
        "description" : ""
    }
}
```


