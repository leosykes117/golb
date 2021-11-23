resource "digitalocean_kubernetes_cluster" "k8s_cluster" {
    name = "${var.project_name}-cluster"
    region = "nyc1"
    version = data.digitalocean_kubernetes_versions.this.latest_version

    node_pool {
        name = "${var.project_name}-k8s-nodes"
        size = var.machine_size
        node_count = var.node_count
    }
    depends_on = [
        digitalocean_container_registry.golb_repository
    ]
}

data "digitalocean_kubernetes_cluster" "k8s_cluster" {
    name = digitalocean_kubernetes_cluster.k8s_cluster.name
}

resource "local_file" "k8s_config" {
    content = data.digitalocean_kubernetes_cluster.k8s_cluster.kube_config.0.raw_config
    filename = "${path.module}/kubernetes/kubeconfig.yaml"
}
