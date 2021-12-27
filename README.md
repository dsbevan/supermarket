# Supermarket API

## Description
The supermarket api allows a user to GET from, POST to, and DELETE from an in-memory database of produce items. Produce items each consist of a name, a produce code, and a unit price.

## Installation
Ensure Docker is installed on your system.
The project can be built using the provided Dockerfile. Builds of the main branch are also available on Docker Hub at `dsbevan/supermarket`. Builds of the `develop` branch are available at `dsbevan/supermarket-develop`. Each image is tagged with the commit hash of the code version used to build the image.

The application is configured using the config file "config.json". An existing image can be configured by mounting a new config.json file with custom values onto the image at /root/config.json.  This file is loaded and validated when the image is run.

By default, the api listens for connections on port 8080 inside the container. The following command will run the container locally (pulling if it has not been pulled), listen for requests at `http://localhost:8080/supermarket/produce`, and remove the container when the process exits:
```
docker run --rm -p 8080:8080 dsbevan/supermarket
```

## API Documentation
#### Produce Items
Produce items have the following format:
```
{
    "name": string,
    "code": string,
    "price": float
}
```

Produce names are alphanumeric, case insensitive, and may contain spaces, but cannot start with a space.

Produce codes are alphanumeric, case insensitive, and have the following format: Four sets of for characters, each separated by a dash ('-'). Ex (Where each '#' represents a character):
`####-####-####-####`

Produce prices are represented using floating-point values and may contain up to two decimal places.

***
### API Methods

#### GET
Get returns a json array of all the produce items in the database.

Example request url:
```
GET http://<domain>/supermarket/produce
```

Example response body:
```
{
    "produce":[
        {
            "name":"Lettuce",
            "code":"A12T-4GH7-QPL9-3N4M",
            "price":3.46
        },
        {
            "name":"Peach",
            "code":"E5T6-9UI3-TH15-QR88",
            "price":2.99
        },
        {
            "name":"Green Pepper",
            "code":"YRT6-72AS-K736-L4AR",
            "price":0.79
        },
        {
            "name":"Gala Apple",
            "code":"TQ4C-VV6T-75ZX-1RMR",
            "price":3.59
        }
    ]
}
```

#### POST
Post allows a user to add items to the database. One or more items may be added at a time. If any item is invalid, the request is rejected and the server responds with a 400 response code. When a POST request is successfully processed, an array of all items successfully added to the database is returned. Any item included in a request but not the response was not added to the database.

Example request url and requst body:
```
POST http://<domain>/supermarket/produce

{
    "produce":[
        {
            "name":"Lettuce",
            "code":"A12T-4GH7-QPL9-3N4M",
            "price":3.46
        },
        {
            "name":"Peach",
            "code":"E5T6-9UI3-TH15-QR88",
            "price":2.99
        },
        {
            "name":"Green Pepper",
            "code":"YRT6-72AS-K736-L4AR",
            "price":0.79
        },
    ]
}
```

Example request response:
```
{
    "produce":[
        {
            "name":"Lettuce",
            "code":"A12T-4GH7-QPL9-3N4M",
            "price":3.46
        },
        {
            "name":"Peach",
            "code":"E5T6-9UI3-TH15-QR88",
            "price":2.99
        },
    ]
}
```
 
#### DELETE
Delete deletes the item with the matching code (case insensitive) from the database, if it exists. The code is passed using the url parameter `Produce Code`.

Example request url:
```
http://<domain>/supermarket/produce?Produce+Code=E5T6-9UI3-TH15-QR88
```

Example response body:
```
{
    "success":true
}
```
