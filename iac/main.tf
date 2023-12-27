terraform {
  required_providers {
    glesys = {
      source = "glesys/glesys"
      # version = "~> 0.4.6" # Här kan du ändra om du vill köra en särskild version av providern.
    }
  }
}

variable "glesys_api_key" {
  type        = string
  description = ""
}

provider "glesys" {
  userid = "CL44666"
  token  = var.glesys_api_key
}

resource "glesys_ip" "ssm_ipv4" {
  address = "188.126.93.221"
  version = "4"
  platform = "KVM"
  datacenter = "Stockholm"
}

resource "glesys_ip" "ssm_ipv6" {
  address = "2a00:1a28:1410:5::10ba"
  version = "6"
  platform = "KVM"
  datacenter = "Stockholm"
}

resource "glesys_server" "ssm_server" {
  count = 1
  datacenter = "Stockholm"
  memory = 2048
  storage = 40
  cpu = 2
  bandwidth = 100

  hostname = "ssm-website"
  description = "Säkerhets-SM:s hemsida (sakerhetssm.se)"

  ipv4_address = glesys_ip.ssm_ipv4.address
  ipv6_address = glesys_ip.ssm_ipv6.address

  platform = "KVM"
  template = "Ubuntu 22.04 LTS (Jammy Jellyfish)"

  cloudconfig = file("vm-config.yaml")
  
}
