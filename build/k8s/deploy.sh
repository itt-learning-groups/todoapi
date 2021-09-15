#!/bin/bash

# Get the deployment environment ("dev", "qa", or "prod") from a command-line argument
ENV=$1
TODOAPI_SECRETS_PATH="${2:-../../../todoapi_secrets}"

# Deploy the environment `namespace`
kubectl apply -f ./"${ENV}"/namespace.yaml

# Deploy the `secret` that will hold the Docker-repository login creds
# (Requires DOCKER_SERVER, DOCKER_USERNAME, DOCKER_PSWD, and DOCKER_PSWD env vars)
kubectl create secret docker-registry dockerlogin --docker-server="${DOCKER_SERVER}" --docker-username="${DOCKER_USERNAME}" --docker-password="${DOCKER_PSWD}" --docker-email="${DOCKER_EMAIL}" -n "${ENV}"
kubectl label secret dockerlogin app=todoapi -n "${ENV}"

# Deploy the `secret` that will hold the DB username/password in this namespace
# (Requires TODOAPI_SECRETS_PATH and files "${TODOAPI_SECRETS_PATH}/${ENV}/secrets/dbusername" and "${TODOAPI_SECRETS_PATH}/${ENV}/secrets/dbpswd")
kubectl create secret generic todoapi-configs --from-file="${TODOAPI_SECRETS_PATH}/${ENV}/secrets" -n "${ENV}"
kubectl label secret todoapi-configs app=todoapi -n "${ENV}"

# Deploy the `config-map` that will hold the other DB config values in this namespace
kubectl apply -f ./"${ENV}"/config.yaml -n "${ENV}"

# Deploy the `deployment`s and `service`s in this namespace
kubectl apply -f ./service.yaml -n "${ENV}"
kubectl apply -f ./deployment.yaml -n "${ENV}"

# Deploy the `ingress` in this namespace
kubectl apply -f ./ingress.yaml -n "${ENV}"

# Wait a few seconds for the service loadbalancer hostname to be available
sleep 5

# Print out the service hostname so we can test it (browser, curl, Postman) for this namespace
printf "\nloadbalancer hostname: "
kubectl get ingress todoapi -n "${ENV}" -o jsonpath="{.status.loadBalancer.ingress[*].hostname}"
echo ""
