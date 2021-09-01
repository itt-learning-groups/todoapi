#!/bin/bash

ENV=$1

kubectl delete secrets,configmaps,deployments,services -l app=todoapi -n "${ENV}"
kubectl delete secrets,configmaps,deployments,services -l app=todoapi1 -n "${ENV}"
kubectl delete secrets,configmaps,deployments,services -l app=todoapi2 -n "${ENV}"
kubectl delete ingress todoapi -n "${ENV}"
kubectl delete namespace "${ENV}"
