# Weather API demo

A single-protocol weather API that returns a simple weather response from a number of backend weather providers.

#### To get started:

* Add API key values to .env file

* To spin up the server

```$ docker-compose up weather``` 


The service is now running on `https://localhost:8080/v1/weather?city=melbourne`

#### Running tests:

```$ docker-compose up test```

## Features

* A failover weather provider is fired when the primary one fails. Providers used:
    * Weatherstack (https://weatherstack.com)
    * Openweather (https://openweathermap.org)

* Built-in cache will save results for up to 3 seconds in attempt to alleviate load on the server and provide up-to-date weather information

* The service can be run in Docker so that any particulars of the app, such as Golang and the packages used don't need to be installed individually

### Improvements

* Middleware to catch errors can be implemented to abstract the in-code error handling and improve readability

