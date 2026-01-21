# examples/complete/variables.tf

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

variable "ssh_public_key" {
  type        = string
  description = "SSH public key for server access"
  sensitive   = true
}
