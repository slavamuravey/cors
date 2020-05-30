package cors

type Config struct {
  ContinuousPreflight    bool
  AllowOriginPattern     string
  AllowAllOrigin         bool
  AllowMethods           []string
  AllowHeaders           []string
  AllowAllHeaders        bool
  AllowCredentials       bool
  ExposedHeaders         []string
  MaxAge                 int
}
