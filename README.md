# Terraform Test Framework

This framework runs unit tests and integration tests for terraform modules.

## The Big Idea

Gruntwork customers rely on our Terraform Modules to build their infrastructure, so we need a high degree of confidence that our modules work as expected.

What would give us such confidence?

- The module should correctly `terraform apply` and `terraform destroy`.
- The module should not become "broken" as other updates are made to it or other modules.
- Each module should work in every AWS region.
- Once the module's infrastructure is created, functionality should work as expected.

The goal of this framework is to automate everything on the list above.  We want to make this automation fast by parallelizing tests, and we want to make it easy to write tests in the first place.

### Traditional Software Tests vs. DevOps Software Tests

In traditional software testing, we have the benefit of near total-control over our local environment.  If our code relies on third-party services, we can "mock" these services to simulate their responses so that our laptops can very quickly validate that our code works as expected.

In DevOps, we can't mock AWS.  It's simply too complex with too many interacting services.  So we accept this as a limitation of our environment, and take steps to introduce other properties of unit tests such as parallelizability, isolation, and fastness.

### Achieving Unit Tests for Our Terraform Modules

We approximate unit tests for our Terraform Modules with the following philosophy:

- Every AWS resource Terraform creates should be "namespaced" so that the same template running a unit test in parallel won't have a resource name conflict.
- We should randomize the AWS region in which we launch our resources to expose resource-specific quirks such different numbers of Availability Zones.
- We should auto-generate select resources such as EC2 KeyPairs since many Terraform Modules assume these are created "out of band".

By adopting this philosophy, terraform-test can run multiple "unit" tests in parallel, reducing our build times!

## How to Use
...

## ToDo
1. Add `circle.yml` to run tests automatically on each commit to `master`.
2. Fill in "How to Use"