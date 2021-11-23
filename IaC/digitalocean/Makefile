all: terraform

CMD=
VARS=
ENVFILE="./.env"

.PHONY: terraform
terraform:
	docker run -it --rm \
	-w /home/terraform/ \
	--env-file ${ENVFILE} \
	--mount type=bind,source=${PWD},destination=/home/terraform/ \
	--name golb-terraform \
	hashicorp/terraform ${CMD} ${VARS}

helm:
	docker run -ti --rm \
	-w /apps \
	-v $(PWD):/apps \
    -v ~/.kube:/root/.kube \
	-v ~/.helm:/root/.helm \
	-v ~/.config/helm:/root/.config/helm \
    -v ~/.cache/helm:/root/.cache/helm \
    alpine/helm ${CMD}

# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo add ingress-nginx https://kubernetes.github.io/ingress-nginx"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml repo update"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml install nginx-ingress ingress-nginx/ingress-nginx --set controller.publishService.enabled=true"
# make helm CMD="--kubeconfig ./kubernetes/kubeconfig.yaml uninstall nginx-ingress"