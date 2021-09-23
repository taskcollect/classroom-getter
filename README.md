classroom-svc
=============
TaskCollect microservice for interacting with Google Classroom.

Developer notes
---------------
Only just setup the development environment; don't expect any functionality.

All platforms which use Google authentication need to be treated separately. As we need to use the Google-preferred way of authentication (with the Google auth dialog and whatnot), the TaskCollect user manager needs to be built in a way that allows this. We shall find out how to do this during `classroom-svc` development.

Another important aspect that Google allows for (and promotes) is working with partial resources. The Google API supports requesting certain data fields to minimise the amount of overhead, thus *significantly* increasing performance. More on this here: https://developers.google.com/classroom/guides/performance