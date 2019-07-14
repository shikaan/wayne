Wayne
===

Minimalistic URL Short(en)er

> [Here](https://www.youtube.com/watch?v=0PQKnjwlN6g) if you don't get the pun.

## :warning: Disclaimer

This is my first experiment in Go. The software itself is still garbage: buggy,
poorly designed and what not. It's just my first attempt to get my hands dirty.

I gave it a fancy name so I am more motivated to keep it going.

## How does it work?

This application exposes three endpoints:

```
POST /api/urls

Request:
{
    "readableHash": string,
    "targetURL"   : string
}

Response:
    * 201 - Created
        {
            "id"          : string,
            "readableHash": string,
            "targetURL"   : string
        }
```

```
GET /api/urls

Response:
    * 200 - OK
        [
            {
                "id"          : string
                "readableHash": string,
                "targetURL"   : string
            }
        ]
```

```
GET /${readableHash}


Response:
    * 301 - Moved Permanently
```

The first two endpoints are normal CRUD operation: create and list URLs.
URLs are stored in a Firebase Database and the application expects a `firebase_auth.json` in the root directory to be able to work.

The database has to have a collection called `urls` whose elements are exactly as you see in the payloads above.

Once you have created a URL with a readable hash, you can use such hash and hit `localhost:8080/${readableHash}` to be redirected to the `targetURL`.


## Getting started

Copy the Firebase service account JSON file in the root directory with the name "firebase_auth.json" and then

```
make start
```