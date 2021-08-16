#!/bin/bash

ENV=$1

kubectl delete secrets,configmaps,deployments,services -l app=todoapi -l env="${ENV}" -n default
