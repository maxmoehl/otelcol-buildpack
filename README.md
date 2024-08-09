# OpenTelemetry Collector Cloud Foundry Buildpack

:warning: While functional, this buildpack is still experimental.

This buildpack can be added to Cloud Foundry apps to run the OpenTelemetry Collector as a side-car.

## Usage

When pushing your app, just add this buildpack to the process, for example:

```
cf push -b https://github.com/maxmoehl/otelcol-buildpack -b go_buildpack .
```

The config is taken from the environment variable `OTELCOL_CONFIG`.

For a working example see the `example-app/`.
