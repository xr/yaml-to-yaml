package models

type RateLimit struct {
	Actions []interface{} `yaml:"actions"`
}

func isRateLimiterExists(rateLimiters map[string][]RateLimiter, rateLimiterName string) bool {
	_, exists := rateLimiters[rateLimiterName]
	return exists
}

func GetRateLimits(rateLimiters map[string][]RateLimiter, route Route) []interface{} {
	rateLimits := []interface{}{}

	for _, rateLimiterName := range route.RateLimiters {
		rateLimits = append(rateLimits, GetRateLimit(rateLimiters, rateLimiterName)...)
	}

	return rateLimits
}

func GetRateLimit(rateLimiters map[string][]RateLimiter, rateLimiterName string) []interface{} {
	rateLimits := []interface{}{}

	if !isRateLimiterExists(rateLimiters, rateLimiterName) {
		return nil
	}

	rateLimiterConfigs := rateLimiters[rateLimiterName]

	for _, rateLimiterConfig := range rateLimiterConfigs {
		for _, limit := range rateLimiterConfig.Limits {
			actions := []interface{}{}
			rateLimitAction := map[string]interface{}{
				"generic_key": map[string]interface{}{
					"descriptor_key":   "group",
					"descriptor_value": rateLimiterName + "-unit-" + limit.Unit,
				},
			}
			actions = append(actions, rateLimitAction)

			for _, target := range rateLimiterConfig.Targets {
				if target.Key == "ip" {
					remoteAddressAction := map[string]interface{}{
						"remote_address": map[string]interface{}{},
					}
					actions = append(actions, remoteAddressAction)
				} else {
					requestHeadersAction := map[string]interface{}{
						"request_headers": map[string]interface{}{
							"header_name":    target.Key,
							"descriptor_key": target.Key,
						},
					}
					actions = append(actions, requestHeadersAction)
				}
			}

			rateLimits = append(rateLimits, &RateLimit{
				Actions: actions,
			})
		}
	}

	return rateLimits
}

func NewConfigPatches(config *Config) []interface{} {
	configPatches := []interface{}{}
	for _, route := range config.Routes {
		configPatch := NewConfigPatch(config.RateLimiters, route)
		configPatches = append(configPatches, configPatch)
	}

	return configPatches
}

func NewConfigPatch(rateLimiters map[string][]RateLimiter, route Route) interface{} {
	rateLimits := GetRateLimits(rateLimiters, route)
	configPatch := NewMap(
		"applyTo", "HTTP_ROUTE",
		"match", NewMap(
			"context", "GATEWAY",
			"routeConfiguration", NewMap(
				"vhost", NewMap(
					"route", NewMap(
						"name", route.Name,
					),
				),
			),
		),
		"patch", NewMap(
			"operation", "MERGE",
			"value", NewMap(
				"route", NewMap(
					"rate_limits", rateLimits,
				),
			),
		),
	)

	return configPatch
}

func NewEnvoyFilter(config *Config) interface{} {
	configPatches := NewConfigPatches(config)
	envoyFilter := NewMap(
		"apiVersion", "networking.istio.io/v1alpha3",
		"kind", "EnvoyFilter",
		"metadata", NewMap(
			"name", config.Name+"-ratelimiter",
		),
		"spec", NewMap(
			"configPatches", configPatches,
			"workloadSelector", NewMap(
				"labels", NewMap(
					"ingress-type", "namespace",
				),
			),
		),
	)

	return envoyFilter
}
