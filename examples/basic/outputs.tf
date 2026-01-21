# examples/basic/outputs.tf

output "server_id" {
  description = "Created server ID"
  value       = vdsina_server.web.id
}

output "server_name" {
  description = "Server name"
  value       = vdsina_server.web.name
}

output "server_ip" {
  description = "Server public IP"
  value       = vdsina_server.web.ip
}

output "datacenter" {
  description = "Datacenter name"
  value       = vdsina_server.web.datacenter_name
}

output "plan" {
  description = "Server plan"
  value       = vdsina_server.web.plan_name
}

output "template" {
  description = "OS template"
  value       = vdsina_server.web.template_name
}

output "status" {
  description = "Server status"
  value       = vdsina_server.web.status
}
