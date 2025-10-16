#!/bin/bash

eval $(minikube docker-env)
docker build -t prom:latest -f Dockerfile.prod .
kubectl delete -f deployment.yml 2>/dev/null
kubectl apply -f deployment.yml
