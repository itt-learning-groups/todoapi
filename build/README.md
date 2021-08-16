# Deploying todoapi to kubernetes cluster

From /build/ directory...

* Deploy the cluster:

        ./cluster-deploy.sh

*The current iteration of this `build` folder assumes that the DB username/password will be hard-coded into the `deployment.yaml` files (which is a terrible idea, but is our starting place for this exercise and would work if you were careful never to push the hard-coded values to the remote repo).*

* For each deployment environment ("dev", "qa", and "prod"), fill in or replace the `DB_HOSTNAME`, `DB_DBNAME`, `DB_TODOS_COLLECTION`, `DB_USERNAME`, and `DB_PSWD` in the `spec.template.spec.containers.env` for the `todoapi` container in the `deployment.yaml` file for that environment (in the `k8s` directory)

*Note that we're using "app" and "env" labels in the `deployment.yaml` and `service.yaml` files: The "app" label will help differentiate these deployments from other apps we deploy to this k8s cluster. We're relying on the "env" label and corresponding postfixes on the metadata names of the deployment & service resources to differentiate them for each environment ("dev", "qa", and "prod"), since each is currently deployed to the same (default) namespace.*

* Deploy the `dev` environment to the cluster's `default` namespace:

        cd k8s
        ./deploy.sh dev

* Deploy the `qa` environment to the cluster's `default` namespace:

        ./deploy.sh qa

* Deploy the `prod` environment to the cluster's `default` namespace:

        ./deploy.sh prod

## Clean up

Run `./destroy.sh dev`, `./destroy.sh qa`, `./destroy.sh prod` in the `k8s` directory, and/or `./cluster-destroy.sh` in the parent `build` directory.
