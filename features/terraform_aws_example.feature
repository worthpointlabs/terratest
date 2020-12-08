Feature: Terraform AWS Example

  Scenario: Run a simple test on AWS
    Given the Terraform module at "./examples/terraform-aws-example"
    And an input variable named "instance_name" with the value "test-instance-1"
    And an environment variable named "AWS_DEFAULT_REGION" with the value "us-east-1"
    When I run "terraform apply"
    Then the "instance_id" output is "Hello, World!"
