## Unity Gateway yaml-to-yaml

Simplify our Helm templating by embracing the flexibility of the programming language.

## Usage

Options:
```
-config string
      Path to the configuration file (default "config.yaml")
-controllers string
      Comma-separated list of controllers to run (default "rate_limiter")
-output string
      Path to the output folder (default "output")
-output-filename string
      Name of the output file (default "default.yaml")
```

```
./program --config config.yaml --output output --controllers rate_limiter --output-filename rate_limiter_actions.yaml
```
