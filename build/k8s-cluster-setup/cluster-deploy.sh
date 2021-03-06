#!/bin/bash

# preserve your work-day kubeconfig if you have one
mv ~/.kube/config ~/.kube/work_config

eksctl create cluster \
    --name learnanddevops2 \
    --version 1.19 \
    --region us-west-2 \
    --nodegroup-name worker-nodes \
    --node-type t3.small \
    --nodes 3 --nodes-min 1 --nodes-max 4 \
    --managed

# create a kubeconfig for the new cluster
aws eks update-kubeconfig --name learnanddevops2 --region us-west-2

# install the Contour ingress controller
kubectl apply -f https://j.hept.io/contour-deployment-rbac

# install the jetstack cert-manager
kubectl apply -f https://github.com/jetstack/cert-manager/releases/latest/download/cert-manager.yaml

# wait a few seconds for the cert-manager webhook to be ready
sleep 10

# install a TLS-cert ClusterIssuer that uses the Let's Encrypt prod server
kubectl apply -f letsencrypt.yaml
