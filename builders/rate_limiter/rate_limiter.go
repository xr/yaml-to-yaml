package rate_limiter

import (
	"fmt"

	"github.com/xr/yaml-to-yaml/types"
	"github.com/xr/yaml-to-yaml/utilities"
	"gopkg.in/yaml.v2"
)

func isRateLimiterExists(rateLimiters map[string][]types.RateLimiter, rateLimiterName string) bool {
	_, exists := rateLimiters[rateLimiterName]
	return exists
}

func GetRateLimits(rateLimiters map[string][]types.RateLimiter, route types.Route) []interface{} {
	rateLimits := []interface{}{}

	for _, rateLimiterName := range route.RateLimiters {
		rateLimits = append(rateLimits, GetRateLimit(rateLimiters, rateLimiterName)...)
	}

	return rateLimits
}

func GetRateLimit(rateLimiters map[string][]types.RateLimiter, rateLimiterName string) []interface{} {
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

			rateLimits = append(rateLimits, utilities.NewMap(
				"actions", actions,
			))
		}
	}

	return rateLimits
}

func NewConfigPatches(config *types.Config) []interface{} {
	configPatches := []interface{}{}
	for _, route := range config.Routes {
		configPatch := NewConfigPatch(config.RateLimiters, route)
		configPatches = append(configPatches, configPatch)
	}

	return configPatches
}

func NewConfigPatch(rateLimiters map[string][]types.RateLimiter, route types.Route) interface{} {
	rateLimits := GetRateLimits(rateLimiters, route)
	configPatch := utilities.NewMap(
		"applyTo", "HTTP_ROUTE",
		"match", utilities.NewMap(
			"context", "GATEWAY",
			"routeConfiguration", utilities.NewMap(
				"vhost", utilities.NewMap(
					"route", utilities.NewMap(
						"name", route.Name,
					),
				),
			),
		),
		"patch", utilities.NewMap(
			"operation", "MERGE",
			"value", utilities.NewMap(
				"route", utilities.NewMap(
					"rate_limits", rateLimits,
				),
			),
		),
	)

	return configPatch
}

func NewEnvoyFilter(config *types.Config) interface{} {
	configPatches := NewConfigPatches(config)
	envoyFilter := utilities.NewMap(
		"apiVersion", "networking.istio.io/v1alpha3",
		"kind", "EnvoyFilter",
		"metadata", utilities.NewMap(
			"name", config.Name+"-ratelimiter",
		),
		"spec", utilities.NewMap(
			"configPatches", configPatches,
			"workloadSelector", utilities.NewMap(
				"labels", utilities.NewMap(
					"ingress-type", "namespace",
				),
			),
		),
	)

	return envoyFilter
}

func Render(config *types.Config) (string, error) {

	envoyFilter := NewEnvoyFilter(config)

	yamlData, err := yaml.Marshal(envoyFilter)
	if err != nil {
		fmt.Printf("Error marshaling to YAML: %v\n", err)
	}

	return string(yamlData), nil
}
