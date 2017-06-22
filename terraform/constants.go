package terraform

// Terraform has a number of lovely errors.  We look for the presence of these substrings in Terraform output to detect them.
const TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY 		  = "diffs didn't match during apply"
const TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY_MSG 	  = "This usually indicates a minor Terraform timing bug (https://github.com/hashicorp/terraform/issues/5200) that goes away when you reapply. Retrying terraform apply."

// Based on the full error message: "error finding matching route for Route table (rtb-762ea111) and destination CIDR block (0.0.0.0/0)
// See https://github.com/hashicorp/terraform/issues/5335
const TF_ERROR_FINDING_MATCHING_ROUTE_FOR_ROUTE_TABLE     = "error finding matching route for Route table"
const TF_ERROR_FINDING_MATCHING_ROUTE_FOR_ROUTE_TABLE_MSG = "This is an eventual consistency issue with AWS where Terraform looks for the route that was just created but doesn't yet see it: https://github.com/hashicorp/terraform/issues/5335"


// Based on the full error message: "Resource 'aws_eip.nat' does not have attribute 'id' for variable 'aws_eip.nat.*.id'"
// See https://github.com/hashicorp/terraform/issues/5185
const TF_ERROR_DOES_NOT_HAVE_ATTRIBUTE_ID_FOR_VARIABLE     = "does not have attribute 'id' for variable"
const TF_ERROR_DOES_NOT_HAVE_ATTRIBUTE_ID_FOR_VARIABLE_MSG = "This is an eventual consistency issue with AWS where Terraform looks for an EIP that was just created but doesn't yet see it: https://github.com/hashicorp/terraform/issues/5185"

// Based on the full error message: "InvalidRouteTableID.NotFound: The routeTable ID 'rtb-2d0d2f4a' does not exist"
// See https://github.com/hashicorp/terraform/issues/7104
const TF_ERROR_INVALID_ROUTE_TABLE_ID                      = "InvalidRouteTableID.NotFound: The routeTable ID"
const TF_ERROR_INVALID_ROUTE_TABLE_ID_MSG                  = "This is an eventual consistency issue with AWS where Terraform looks for a Route Table that was just created but doesn't yet see it: https://github.com/hashicorp/terraform/issues/7104"

// Based on the full error message: "aws_network_acl_rule.private_app_subnet_all_traffic_from_public_subnet.0: Expected to find one Network ACL, got: []*ec2.NetworkAcl(nil)"
// See https://github.com/hashicorp/terraform/issues/5392
const TF_ERROR_EXPECTED_TO_FIND_ONE_NETWORK_ACL            = "Expected to find one Network ACL, got:"
const TF_ERROR_EXPECTED_TO_FIND_ONE_NETWORK_ACL_MSG        = "This is an eventual consistency issue with AWS where Terraform looks for a Network ACL that was just created but doesn't yet see it: https://github.com/hashicorp/terraform/issues/5392"

// Based on the full error message: "aws_subnet.private-persistence.2: InvalidSubnetID.NotFound: The subnet ID 'subnet-xxxxxxx' does not exist"
// See https://github.com/hashicorp/terraform/issues/6813#issuecomment-229142897
const TF_ERROR_INVALID_SUBNET_ID                           = "InvalidSubnetID.NotFound:"
const TF_ERROR_INVALID_SUBNET_ID_MSG                       = "This is an eventual consistency issue with AWS where Terraform looks for a Subnet ID that was just created but doesn't yet see it: https://github.com/hashicorp/terraform/issues/6813#issuecomment-229142897"

// Based on the full error message: "Error finding route after creating it: error finding matching route for Route table (rtb-xxxxx) and destination CIDR block (xxx.xxx.xxx.xxx/xxx)"
// See https://github.com/hashicorp/terraform/issues/8542
const TF_ERROR_FINDING_ROUTE_AFTER_CREATING                = "Error finding route after creating it:"
const TF_ERROR_FINDING_ROUTE_AFTER_CREATING_MSG            = "This is an eventual consistency issue with AWS where Terraform looks for a Route that was just created but doesn't yet see it: https://github.com/hashicorp/terraform/issues/8542"
