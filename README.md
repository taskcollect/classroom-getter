classroom-svc
=============
TaskCollect microservice for interacting with Google Classroom.

Developer notes
---------------
**Only just setup the development environment; don't expect any functionality.**

All platforms which use Google authentication need to be treated separately. As we need to use the Google-preferred way of authentication (with the Google auth dialog and whatnot), the TaskCollect user manager needs to be built in a way that allows this. We shall find out how to do this during `classroom-svc` development.

Another important aspect that Google allows for (and promotes) is working with partial resources. The Google API supports requesting certain data fields to minimise the amount of overhead, thus *significantly* increasing performance. More on this here: https://developers.google.com/classroom/guides/performance

HTTP Spec
---------
* GET /v1/tasks
  * Request body (JSON)

    ```jsonc
    {
        "user": "The person's username",
        "secret": "The person's authentication token",
        "fields": [
            // Just an example of what fields could be requested
            "task",
            "link",
            "overdue"
        ]
    }
    ```

  * Response codes
     * 200 OK - The request proceeded without any errors
     * 400 Bad Request - The request sent to the microservice was malformed
     * 401 Unauthorized - Failed Google Classroom authentication
     * 403 Forbidden - Yikes. Banned from Google Classroom?
     * 500 Internal Server Error - The microservice hit an error

  * Response body (JSON)
    
    ```jsonc
    [
        {
            "task": "Task name",
            "subbject": "Subject name",
            "desc": "Task description",
            "link": "Task hyperlink",
            "duedate": [2021, 06, 23],
            "duetime": [23, 59],
            // Overdue status: 1 if overdue, 0 if not.
            "overdue": 1
        }
    ]
    ```