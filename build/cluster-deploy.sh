#!/bin/bash

# preserve your work-day kubeconfig if you have one
mv ~/.kube/config ~/.kube/work_config

eksctl create cluster \
    --name learnanddevops \
    --version 1.19 \
    --region us-west-2 \
    --nodegroup-name worker-nodes \
    --node-type t3.small \
    --nodes 3 --nodes-min 1 --nodes-max 4 \
    --managed

# create a kubeconfig for the new cluster
aws eks update-kubeconfig --name learnanddevops --region us-west-2
