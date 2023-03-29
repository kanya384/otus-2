
up: install-nginx-ingress apply-manifests

apply-manifests:
	@kubectl apply -f .

install-nginx-ingress:
	@helm repo add nginx-stable https://helm.nginx.com/stable && helm repo update && helm install nginx-ingress-controller nginx-stable/nginx-ingress --set controller.service.httpPort.port=80 --set controller.enableSnippets=True