```
$ curl http://localhost:8080/timestamp -X POST
$ curl http://localhost:8080/metrics -s | grep last_posted_unixtime_seconds
# HELP last_posted_unixtime_seconds The last posted Unix time in seconds
# TYPE last_posted_unixtime_seconds gauge
last_posted_unixtime_seconds 1.7383421489568262e+09
```

```
docker build -t simple-last-time-api .
```
