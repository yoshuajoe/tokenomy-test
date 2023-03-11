# Go HTTP Server
This is a simple HTTP server written in Go that returns data based on the provided ID(s).

# Prerequisites
1. Go 1.16 or later installed
2. Docker and Docker Compose installed (if running using Docker)

# Running using Go
1. Clone this repository to your local machine
2. Navigate to the cloned repository directory
3. Run go run . command in your terminal
4. The server will start running on port 8080

# Running using Docker Compose
1. Clone this repository to your local machine
2. Navigate to the cloned repository directory
3. Run docker-compose up command in your terminal
4. The server will start running on port 8080

# API Endpoint
## Request
`GET /?id=<ID1,ID2,...>`

`id`: a comma-separated list of IDs to retrieve data for. If id is not provided, all available data will be returned.
## Response
The response will be in JSON format and will have the following structure:

```json
{
    "code": HTTP_STATUS_CODE,
    "data": [
        {
            "id": ID,
            "name": NAME
        }
    ],
    "message": MESSAGE
}
```

- `HTTP_STATUS_CODE`: the HTTP status code of the response.
- `ID`: the ID of the data.
- `NAME`: the name of the data.
- `MESSAGE`: a message accompanying the response.

If no data is found for the provided ID(s), a `404 Not Found` HTTP status code will be returned.

If the provided ID(s) is/are invalid, a `400 Bad Request` HTTP status code will be returned.

License
This project is licensed under the MIT License.