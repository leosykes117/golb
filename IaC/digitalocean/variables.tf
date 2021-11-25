variable "do_token" {}

variable "project_name" {
    description = "Nombre del proyecto"
    type = string
    default = "golb"
}

variable "machine_size" {
    description = "Tamaño de los droplets para el cluster de kubernetes"
    type = string
    default = "s-2vcpu-2gb"
}

variable "node_count" {
    description = "Numero de instancias droplet para los nodos del cluster de kubernetes"
    type = number
    default = 1
}

variable "k8s_app_namespace" {
    description = "Nombre del namespace donde se crearan los recursos dentro del cluster"
    type = string
    default = "golb-test"
}