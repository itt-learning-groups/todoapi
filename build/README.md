# Deploying todoapi to a Kubernetes cluster in AWS EKS

* Deploy the cluster:

      cd build/k8s-cluster-setup
      ./cluster-deploy.sh

* Build the Dockerfile and push the image to the private image repository of your choice (which matches the Docker login creds set in the following step).

* Set your image-repository and DB credentials as environment variables so you can use them in `--set` options when running the helm install/upgrade:

      export DOCKER_SERVER=...
      export DOCKER_USERNAME=...
      export DOCKER_PSWD=...
      export DOCKER_EMAIL=...
      export DB_USERNAME=...
      export DB_PSWD_DEV=...
      export DB_PSWD_QA=...
      export DB_PSWD_PROD=...

* Deploy the `dev` environment:

      cd build/helm/chart_from_helm_starter
      helm upgrade --install todoapi ./todoapi \
        --create-namespace \
        --set 'ingress.hosts.host=todoapi.dev.ittlearninggroups.com' \
        --set 'ingress.tls.hosts={todoapi.dev.ittlearninggroups.com}' \
        --set 'db.name.value=todoapidev' \
        --set "imageCredentials.registry=${DOCKER_SERVER}" \
        --set "imageCredentials.username=${DOCKER_USERNAME}" \
        --set "imageCredentials.password=${DOCKER_PSWD}" \
        --set "imageCredentials.email=${DOCKER_EMAIL}" \
        --set "dbCredentials.dbusername=${DB_USERNAME}" \
        --set "dbCredentials.dbpswd=${DB_PSWD_DEV}" \
        -n dev

* Deploy the `qa` environment:

      cd build/helm/chart_from_helm_starter
      helm upgrade --install todoapi ./todoapi \
        --create-namespace \
        --set 'ingress.hosts.host=todoapi.qa.ittlearninggroups.com' \
        --set 'ingress.tls.hosts={todoapi.qa.ittlearninggroups.com}' \
        --set 'db.name.value=todoapiqa' \
        --set "imageCredentials.registry=${DOCKER_SERVER}" \
        --set "imageCredentials.username=${DOCKER_USERNAME}" \
        --set "imageCredentials.password=${DOCKER_PSWD}" \
        --set "imageCredentials.email=${DOCKER_EMAIL}" \
        --set "dbCredentials.dbusername=${DB_USERNAME}" \
        --set "dbCredentials.dbpswd=${DB_PSWD_QA}" \
        -n qa

* Deploy the `prod` environment:

      cd build/helm/chart_from_helm_starter
      helm upgrade --install todoapi ./todoapi \
        --create-namespace \
        --set 'ingress.hosts.host=todoapi.qa.ittlearninggroups.com' \
        --set 'ingress.tls.hosts={todoapi.qa.ittlearninggroups.com}' \
        --set 'db.name.value=todoapiqa' \
        --set "imageCredentials.registry=${DOCKER_SERVER}" \
        --set "imageCredentials.username=${DOCKER_USERNAME}" \
        --set "imageCredentials.password=${DOCKER_PSWD}" \
        --set "imageCredentials.email=${DOCKER_EMAIL}" \
        --set "dbCredentials.dbusername=${DB_USERNAME}" \
        --set "dbCredentials.dbpswd=${DB_PSWD_QA}" \
        -n prod

## Clean up

* Clean up the `dev` environment:

      helm delete todoapi -n dev

* Clean up the `qa` environment:

      helm delete todoapi -n qa

* Clean up the `prod` environment:

      helm delete todoapi -n prod

* Clean up the cluster:

      cd build/k8s-cluster-setup
      ./cluster-destroy.sh
