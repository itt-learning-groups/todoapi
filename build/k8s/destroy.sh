#!/bin/bash

ENV=$1

kubectl delete secrets,configmaps,deployments,services -l app=todoapi -n "${ENV}"
kubectl delete namespace "${ENV}"
