# OpenTelemetry Collector Cloud Foundry Buildpack

:warning: While functional, this buildpack is still experimental.

This buildpack can be added to Cloud Foundry apps to run the OpenTelemetry Collector as a sidecar.

## Usage

When pushing your app, just add this buildpack to the process, for example:

```
cf push -b https://github.com/maxmoehl/otelcol-buildpack -b go_buildpack .
```

The config is taken from the environment variable `OTELCOL_CONFIG`.

For a working example see the `example-app/`.

## Caveats

Sidecars can crash the entire application, so if the otelcol process crashes the entire application
will. The plan is to address this by having a wrapper script to catch such crashes and ensure they
don't bring down the entire application.

The limits imposed by environment variables will prevent more complex otelcol configurations. This
should be addressed by allowing the configuration to be read from the disk where such limits do not
apply.
