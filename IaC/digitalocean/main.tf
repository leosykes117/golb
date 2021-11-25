data "digitalocean_kubernetes_versions" "this" {}

data "digitalocean_regions" "available" {
    filter {
        key    = "available"
        values = ["true"]
    }
}

resource "random_uuid" "project_id" {}

locals {
    k8s_cluster_name = "${var.project_name}-cluster-${random_uuid.project_id.result}"
    node_pool_name = "${var.project_name}-k8s-nodes-${random_uuid.project_id.result}"
    registry_name = "${var.project_name}-repository"
#    registry_name = "${var.project_name}-repository-${random_uuid.project_id.result}"
}
