apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: {{ .Name }}-ratelimiter
spec:
  workloadSelector:
    labels:
      ingress-type: namespace
  configPatches:
    - applyTo: HTTP_ROUTE
      match:
        context: GATEWAY
        routeConfiguration:
          vhost:
            route:
              name: name