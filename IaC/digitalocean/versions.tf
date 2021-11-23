terraform {
    required_providers {
        digitalocean = {
            source = "digitalocean/digitalocean"
            version = ">= 2.0"
        }

        local = {
            source = "hashicorp/local"
            version = ">= 2.1.0"
        }

        kubernetes = {
            source  = "hashicorp/kubernetes"
            version = ">= 2.6.1"
        }
    }
    required_version = ">= 0.6"
}