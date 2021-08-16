#!/bin/bash

# Get the deployment environment ("dev", "qa", or "prod") from a command-line argument
ENV=$1

# Deploy the `secret` that will hold the DB username/password
kubectl create secret generic todoapi-configs-"${ENV}" --from-file="${ENV}"/secrets -n default
kubectl label secret todoapi-configs-"${ENV}" app=todoapi env=dev

# Deploy the `config-map` that will hold the other DB config values
kubectl apply -f ./"${ENV}"/config.yaml -n default

# Deploy the `deployment` and `service`
kubectl apply -f ./"${ENV}"/service.yaml -n default
kubectl apply -f ./"${ENV}"/deployment.yaml -n default

# Wait a few seconds for the service loadbalancer hostname to be available
sleep 5

# Print out the service hostname so we can test it (browser, curl, Postman)
printf "\nloadbalancer hostname: "
kubectl get service todoapi-"${ENV}" -n default -o jsonpath="{.status.loadBalancer.ingress[*].hostname}"
echo ""
