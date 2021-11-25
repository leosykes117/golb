resource "digitalocean_container_registry" "golb_repository" {
    name                   = local.registry_name
    subscription_tier_slug = "basic"
}

resource "digitalocean_container_registry_docker_credentials" "golb_repository" {
    depends_on = [
        digitalocean_container_registry.golb_repository
    ]
    registry_name = local.registry_name
}
