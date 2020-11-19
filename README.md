# Mujina 

Mujina takes a configuration (JSON file) and starts a server that returns the configured response (static file).

It should be used for development environment only when you need to "stub" a response from any backend, e.g you are creating a complex page in react that calls multiple backend endpoints instead running the entire backend you can use `Mujina` with all the endpoints and switch to the real backend once your development is in a later state.

## Features

Wildcard: Route configuration supports wildcards using `:`, value will be ignored but this allows matching multiple request to the same endpoint e.g with a route config `/foo/:bar/qux` will match any of this requests `/foo/1/qux` , `/foo/ok/qux`.

NoContent: If configuration does not have a body path the response will have a `Content-Length: 0`, if status code is `204` the response will be a no content even if body path is provided.

Note: It supports only one response per endpoint if you need different responses you would need stop the server and run a different configuration file or run multiple servers (in different port).

## Config

Configuration file should have a JSON format, app by default reads a file called `mujina.json` in the same path as the app but you can run any config file using the param `-configPath=<PATH>`

* method: HTTP method (get, post, put, patch and delete) of the enpoint. default: `get`
* status_code: The http status code that response will have.
* route: Configured route for the endpoint.
* body_path: Path to file that contains the wanted response body for the endpoint, if key is not present the response will be a no content type.

```json
{
  "endpoints": [
    {
      "method":     "get",
      "status_code": 200,
      "route":      "/u/foo",
      "body_path":   "<PATH_TO_PROJECT>/mujina/internal/samples/ok.json"
    },
    {
      "method":     "get",
      "status_code": 200,
      "route":      "/u/foo/:t/wow",
      "body_path":   "<PATH_TO_PROJECT>/mujina/internal/samples/error.json"
    },    
    {
      "method":     "post",
      "status_code": 200,
      "route":      "/u/bar",
      "body_path":   "<PATH_TO_PROJECT>/mujina/internal/samples/created.json"
    }
  ]
}
```

## Usage

1. Download the binary (or build it from the project) and make it available in your path.
2. Create the configuration file with the required endpoints.
3. Run the server:

    * configPath: If config file is in the same folder as the app and name is `mujina.json` app will read it automatically otherwise you need to pass the path with the flag `--configPath` e.g `--configPath=/foo/bar/happy_path.json`.
    * port: By default Mujina will run in port `8080` but you can specify it with the flag `--port`
