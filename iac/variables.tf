variable "app" {
  type        = string
  description = "Name of the application"
}

variable "tags" {
  type        = map(string)
  description = "A map of the tags to apply to various resources"
  default     = {}
}

variable "repository_url" {
  description = "ECR repository URL"
  type        = string
}

variable "health_check" {
  description = "healthcheck path"
  type        = string
  default     = "/health"
}

variable "port" {
  description = "port"
  type        = string
  default     = "8080"
}

variable "cpu" {
  description = "cpu"
  type        = number
  default     = 1024
}

variable "memory" {
  description = "memory"
  type        = number
  default     = 2048
}

variable "health_check_protocol" {
  description = "The IP protocol that App Runner uses to perform health checks for your service."
  type        = string
  default     = "HTTP"
}

variable "health_check_healthy_threshold" {
  description = "The number of consecutive checks that must succeed before App Runner decides that the service is healthy."
  type        = number
  default     = 1
}

variable "health_check_interval" {
  description = "The time interval, in seconds, between health checks."
  type        = number
  default     = 5
}

variable "health_check_unhealthy_threshold" {
  description = "The number of consecutive checks that must fail before App Runner decides that the service is unhealthy. "
  type        = number
  default     = 5
}

variable "health_check_timeout" {
  description = "The time, in seconds, to wait for a health check response before deciding it failed."
  type        = number
  default     = 2
}
