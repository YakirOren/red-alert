### What is this service?

* This service exposes a web socket that you could easily subscribe to 
and receive notifications.

### Why?
* The service aims to help lower the amount of requests headed to the API and expose an nicer way to intergrate with it.

### What is polling?

* When entering the alerts section in https://www.oref.org.il/ 
* after opening the network tab you can see
* a GET request begin sent **every second** to the API 😧(not very healthy). 

<img src="https://cdn.discordapp.com/attachments/950238983571513367/1005549165490749461/Screen_Shot_2022-08-06_at_21.52.26.png">

* Imagine the amount of traffic of multiple users 😰😰 
* This services does this polling in the background and allows multiple users to listen
with a web socket. 

### How to run it ?
* create a new `.env` file like described in the example file.
* Use the provided Dockerfile.
* Or build from source
`go run cmd/main.go`

