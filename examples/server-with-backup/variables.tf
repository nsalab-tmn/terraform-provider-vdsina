# examples/server-with-backup/variables.tf

variable "vdsina_api_token" {
  type        = string
  description = "VDSina API token"
  sensitive   = true
}

variable "vdsina_base_url" {
  type        = string
  description = "VDSina API base URL"
  default     = "https://userapi.vdsina.com/v1"
}

variable "datacenter_id" {
  type        = number
  description = "Datacenter ID"
  default     = 4  # Amsterdam
}

variable "server_plan_id" {
  type        = number
  description = "Server plan ID"
  default     = 1  # 2 RAM / 1 CPU / 40 NVMe
}

variable "template_id" {
  type        = number
  description = "OS template ID"
  default     = 23  # Ubuntu 24.04
}
