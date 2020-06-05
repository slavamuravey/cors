package cors

// Config is structure contains all necessary configurations to handle CORS requests
type Config struct {
  // AllowAllHeaders allows all headers which presented in Access-Control-Request-Headers
  AllowAllHeaders            bool
  // AllowAllOrigin allows requests from all origins
  AllowAllOrigin             bool
  // AllowCredentials sets Access-Control-Allow-Credentials header
  AllowCredentials           bool
  // AllowHeaders is list of allowed headers for request
  AllowHeaders               []string
  // AllowMethods is list of allowed methods for request
  AllowMethods               []string
  // AllowOriginPattern is regexp pattern (ex. ^https?://localhost(:[0-9]+)?$) that is used on origin validation stage.
  AllowOriginPattern         string
  // ContinuousPreflight determines if preflight request must be terminated or it must go next
  ContinuousPreflight        bool
  // ExposedHeaders is list of headers presented in Access-Control-Expose-Headers header
  ExposedHeaders             []string
  // MaxAge sets Access-Control-Max-Age header
  MaxAge                     int
  // PreflightTerminationStatus sets status code of response for preflight request
  PreflightTerminationStatus int
}
