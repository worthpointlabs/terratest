# local-resources-only-example

A set of templates that create only *local* Terraform resources. That is, resources such as `template_file` that exist
only locally, as opposed to resources such as `aws_instance` that require an external provider such as AWS. When
testing Terraform wrappers, this allows the unit tests to run much faster than if we had to wait
on an external provider, such as waiting for AWS to create EC2 instances.