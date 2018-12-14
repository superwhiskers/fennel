# api documentation for the **official** nintendo network account servers

## api endpoints

all of these endpoints are located on the host of the target account server unless otherwise specified

### **get** /people/:nnid

this endpoint lets you check if a user exists or not

#### request parameters

    - nnid: the nnid of the user you want to check

#### responses

 status code | response body | error code | meanining
-------------|---------------|------------|------------------------------------
 200         | empty         | none       | user doesn't exist
 400         | error xml     | 0100       | user exists
 400         | error xml     | 1104       | an invalid account id was provided
 400         | error xml     | 0004       | invalid credentials were provided

#### examples

##### nonexistent user
```
GET https://account.nintendo.net/v1/api/people/<any nonexistent user>

response status code: 200

response:
<empty>
```

##### existing user
```
GET https://account.nintendo.net/v1/api/people/whiskers

response status code: 400

response:
<errorxml with code 0100>
```

#### invalid account id
```
GET https://account.nintendo.net/v1/api/people/t

response status code: 400

response:
<errorxml with error code 1104>
```

### **get** /content/agreements/Nintendo-Network-EULA/:country-code/:version

#### request parameters

    - country-code: pretty much any valid ISO 3166-1 alpha-2 country code
    - version: the version of the eula to get. providing `@latest` gets you the latest one

#### responses

 status code | response body | error code | meanining
-------------|---------------|------------|----------------------------------------------------
 200         | eula xml      | none       | your request was valid
 400         | error xml     | 0002       | one or more of your request parameters are invalid
 400         | error xml     | 1101       | an invalid version of the eula was requested
 400         | error xml     | 0004       | invalid credentials were provided

#### examples

##### TODO: add request examples