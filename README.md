# Country Information and Population Service

This project is a REST web application developed in Go that provides information about countries and their historical populations. The service integrates with the CountriesNow API and the REST Countries API to fetch and return the required data.

## Installing

Clone the repository and move to the folder.

## Dependencies

- go
- github.com/joho/godotenv

Copy `.env.example` to `.env`. This allows you to update environment variables of the program from the files.

### External REST web services

- _CountriesNow API_
  - Endpoint: http://129.241.150.113:3500/api/v0.1/
  - Documentation: https://documenter.getpostman.com/view/1134062/T1LJjU52
- _REST Countries API_
  - Endpoint: http://129.241.150.113:8080/v3.1/
  - Documentation: http://129.241.150.113:8080/

## Building and Running from Terminal

To build, run the command. You can replace `app` with what you want the executable to be called.

```sh
go build -tags netgo -o app ./cmd
```

Then to run:

```sh
./app
```

# Endpoints

The web service has three resource root paths:

```
/countryinfo/v1/info/
/countryinfo/v1/population/
/countryinfo/v1/status/
```

Assuming the web service should run on localhost, port 8080, the resource root paths would look something like this:

```
http://localhost:8080/countryinfo/v1/info/
http://localhost:8080/countryinfo/v1/population/
http://localhost:8080/countryinfo/v1/status/
```

## Country Info Endpoint: Return general country infos

The initial endpoint focuses returns general information for a given country, [2-letter country codes (ISO 3166-2)](https://en.wikipedia.org/wiki/ISO_3166-2).

### Request

```
Method: GET
Path: info/{:two_letter_country_code}{?limit=10}
```

- `two_letter_country_code` is the corresponding [2-letter country ISO codes](https://en.wikipedia.org/wiki/ISO_3166-2)
- `limit` is the number of cities that are listed in the response. The listing of cities should be in ascending alphabetical order. The parameter is optional.

Example request: `info/no`

### Response

- Content type: `application/json`
- Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):

```json
[
  {
    "name": "Norway",
    "continents": ["Europe"],
    "population": 4700000,
    "languages": {
      "nno": "Norwegian Nynorsk",
      "nob": "Norwegian Bokmål",
      "smi": "Sami"
    },
    "borders": ["FIN", "SWE", "RUS"],
    "flag": "https://flagcdn.com/w320/no.png",
    "capital": "Oslo",
    "cities": ["Abelvaer", "Adalsbruk", "Adland"]
  }
]
```

## Country Population Endpoint: Return population levels for given time frames

The second endpoint should return population levels for individual years for a given country (identified based on country code), as well as the mean value of those. Optionally, the endpoint should allow you to limit the number of returned values by time frames. Otherwise, all values are returned.

### Request

```
Method: GET
Path: population/{:two_letter_country_code}{?limit={:startYear-endYear}}
```

- `{:two_letter_country_code}` refers to the ISO 3166-2 identifier of the country.
- `{?limit={:startYear-endYear}}` is an optional parameter that constrains the population history to values between start year and end year (boundary values are included).

Example requests:

- `population/no`
- `population/no?limit=2010-2015`

### Response

- Content type: `application/json`
- Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):

```json
{
  "mean": 5044396,
  "values": [
    { "year": 2010, "value": 4889252 },
    { "year": 2011, "value": 4953088 },
    { "year": 2012, "value": 5018573 },
    { "year": 2013, "value": 5079623 },
    { "year": 2014, "value": 5137232 },
    { "year": 2015, "value": 5188607 }
  ]
}
```

## Diagnostics Endpoint: Getting a status overview of services

The diagnostics interface indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.

### Request

```
Method: GET
Path: status/
```

### Response

- Content type: `application/json`
- Status code: 200 if everything is OK, appropriate error code otherwise.

Body:

```json
{
   "countriesnowapi": "<http status code for CountriesNow API>",
   "restcountriesapi": "<http status code for RestCountries API>",
   "version": "v1",
   "uptime": <time in seconds since the last re/start of the service>
}
```

# Extra features

- Integrated support for defining variables in a `.env` file.
- Request timeout after 5 seconds if the services don't respond.
- HTML pages for each endpoint to guide the user on how to use the service.
- Redirecting to the startpage from the root path.
