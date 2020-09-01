# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "client_id" {
  description = "The Service Principal Client Id for AKS to modify Azure resources."
}
variable "client_secret" {
  description = "The Service Principal Client Password for AKS to modify Azure resources."
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "acr_name" {
  description = "The name to set for the ACR."
  default     = "acrtest"
}

variable "sku" {
  description = "SKU tier for the ACR."
  default     = "Premium"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  default     = "acr-rg"
}

variable "location" {
  description = "The location to set for the ACR."
  default     = "Central US"
}
