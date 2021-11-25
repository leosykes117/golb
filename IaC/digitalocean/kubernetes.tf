resource "kubernetes_namespace" "golb" {
    metadata {
        name = "${var.k8s_app_namespace}"
    }
}

resource "kubernetes_secret" "do_registry_credentials" {
    metadata {
        namespace = "${var.k8s_app_namespace}"
        name = "docker-cfg"
    }

    data = {
        ".dockerconfigjson" = digitalocean_container_registry_docker_credentials.golb_repository.docker_credentials
    }

    type = "kubernetes.io/dockerconfigjson"

    depends_on = [
        kubernetes_namespace.golb
    ]
}