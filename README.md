# Go CORS middleware [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/slavamuravey/cors) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/slavamuravey/cors/master/LICENSE) [![build](https://img.shields.io/travis/slavamuravey/cors.svg?style=flat)](https://travis-ci.org/slavamuravey/cors) [![codecov](https://codecov.io/gh/slavamuravey/cors/branch/master/graph/badge.svg)](https://codecov.io/gh/slavamuravey/cors)

<p align="center">
  <img
      alt="Go CORS middleware"
      src="https://user-images.githubusercontent.com/3774019/83374949-6c73cf00-a3f7-11ea-9caa-5fbfe018322d.png"
  />
</p>

This package provides middleware for Golang http server to support CORS requests.

## Getting Started

More and more web applications based on distributive architecture. 
Often frontend and backend are separated parts, so we have to
do cross-origin requests from frontend to backend. 
This package provides support for this requests for golang servers.

Let's create simple server with CORS support.

```go
package main

import (
	"github.com/slavamuravey/cors"
	"net/http"
)

func main() {
	config := &cors.Config{
		AllowAllOrigin: false,
		AllowOriginPattern: "^https?://localhost(:[0-9]+)?$",
		AllowMethods: []string{"DELETE"},
	}

	corsMiddleware := cors.CreateMiddleware(config)

	yourProjectHandler := func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`I am handler in your project`))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/entry", yourProjectHandler)

	http.ListenAndServe(":8080", corsMiddleware(mux))
}
```

```bash
$ curl -XDELETE -D - -H 'Origin: http://localhost:80' http://localhost:8080/entry
HTTP/1.1 200 OK
Access-Control-Allow-Origin: http://localhost:80
Vary: Origin
Vary: Access-Control-Request-Method
Vary: Access-Control-Request-Headers
Date: Fri, 05 Jun 2020 11:46:05 GMT
Content-Length: 28
Content-Type: text/plain; charset=utf-8

I am handler in your project
```

In this example every entry point will support CORS requests.

If you need to add CORS support for single entry point you can do it as in following example:

```go
package main

import (
	"github.com/slavamuravey/cors"
	"net/http"
)

func main() {
	config := &cors.Config{
		AllowAllOrigin: false,
		AllowOriginPattern: "^https?://localhost(:[0-9]+)?$",
		AllowMethods: []string{"DELETE"},
	}

	corsMiddleware := cors.CreateMiddleware(config)

	yourProjectHandler := func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`I am handler in your project`))
	}

	http.HandleFunc("/cors", corsMiddleware(http.HandlerFunc(yourProjectHandler)))
	http.HandleFunc("/nocors", yourProjectHandler)
	http.ListenAndServe(":8080", nil)
}
```

```bash
$curl -XDELETE -D - -H 'Origin: http://localhost:80' http://localhost:8080/cors
HTTP/1.1 200 OK
Access-Control-Allow-Origin: http://localhost:80
Vary: Origin
Vary: Access-Control-Request-Method
Vary: Access-Control-Request-Headers
Date: Fri, 05 Jun 2020 11:52:18 GMT
Content-Length: 28
Content-Type: text/plain; charset=utf-8

I am handler in your project

$ curl -XDELETE -D - -H 'Origin: http://localhost:80' http://localhost:8080/nocors
HTTP/1.1 200 OK
Date: Fri, 05 Jun 2020 11:52:22 GMT
Content-Length: 28
Content-Type: text/plain; charset=utf-8

I am handler in your project
```

## Configuration
- **AllowAllHeaders** `bool` allows all headers which presented in Access-Control-Request-Headers
- **AllowAllOrigin** `bool` allows requests from all origins
- **AllowCredentials** `bool` sets Access-Control-Allow-Credentials header
- **AllowHeaders** `[]string` is list of allowed headers for request
- **AllowMethods** `[]string` is list of allowed methods for request
- **AllowOriginPattern** `string` is regexp pattern (ex. ^https?://localhost(:[0-9]+)?$) that is used on origin validation stage.
- **ContinuousPreflight** `bool` determines if preflight request must be terminated or it must go next
- **ExposedHeaders** `[]string` is list of headers presented in Access-Control-Expose-Headers header
- **MaxAge** `int` sets Access-Control-Max-Age header
- **PreflightTerminationStatus** `int` sets status code of response for preflight request

## Contributing

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/slavamuravey/cors/tags).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
