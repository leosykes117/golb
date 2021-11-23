resource "digitalocean_container_registry" "golb_repository" {
    name                   = "${var.project_name}-repository"
    subscription_tier_slug = "starter"
}

resource "digitalocean_container_registry_docker_credentials" "golb_repository" {
    depends_on = [
        digitalocean_container_registry.golb_repository
    ]
    registry_name = "${var.project_name}-repository"
}
