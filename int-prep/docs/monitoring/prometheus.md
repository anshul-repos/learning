Time series data model:

Collects and stores metrics as time series data 


    • Prometheus Server is the brain of the metric-based monitoring system. The main job of the server is to collect the metrics from various targets using pull model.
    
    • Target is the source where Prometheus scrape the metrics. A target could be servers, services, Kubernetes pods, application endpoints, etc.
    
    • The general term for collecting metrics from the targets using Prometheus is called scraping.
    
    • Prometheus periodically scrapes the metrics, based on the scrape interval that we mention in the Prometheus configuration file.
    
    • The metric data which prometheus receives changes over time (CPU, memory, network IO etc..). It is called time-series data. So Prometheus uses a Time Series Database (TSDB) to store all its data.
    • By default prometheus looks for metrics under /metrics
    • PromQL is a flexible query language that can be used to query time series metrics from the prometheus.
    • Exporter: fetches metrics from target for some components


Metrics:

    1. Help : description
    2. Type :  one of the three types (counter, guage, histogram)

Prometheus pulls data from endpoints unlike aws-cloud watch which pushes the data.

Exception: Short lived metrics uses pushgateway to push the data