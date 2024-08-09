# Example App

You can push this app as usual:

```
cf push -f ./manifest.yml
```

After it started you can observe the metrics from the logs via `cf logs otelcol-test`. To look at
our custom HTTP request metric you can use the following command:

```
cf logs otelcol-test | grep 'app_requests_total' -B 2 -A 9
```

Perform some requests and watch the metric go!
