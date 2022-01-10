classroom-getter
=============
TaskCollect microservice for interacting with Google Classroom.

Currently in prototyping stage, in a rewrite of the entire thing.

## HTTP SPEC

**GET** `/v1/tasks`

Returns all active assignments for the user, alongside their materials, and submission information.

Request:

```jsonc
{
    "token": {
        "refresh": "AbCdEfGhIjKlMn...",
        "access": "OpQrStUvWxYz1234...",
        "expires": 1234567890 // unix timestamp
    }
}
```

Response:

```jsonc
[
    {
        "materials": [
            {
                "title": "Some Document",
                "link": "https://docs.google.com/abc123"
            }
        ],
        "id": "12345678910",
        "name": "Example Assignment",
        "courseId": "987654321",
        "courseName": "11 Maths Methods or something",
        "description": "In this assignment we do lots of maths",
        "dueOn": 1624976940, // unix timestamp
        "setOn": 1624928174, // unix timestamp
        "submission": {
            "id": "AbCDeF156172gHIj",
            "state": "TURNED_IN", // refer to google classroom for enum values
            "late": "false" // is it late?
        }
    },
    { /* another assignment */ },
    { /* another assignment */ },
    { /* another assignment */ }
]
```