output "service_url" {
  description = "app runner endpoint"
  value       = "https://${aws_apprunner_service.main.service_url}"
}
