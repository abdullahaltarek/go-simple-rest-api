# Go Simple CRUD REST API

Simple Go REST API for CRUD operations

No database is used, just in memory data structures. No 3rd party library was used.
This was the main motto for this project, fully implemented using Go standard libraries

Endpoints:
```
/people
/people/get/{id}
/people/create/{id}
/people/update/{id}
/people/delete/{id}
```

Also, a logger is implemented for logging incoming requests onto the console

The whole application is Dockerized.
Just build the docker image and run it.
Then you can access it from :8075 port for the above mentioned endpoints

This is my very first GO Api.

I am planning to develop API and full blown database driven application from next

PS: Pardon me for the .idea directory, I've added it to the .gitignore file. But it's not working for some strange reason.