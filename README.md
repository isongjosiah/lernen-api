# Lernen-api

[![DeepSource](https://deepsource.io/gh/isongjosiah/lernen-api.svg/?label=active+issues&show_trend=true&token=iqPqGUucIDx_hvL69dd7Ke0X)](https://deepsource.io/gh/isongjosiah/lernen-api/?ref=repository-badge)     [![DeepSource](https://deepsource.io/gh/isongjosiah/lernen-api.svg/?label=resolved+issues&show_trend=true&token=iqPqGUucIDx_hvL69dd7Ke0X)](https://deepsource.io/gh/isongjosiah/lernen-api/?ref=repository-badge)


## Development Set-up
1. Have `go` installed on your machine
2. clone the repo.

You can use `postman` to query the api at port `8080` i.e `localhost:8080`. Should you 
need to adjust the port you can do so in the .env file at the root of the api project.

To test the sign-up implementation use `localhost:8080/auth/register`, you should also provide the input in json format like so defining your own values. 
```JSON
{
  "firstname":"any name",
  "lastname" : "any name",
  "email": "any email",
  "username": "any username",
  "password": "any password"
  
}
```

The document would be updated as other routes are implemented.
