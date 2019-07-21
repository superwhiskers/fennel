# api documentation for the **official** nintendo network account servers

## api routes

all of these routes are located on the host of the target account server unless otherwise specified

### **get** /people/:nnid

checks if a user id is taken or not

#### request parameters

    - **nnid**: the nnid of the user you want to check the existance of

#### responses

 status code | response body | error code | meaning
-------------|---------------|------------|------------------------------------
 200         | empty         | none       | user doesn't exist
 400         | error xml     | 0100       | user exists
 400         | error xml     | 1104       | an invalid account id was provided
 401         | error xml     | 0004       | invalid credentials were provided

#### examples

##### nonexistant user

```
GET https://account.nintendo.net/v1/api/people/<any nonexistent user>

response status code: 200

response:
<empty>
```

##### existing user

```
GET https://account.nintendo.net/v1/api/people/abcdefg

response status code: 400

response:
<errorxml with code 0100>
```

#### invalid account id

```
GET https://account.nintendo.net/v1/api/people/a

response status code: 400

response:
<errorxml with error code 1104>
```

### **get** /content/agreements/Nintendo-Network-EULA/:country-code/:version

returns the nn eula of the specified version for the specified country

#### request parameters

    - **country-code**: pretty much any valid ISO 3166-1 alpha-2 country code
    - **version**: the version of the eula to get. providing `@latest` gets you the latest one

#### responses

 status code | response body | error code | meaning
-------------|---------------|------------|----------------------------------------------------------------------------------
 200         | eula xml      | none       | your request was valid
 400         | error xml     | 0002       | one or more of your request parameters are invalid
 400         | error xml     | 1101       | an invalid version/country code combination was provided
 401         | error xml     | 0004       | invalid credentials were provided
 405         | empty         | none       | invalid method (appears when get-ing without providing country code and version)
 404         | empty         | none       | agreement not found. appears when providing one parameter but not the other

#### examples

##### TODO: add request examples

### **get** /content/time\_zones/:region/:lang

returns timezones for the specified region in xml format, with names in the specified
language

#### request parameters

	- **region**: pretty much any valid ISO 3166-1 alpha-2 country code
	- **lang**: some valid ISO 639-1 language codes. valid ones vary by selected region. if an invalid one is provided, it provides it in english (`en`)

#### responses

 status code | response body | error code | meaning
-------------|---------------|------------|-----------------------------------
 404         | empty         | none       | invalid parameters were provided
 200         | timezone xml  | none       | your request was valid
 401         | error xml     | 0004       | invalid credentials were provided

#### examples

##### TODO: add request examples

### **get** /admin/time

route that seemingly does nothing aside from allow you to get the `X-Nintendo-Date`
without having to make a request to another route

#### responses

 status code | response body | error code | meaning
-------------|---------------|------------|-----------------------------------
 200         | empty         | none       | your request was valid
 401         | error xml     | 0004       | invalid credentials were provided

#### examples

##### getting `X-Nintendo-Date`

```
GET https://account.nintendo.net/v1/api/admin/time

response status code: 200

response: empty
```

### **get** /admin/mapped\_ids?input\_type=<input\_type>&output\_type=<output\_type>&input=<input>

map given input to an output

#### valid types

	- `user_id`: an nnid
	- `pid`: a pid

#### request parameters

	- **input\_type**: the type of the input. must be one of the valid types
	- **output\_type**: the type of the output expected. must be one of the valid types
	- **input**: the input

#### responses

 status code | response body | error code | meaning
-------------|---------------|------------|-----------------------------------
 400         | error xml     | 1600       | invalid input was given
 200         | mapped id xml | none       | your request was valid
 401         | error xml     | 0004       | invalid credentials were provided

#### examples

##### TODO: add request examples
