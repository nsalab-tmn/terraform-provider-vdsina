# examples/iso-custom/variables.tf

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

variable "iso_url" {
  type        = string
  description = "URL to download ISO from"
  default     = "https://boot.netboot.xyz/ipxe/netboot.xyz.iso"
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
