# Deploying todoapi to kubernetes cluster

From /build/ directory...

* Deploy the cluster:

        ./cluster-deploy.sh

*Note that we're still relying on the "env" labels and corresponding postfixes on the metadata names of the deployment & service resources to differentiate them for each environment ("dev", "qa", and "prod"), since each is currently deployed to the same (default) namespace.*

* Deploy the `dev` environment to the cluster's `default` namespace:

        cd k8s
        ./deploy.sh dev

* Deploy the `qa` environment to the cluster's `default` namespace:

        ./deploy.sh qa

* Deploy the `prod` environment to the cluster's `default` namespace:

        ./deploy.sh prod

## Clean up

Run `./destroy.sh dev`, `./destroy.sh qa`, `./destroy.sh prod` in the `k8s` directory, and/or `./cluster-destroy.sh` in the parent `build` directory.
