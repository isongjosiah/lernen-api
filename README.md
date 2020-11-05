# Lernen-api

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

To view the instance of the database running on `ElephantSQL` you can do so
at this link [ElephantSQL](https://api.elephantsql.com/console/6fb7001a-11f5-49f0-88e7-61a08c8ac348/browser?# "ElephantSQL home") using these login credentials.
|Email|Password|
|-----|--------|
|isongjosiah@gmail.com|@c@dem!@2020.|

Select the lernen instance an in the page that appears select the browser option in the side navbar 
which brings up an interface for you to query the table. 
Basically just clicking on the `table queries` button and then the `Execute` button should bring the result of a query up.
