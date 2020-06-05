package cors

const (
  // AllowCredentialsHeader contains the Access-Control-Allow-Credentials response header that tells browsers
  // whether to expose the response to frontend JavaScript code
  // when the request's credentials mode (Request.credentials) is include.
  AllowCredentialsHeader string = "Access-Control-Allow-Credentials"
  // AllowOriginHeader contains the Access-Control-Allow-Origin response header that indicates whether the response
  // can be shared with requesting code from the given origin.
  AllowOriginHeader string = "Access-Control-Allow-Origin"
  // AllowHeadersHeader contains the Access-Control-Allow-Headers response header that is used in response to a
  // preflight request which includes the Access-Control-Request-Headers to indicate which HTTP headers can be used
  // during the actual request.
  AllowHeadersHeader string = "Access-Control-Allow-Headers"
  // AllowMethodsHeader contains the Access-Control-Allow-Methods response header that specifies the method or methods
  // allowed when accessing the resource in response to a preflight request.
  AllowMethodsHeader string = "Access-Control-Allow-Methods"
  // ExposeHeadersHeader contains the Access-Control-Expose-Headers response header that indicates which headers can be
  // exposed as part of the response by listing their names.
  ExposeHeadersHeader string = "Access-Control-Expose-Headers"
  // MaxAgeHeader contains the Access-Control-Max-Age response header that indicates how long the results of a preflight
  // request (that is the information contained in the
  // Access-Control-Allow-Methods and Access-Control-Allow-Headers headers) can be cached.
  MaxAgeHeader string = "Access-Control-Max-Age"
  // OriginHeader contains the Origin request header that indicates where a fetch originates from.
  // It doesn't include any path information, but only the server name.
  // It is sent with CORS requests, as well as with POST requests.
  // It is similar to the Referer header, but, unlike this header, it doesn't disclose the whole path.
  OriginHeader string = "Origin"
  // RequestHeadersHeader contains the Access-Control-Request-Headers request header that is used by browsers
  // when issuing a preflight request, to let the server know which HTTP headers the client might send when
  // the actual request is made.
  RequestHeadersHeader string = "Access-Control-Request-Headers"
  // RequestMethodHeader contains the Access-Control-Request-Method request header that is used by browsers when issuing
  // a preflight request, to let the server know which HTTP method will be used when the actual request is made.
  // This header is necessary as the preflight request is always an OPTIONS and doesn't use the same method
  // as the actual request.
  RequestMethodHeader string = "Access-Control-Request-Method"
  // VaryHeader contains the Vary HTTP response header that determines how to match future request headers to decide
  // whether a cached response can be used rather than requesting a fresh one from the origin server.
  // It is used by the server to indicate which headers it used when selecting a representation of a resource
  // in a content negotiation algorithm.
  VaryHeader string = "Vary"
)
