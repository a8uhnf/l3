# Prometheus Metrics API

Expose prometheus metrics. Various golang related metrics and web server's each endpoints related metrics exposed as well.
Endpoint related metrics consist of `Http Method`, `url endpoint`, `Latency` etc. 

- URL: `/metrics`
- Method: `GET`
- Success Response:

    **StatusCode:** 200
    **body:**
    
    ```text
    ...
    ... various golang related metrics
    ... Moreover l3 prefix metrics are related to various 
    ... endpoint metrics like latency, method, response code etc.
    ...
    # HELP l3 this histogram is for l3 api
    # TYPE l3 histogram
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.005"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.01"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.025"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.05"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.1"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.25"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="0.5"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="1"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="2.5"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="5"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="10"} 0
    l3_bucket{endpoint="/api/v1/location",method="POST",le="+Inf"} 1
    l3_sum{endpoint="/api/v1/location",method="POST"} 1.537534e+06
    l3_count{endpoint="/api/v1/location",method="POST"} 1
    l3_bucket{endpoint="/metrics",method="POST",le="0.005"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="0.01"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="0.025"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="0.05"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="0.1"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="0.25"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="0.5"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="1"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="2.5"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="5"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="10"} 0
    l3_bucket{endpoint="/metrics",method="POST",le="+Inf"} 1
    l3_sum{endpoint="/metrics",method="POST"} 38310
    l3_count{endpoint="/metrics",method="POST"} 1
    # HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
    # TYPE promhttp_metric_handler_requests_in_flight gauge
    promhttp_metric_handler_requests_in_flight 1
    # HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
    # TYPE promhttp_metric_handler_requests_total counter
    promhttp_metric_handler_requests_total{code="200"} 0
    promhttp_metric_handler_requests_total{code="500"} 0
    promhttp_metric_handler_requests_total{code="503"} 0
    ```
    

#Metrics Objective

`l3` prefixed metrics are related to our `l3` service. It contains information like, endpoint latency, endpoint request count, sum of endpoint latency etc. 

We could use these metrics to decide service level SLO. 
Like if we want to define a SLO that 95% of all the request will be served under 500ms, 
then we could use these metrics 
and define a alert if this percentage goes below 95%.

