#!/bin/bash

eval $(minikube docker-env)
docker build -t prom:latest -f Dockerfile.prod .
kubectl delete -f deployment.yml
kubectl apply -f deployment.yml
