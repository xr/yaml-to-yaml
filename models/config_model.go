package models

type Config struct {
	Name             string                   `yaml:"name"`
	EnableEchoServer bool                     `yaml:"enableEchoServer"`
	Hosts            []string                 `yaml:"hosts"`
	Upstream         UpstreamConfig           `yaml:"upstream"`
	Headers          HeadersConfig            `yaml:"headers"`
	RateLimiters     map[string][]RateLimiter `yaml:"rateLimiters"`
	Routes           []Route                  `yaml:"routes"`
	Debug            bool                     `yaml:"debug"`
}

type UpstreamConfig struct {
	Host          string `yaml:"host"`
	SupportsHttps bool   `yaml:"supportsHttps"`
}

type HeadersConfig struct {
	In []string `yaml:"in"`
}

type RateLimiter struct {
	Targets []struct {
		Key string `yaml:"key"`
	} `yaml:"targets"`
	Limits []struct {
		Unit  string `yaml:"unit"`
		Limit int    `yaml:"limit"`
	} `yaml:"limits"`
}

type Route struct {
	Path           string `yaml:"path"`
	Authentication bool   `yaml:"authentication"`
	Filters        struct {
		DisableAuthorization bool `yaml:"disableAuthorization"`
	} `yaml:"filters"`
	Methods      []string      `yaml:"methods"`
	Name         string        `yaml:"name"`
	Headers      HeadersConfig `yaml:"headers"`
	RateLimiters []string      `yaml:"rateLimiters"`
}
