package terraform

// Terraform has a number of lovely errors.  We look for the presence of these substrings in Terraform output to detect them.
const TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY 		= "diffs didn't match during apply"
const TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY_MSG 	= "This usually indicates a minor Terraform timing bug (https://github.com/hashicorp/terraform/issues/5200) that goes away when you reapply. Retrying terraform apply."