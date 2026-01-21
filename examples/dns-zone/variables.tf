# examples/dns-zone/variables.tf

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

variable "domain_name" {
  type        = string
  description = "Domain name for DNS zone"
  default     = "example.com"
}

variable "default_ip" {
  type        = string
  description = "Default IP address for A records"
  default     = "78.40.199.136"
}
