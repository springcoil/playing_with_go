# My first Go program
* I've become interested in Go for microservices - so this is my first one
* This repository is just put together to get myself to learn this stuff. 
* Some examples after running `go run microservice.go` using [`httpie`](http://radek.io/2015/10/20/httpie/)
* Peadars-MBP% http GET http://localhost:8081/John_Smith
HTTP/1.1 200 OK
Content-Length: 20
Content-Type: text/plain; charset=utf-8
Date: Tue, 20 Sep 2016 20:26:38 GMT

Welcome, John_Smith!

*Peadars-MBP% http GET http://localhost:8081/about.html
HTTP/1.1 200 OK
Content-Length: 20
Content-Type: text/plain; charset=utf-8
Date: Tue, 20 Sep 2016 20:26:46 GMT

Welcome, about.html!

* TODO: Add in your Reflections on Go. 
