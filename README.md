classroom-svc
=============
TaskCollect microservice for interacting with Google Classroom.

Description
-----------

`classroom-svc` is a microservice which runs an HTTP server on port 2000. Its
purpose is to retrieve information for a student's assignments from Google
Classroom. Requests and responses are done via the HTTP server, with their
format detailed in the [HTTP Spec](#http-spec).

The microservice uses the following Go libraries provided by Google for easier
communication with the Google Classroom API:

  * `google.golang.org/api/classroom/v1`
  * `golang.org/x/oauth2/google`

Not that it really matters, as dependencies are automatically managed.

HTTP Spec
---------

* GET /v1/tasks
  * Request body (JSON)

	```jsonc
	{
		"user": "The person's student ID. No 'CURRIC\\'.",
		"secret": "The person's authentication token",
		"fields": [
			/*
				Just an example of what fields could be requested.
				If all fields are required, simply omit the "fields" field from
				the request.
			*/
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
			"class": "Class name",
			"desc": "Task description",
			"link": "Task hyperlink",
			"res": ["reslink1", "reslink2"],
			"duedate": /* Timestamp. */,
			"overdue": false
		}
	]
	```