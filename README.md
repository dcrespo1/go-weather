# Weather API

This is a simple Go application that serves weather data for a given city. It fetches data from the OpenWeatherMap API, converts the temperature from Celsius to Fahrenheit, and returns both temperatures in the response.

## Structure

The application is structured into several parts:

- `main.go`: This is the entry point of the application. It sets up the HTTP server and routes.

- `apiConfig`: This struct holds the configuration for the OpenWeatherMap API.
  - You will need to create a file called `.api-conf` in the root of the project and store you api key as a JSON obj

- `weatherData`: This struct holds the weather data returned by the OpenWeatherMap API. It includes a method to convert the temperature from Celsius to Fahrenheit.

- `readAPIConfig`: This function reads the API configuration from a file.

- `getCity`: This function is the handler for the `/weather/:city` route. It fetches the weather data for the given city and returns it in the response.

- `query`: This function fetches the weather data from the OpenWeatherMap API.

## Usage

To run the application, use the following command:

```bash
go run main.go
```
Alternatively, ive included `.air.toml` file so that you can use [cosmtrek/air](https://github.com/cosmtrek/air) go utility.

```bash
air
```
## Debrief

This was another attempt at using the [labstack/echo](https://github.com/labstack/echo) module and i think it was pretty successful.
