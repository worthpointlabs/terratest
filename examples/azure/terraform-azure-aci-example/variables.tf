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

variable "aci_name" {
  description = "The name to set for the ACI."
  default     = "acitest"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  default     = "aci-rg"
}

variable "location" {
  description = "The location to set for the ACR."
  default     = "Central US"
}
