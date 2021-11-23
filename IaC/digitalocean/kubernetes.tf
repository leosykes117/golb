resource "kubernetes_secret" "example" {
    metadata {
        name = "docker-cfg"
    }

    data = {
        ".dockerconfigjson" = digitalocean_container_registry_docker_credentials.golb_repository.docker_credentials
    }

    type = "kubernetes.io/dockerconfigjson"
}