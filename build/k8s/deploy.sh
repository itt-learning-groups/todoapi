#!/bin/bash

# Get the deployment environment ("dev", "qa", or "prod") from a command-line argument
ENV=$1

# Deploy the environment `namespace`
kubectl apply -f ./"${ENV}"/namespace.yaml

# Deploy the `secret` that will hold the DB username/password in this namespace
kubectl create secret generic todoapi-configs --from-file="${ENV}"/secrets -n "${ENV}"
kubectl label secret todoapi-configs app=todoapi env=dev -n "${ENV}"

# Deploy the `config-map` that will hold the other DB config values in this namespace
kubectl apply -f ./"${ENV}"/config.yaml -n "${ENV}"

# Deploy the `deployment` and `service` in this namespace
kubectl apply -f ./service.yaml -n "${ENV}"
kubectl apply -f ./deployment.yaml -n "${ENV}"

# Wait a few seconds for the service loadbalancer hostname to be available
sleep 5

# Print out the service hostname so we can test it (browser, curl, Postman) for this namespace
printf "\nloadbalancer hostname: "
kubectl get service todoapi -n "${ENV}" -o jsonpath="{.status.loadBalancer.ingress[*].hostname}"
echo ""
