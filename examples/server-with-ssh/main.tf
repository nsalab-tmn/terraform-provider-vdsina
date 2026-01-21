# examples/server-with-ssh/main.tf
# Server with SSH key for secure access

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

# Create SSH key
resource "vdsina_ssh_key" "deploy" {
  name = "deploy-key"
  data = var.ssh_public_key
}

# Create server with SSH key
resource "vdsina_server" "app" {
  datacenter  = var.datacenter_id
  server_plan = var.server_plan_id
  template    = var.template_id
  ssh_key     = vdsina_ssh_key.deploy.id
  name        = var.server_name
  host        = var.hostname
}

output "server_ip" {
  description = "Server IP for SSH connection"
  value       = vdsina_server.app.ip
}

output "ssh_command" {
  description = "SSH command to connect"
  value       = "ssh root@${vdsina_server.app.ip}"
}
