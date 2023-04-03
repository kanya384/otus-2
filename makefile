
up: install-nginx-ingress apply-manifests

apply-manifests:
	@kubectl apply -f .

install-nginx-ingress:
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && helm repo update && helm install ingress-nginx ingress-nginx/ingress-nginx