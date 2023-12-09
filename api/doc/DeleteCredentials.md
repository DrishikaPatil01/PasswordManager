# Delete Credentials LLD/API Contract

## Description
The delete credentials api is used to delete existing credentials for an authorized user in the password manager.

## LLD
![Delete Credentials Low Level Diagram](../assets/deleteCredsLld.png)


## Request

### Path
| **Field** | **Value**                             |
|-----------|-----------                            |
| Base Url  | http://localhost:8080/password-manager|
| Path      |    /credentials?id=id                 |
| Headers   | Content-Type: application/json        |
| Headers   |Authorization: Bearer aksjfkfjd        |


### Request Body
N/a

### Sample Request
```
curl --location --request DELETE 'http://localhost:8080/password-manager/credentials?id=id' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer aksjfkfjdlkjfkdsjfkl' \
```

## Response

### Response Body
| **Field**             | **Description**                                |
|-----------------------|------------------------------------------------|
| status                | status of delete credentials - SUCCESS/FAILED  |
| credential.id         | id of credentials assigned                     |
| error.Code            |                                                |
| error.Description     |                                                |

### Sample Response
```
{
    "status" : "SUCCESS"/"FAILIURE",
    "credential" : {
        "id" : "<id>"
    }
    "error" : {
        "code" : "",
        "description" : ""
    }
}
```


