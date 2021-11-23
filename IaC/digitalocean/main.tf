data "digitalocean_kubernetes_versions" "this" {}

data "digitalocean_regions" "available" {
    filter {
        key    = "available"
        values = ["true"]
    }
}
