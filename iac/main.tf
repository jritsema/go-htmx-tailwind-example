provider "aws" {
  default_tags {
    tags = var.tags
  }
}

resource "aws_apprunner_service" "main" {
  service_name = var.app

  source_configuration {
    auto_deployments_enabled = true

    image_repository {
      image_repository_type = "ECR"
      image_identifier      = "${var.repository_url}:latest"
      image_configuration {
        port = var.port
        runtime_environment_variables = {
          "PORT" = var.port
        }
      }
    }

    authentication_configuration {
      access_role_arn = aws_iam_role.apprunner.arn
    }
  }

  instance_configuration {
    cpu    = var.cpu
    memory = var.memory
    # instance_role_arn = ""
  }

  health_check_configuration {
    protocol            = var.health_check_protocol
    path                = var.health_check
    timeout             = var.health_check_timeout
    healthy_threshold   = var.health_check_healthy_threshold
    interval            = var.health_check_interval
    unhealthy_threshold = var.health_check_unhealthy_threshold
  }
}

resource "aws_iam_role" "apprunner" {
  name               = var.app
  assume_role_policy = data.aws_iam_policy_document.apprunner.json

  # workaround for https://github.com/hashicorp/terraform-provider-aws/issues/6566
  provisioner "local-exec" {
    command = "sleep 15"
  }
}

data "aws_iam_policy_document" "apprunner" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = [
        "tasks.apprunner.amazonaws.com",
        "build.apprunner.amazonaws.com",
      ]
    }
  }
}

resource "aws_iam_role_policy_attachment" "apprunner" {
  role       = aws_iam_role.apprunner.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSAppRunnerServicePolicyForECRAccess"
}
