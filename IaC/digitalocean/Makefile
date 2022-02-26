all: terraform

CMD=
ARG=
VARS=
ENVFILE="./.env"

.PHONY: terraform
terraform:
	docker run -it --rm \
	-w /home/terraform/ \
	--env-file ${ENVFILE} \
	--mount type=bind,source=${PWD},destination=/home/terraform/ \
	--name golb-terraform \
	hashicorp/terraform ${CMD}

helm:
	docker run -ti --rm \
	-w /apps \
	-v $(PWD):/apps \
	-v ~/.kube:/root/.kube \
	-v ~/.helm:/root/.helm \
	-v ~/.config/helm:/root/.config/helm \
	-v ~/.cache/helm:/root/.cache/helm \
	alpine/helm ${CMD}

kubeseal:
	docker run -i --rm \
	-w /apps \
	-v $(PWD):/apps \
	kubeseal:0.16.0 < ${ARG} ${CMD}

# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo add ingress-nginx https://kubernetes.github.io/ingress-nginx"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo update"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace --set controller.publishService.enabled=true"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml uninstall ingress-nginx"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml upgrade --install ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --create-namespace"

### PROMETHEUS STACK ###
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo add prometheus-community https://prometheus-community.github.io/helm-charts"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo add stable https://charts.helm.sh/stable"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo update"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml install kube-prometheus --namespace monitoring --create-namespace -f kubernetes/monitoring/kube-prometheus/custom-values.yaml prometheus-community/kube-prometheus-stack"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml install mysql-exporter --namespace golb-test -f kubernetes/monitoring/kube-prometheus/mysqld-values.yaml prometheus-community/prometheus-mysql-exporter"
# ---
# kubectl --kubeconfig ./kubernetes/kubeconfig.yaml apply --server-side -f ~/Repositories/kubernetes/kube-prometheus/manifests/setup
# until kubectl --kubeconfig ./kubernetes/kubeconfig.yaml get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
# kubectl --kubeconfig ./kubernetes/kubeconfig.yaml apply -f ~/Repositories/kubernetes/kube-prometheus/manifests/

### KUBESEALED ###
# kubectl --kubeconfig ./kubernetes/kubeconfig.yaml apply -f https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.16.0/controller.yaml
