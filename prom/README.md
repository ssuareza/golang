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

2. Set environment variables:

```sh
eval $(minikube docker-env)
```

3. Build the docker image:

```sh
docker build -t prom:latest -f Dockerfile.prod .
```

4. Deploy it!

```sh
kubectl apply -f deployment.yml
```

5. Test it!

Make a few requests to the "ping" endpoint:

```sh
kubectl exec -it deployment/prom -- sh -c "wget -O - http://localhost:2112/hello"
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
