#!/bin/bash

eksctl delete cluster learnanddevops2

# restore your work-day kubeconfig if you have one
rm ~/.kube/config
mv ~/.kube/work_config ~/.kube/config
