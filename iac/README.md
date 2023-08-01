# App Runner Terraform

This repo uses [asdf](https://asdf-vm.com/) to manage the `terraform` CLI and the various other tools it depends upon.

```
 Choose a make command to run

  init     project initialization - install tools and register git hook
  checks   run all pre-commit checks
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_apprunner_service.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/apprunner_service) | resource |
| [aws_iam_role.apprunner](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.apprunner](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_policy_document.apprunner](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_app"></a> [app](#input\_app) | Name of the application | `string` | n/a | yes |
| <a name="input_cpu"></a> [cpu](#input\_cpu) | cpu | `number` | `1024` | no |
| <a name="input_health_check"></a> [health\_check](#input\_health\_check) | healthcheck path | `string` | `"/health"` | no |
| <a name="input_health_check_healthy_threshold"></a> [health\_check\_healthy\_threshold](#input\_health\_check\_healthy\_threshold) | The number of consecutive checks that must succeed before App Runner decides that the service is healthy. | `number` | `1` | no |
| <a name="input_health_check_interval"></a> [health\_check\_interval](#input\_health\_check\_interval) | The time interval, in seconds, between health checks. | `number` | `5` | no |
| <a name="input_health_check_protocol"></a> [health\_check\_protocol](#input\_health\_check\_protocol) | The IP protocol that App Runner uses to perform health checks for your service. | `string` | `"HTTP"` | no |
| <a name="input_health_check_timeout"></a> [health\_check\_timeout](#input\_health\_check\_timeout) | The time, in seconds, to wait for a health check response before deciding it failed. | `number` | `2` | no |
| <a name="input_health_check_unhealthy_threshold"></a> [health\_check\_unhealthy\_threshold](#input\_health\_check\_unhealthy\_threshold) | The number of consecutive checks that must fail before App Runner decides that the service is unhealthy. | `number` | `5` | no |
| <a name="input_memory"></a> [memory](#input\_memory) | memory | `number` | `2048` | no |
| <a name="input_port"></a> [port](#input\_port) | port | `string` | `"8080"` | no |
| <a name="input_repository_url"></a> [repository\_url](#input\_repository\_url) | ECR repository URL | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of the tags to apply to various resources | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_service_url"></a> [service\_url](#output\_service\_url) | app runner endpoint |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
