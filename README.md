# OpenTelemetry Collector Cloud Foundry Buildpack

:warning: This buildpack is still under development.

This buildpack can be added to Cloud Foundry apps to run the OpenTelemetry Collector as a side-car.

## Usage

When pushing your app, just add the buildpack to the process, for example:

```
cf push -b https://github.com/maxmoehl/otelcol-buildpack -b go_buildpack .
```
