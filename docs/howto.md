# how to work with nintendo network (on the command line)

## making http requests

### setting up curl

curl is capable of interacting with the nintendo network rest api, but first
you need to make it resemble a nintendo console

first, you need to write a `headers` file containing headers that the 3ds/wiiu
use to provide information to nn. it should look a little like this

```
X-Nintendo-Client-ID: ea25c66c26b403376b4c5ed94ab9cdea
X-Nintend-FPD-Version: 0000
X-Nintendo-Client-Secret: d137be62cb6a2b831cad8c013b92fb55
X-Nintendo-Platform-ID: 1
X-Nintendo-Device-Type;
X-Nintendo-Device-ID: 1
X-Nintendo-Serial-Number: 1
X-Nintendo-System-Version: 1111
X-Nintendo-Region: 2
X-Nintendo-Country: US
X-Nintendo-Environment: L1
X-Nintendo-Device-Cert;
```
(currently filled with mostly garbage data, will work on fixing some of that later)

next, you need to obtain the `cert.pem` and `key.pem` files. these are the client
certificate and the key used to decrypt it, respectively. the information found in
them can be found in this repository, but i'll let you find that yourself.

that's it! now you can interact with the nn rest api using curl

### interacting with the api

assuming the files prepared in the previous step have all been placed in the same
directory, open up a shell in that dir and run the following command (not including
`$`)

```
$ curl -H @headers --cert ./cert.pem --key ./key.pem -k https://account.nintendo.net/v1/api/people/abcdefg
```

you should see output like this

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<errors>
	<error>
		<code>0100</code>
		<message>Account ID already exists</message>
	</error>
</errors>
```

### interacting with the api easily

if you haven't already noticed, there's a perl script called `nn-curl` in this directory that does exactly
what the `curl` command can do, without having to add all of that additional flags each time. give it a try!
if not, then something is wrong, and you should check the command you just typed,
or check for the files `headers`, `cert.pem`, and `key.pem` in your current diretory

if you did, then you are ready to interact with the nn rest api using curl!
