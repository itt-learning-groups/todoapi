#!/bin/bash

ENV=$1

kubectl delete deployments,services -l app=todoapi -l env="${ENV}" -n default
