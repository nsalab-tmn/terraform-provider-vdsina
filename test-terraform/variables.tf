# variables.tf

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
