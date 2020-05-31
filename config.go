package cors

type Config struct {
  AllowAllHeaders     bool
  AllowAllOrigin      bool
  AllowCredentials    bool
  AllowHeaders        []string
  AllowMethods        []string
  AllowOriginPattern  string
  ContinuousPreflight bool
  ExposedHeaders      []string
  MaxAge              int
}
