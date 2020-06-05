/*
Package cors provides middleware for net/http package to handle CORS requests according to http://www.w3.org/TR/cors/

There are two scenarios: use http.HandleFunc and use http.NewServeMux.

If you use http.HandleFunc you should:

1. Create configuration first. Example:

  config := &cors.Config{
		AllowAllOrigin: false,
		AllowOriginPattern: "^https?://localhost(:[0-9]+)?$",
		AllowMethods: []string{"DELETE"},
	}

2. Create CORS middleware. Pass configuration.

	corsMiddleware := cors.CreateMiddleware(config)

3. Imagine you have some handler function:

	yourProjectHandler := func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`I am handler in your project`))
	}

4. Convert it to http.HandlerFunc and pass to corsMiddleware function.

	http.HandleFunc("/entry", corsMiddleware(http.HandlerFunc(yourProjectHandler)))

5. Start your server.

  http.ListenAndServe(":8080", nil)

If you use http.NewServeMux you should:

1. Create configuration first. Example:

  config := &cors.Config{
		AllowAllOrigin: false,
		AllowOriginPattern: "^https?://localhost(:[0-9]+)?$",
		AllowMethods: []string{"DELETE"},
	}

2. Create CORS middleware. Pass configuration.

	corsMiddleware := cors.CreateMiddleware(config)

3. Imagine you have some handler function:

  yourProjectHandler := func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`I am handler in your project`))
	}

4. Create mux and add handler function to it.

  mux := http.NewServeMux()
	mux.HandleFunc("/entry", yourProjectHandler)

5. Start your server.

  http.ListenAndServe(":8080", corsMiddleware(mux))
 */
package cors
