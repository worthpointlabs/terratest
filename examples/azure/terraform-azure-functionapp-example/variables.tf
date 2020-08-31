# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

resource "random_string" "lower" {
  length  = 12
  upper   = false
  lower   = true
  number  = false
  special = false
}

variable "location" {
  description = "The location to set for the project"
  default     = "West Europe"
}

variable "project_name" {
  description = "Name of the project"
  default     = ""
}
