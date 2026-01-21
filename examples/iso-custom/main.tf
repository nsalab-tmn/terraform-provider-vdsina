# examples/iso-custom/main.tf
# Custom ISO image for installation

terraform {
  required_providers {
    vdsina = {
      source  = "scinfra-pro/vdsina"
      version = "~> 0.2.0"
    }
  }
}

provider "vdsina" {
  api_token = var.vdsina_api_token
  base_url  = var.vdsina_base_url
}

# Upload ISO image
resource "vdsina_iso" "custom" {
  url = var.iso_url
}

# Create server with ISO
resource "vdsina_server" "custom_os" {
  datacenter  = var.datacenter_id
  server_plan = var.server_plan_id
  iso         = vdsina_iso.custom.id
  name        = "custom-os-server"
}

# Note: after OS installation you need to detach ISO
# resource "vdsina_server_iso" "detach" {
#   server_id = vdsina_server.custom_os.id
#   iso_id    = 0  # 0 = detach
# }

output "iso_id" {
  description = "ISO image ID"
  value       = vdsina_iso.custom.id
}

output "iso_name" {
  description = "ISO file name"
  value       = vdsina_iso.custom.name
}

output "iso_md5" {
  description = "ISO MD5 checksum"
  value       = vdsina_iso.custom.file_md5
}

output "server_ip" {
  description = "Server IP"
  value       = vdsina_server.custom_os.ip
}
