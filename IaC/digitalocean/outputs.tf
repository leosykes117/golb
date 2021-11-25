output "k8s_versions" {
    value = data.digitalocean_kubernetes_versions.this.valid_versions
}

output "do_regions" {
    value = [
        for i, v in data.digitalocean_regions.available.regions : tomap(
            {
                slug = v.slug
                name = v.name
            }
        )
    ]
}

output "do_registry" {
    value = tomap(
        {
            name = digitalocean_container_registry.golb_repository.name
            endpoint = digitalocean_container_registry.golb_repository.endpoint
            server_url = digitalocean_container_registry.golb_repository.server_url
        }
    )
}

output "k8s_app_namespace" {
    value = var.k8s_app_namespace
}