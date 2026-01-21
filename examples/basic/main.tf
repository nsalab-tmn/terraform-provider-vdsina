# examples/basic/main.tf
# Minimal VPS server creation example

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

# Get available datacenters
data "vdsina_datacenters" "all" {}

# Get server groups
data "vdsina_server_groups" "all" {}

# Get server plans for Standard servers
data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

# Get OS templates
data "vdsina_templates" "all" {}

# Create a minimal server
resource "vdsina_server" "web" {
  datacenter  = data.vdsina_datacenters.all.datacenters[0].id
  server_plan = data.vdsina_server_plans.standard.plans[0].id
  template    = data.vdsina_templates.all.templates[0].id
  name        = "my-first-server"
}

# Outputs defined in outputs.tf
