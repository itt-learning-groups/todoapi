#!/bin/bash

ENV=$1

kubectl apply -f ./"${ENV}"/deployment.yaml -n default
kubectl apply -f ./"${ENV}"/service.yaml -n default
