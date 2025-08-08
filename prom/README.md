# prom

An example of how to publish metrics for Prometheus.

## Usage

Run:

```sh
go run main.go
```

and test it!

```sh
curl http://localhost:2112/ping
curl http://localhost:2112/metrics
```

## Minikube

Let's see how to add this service inside a Minikube kubernetes cluster.

This is assuming that you already have Prometheus running inside your kubernetes cluster with the Service Discovery enabled.

1. Start minikube:

```sh
minikube start
```

4. Deploy it!

```sh
bin/deploy.sh
```

5. Test it!

Make a few requests to the "ping" endpoint:

```sh
kubectl exec -it deployment/prom -- sh -c "wget -O - http://localhost:2112/ping"
```

And the metric should be published on **Prometheus**:

```sh
kubectl exec -it deployment/prom -- sh -c "wget -O - http://localhost:2112/metrics" | grep ping_request_count
```

```text
# HELP ping_request_count No of request handled by Ping handler
# TYPE ping_request_count counter
ping_request_count 6
```

## References

- <https://prometheus.io/docs/guides/go-application/>
