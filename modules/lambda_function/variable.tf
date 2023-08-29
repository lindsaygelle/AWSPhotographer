variable "architectures" {
  default   = null
  sensitive = false
  type      = list(string)
}

variable "code_signing_config_arn" {
  default   = null
  sensitive = false
  type      = string
}

variable "dead_letter_config" {
  default   = null
  sensitive = false
  type = object({
    target_arn = string
  })
}

variable "description" {
  default   = null
  sensitive = false
  type      = string
}

variable "environment" {
  default   = null
  sensitive = false
  type      = map(string)
}

variable "ephemeral_storage" {
  default   = null
  sensitive = false
  type = object({
    size = number
  })
}

variable "event_invoke_destination_config" {
  default   = null
  sensitive = false
  type = object({
    on_failure = object({
      destination_arn = string
    })
    on_success = object({
      destination_arn = string
    })
  })
}

variable "file_system_config" {
  default   = null
  sensitive = false
  type = object({
    arn              = string
    local_mount_path = string
  })
}

variable "filename" {
  default   = null
  sensitive = false
  type      = string
}


variable "function_name" {
  sensitive = false
  type      = string
}

variable "handler" {
  default   = null
  sensitive = false
  type      = string
}

variable "image_config" {
  default   = null
  sensitive = false
  type = object({
    command           = string
    entry_point       = string
    working_directory = string
  })
}

variable "image_uri" {
  default   = null
  sensitive = false
  type      = string
}

variable "kms_key_arn" {
  default   = null
  sensitive = false
  type      = string
}

variable "layers" {
  default   = null
  sensitive = false
  type      = list(string)
}

variable "maximum_event_age_in_seconds" {
  default   = null
  sensitive = false
  type      = number
}

variable "maximum_retry_attempts" {
  default   = null
  sensitive = false
  type      = number
}

variable "memory_size" {
  default   = null
  sensitive = false
  type      = number
}

variable "package_type" {
  default   = null
  sensitive = false
  type      = string
}

variable "reserved_concurrent_executions" {
  default   = null
  sensitive = false
  type      = number
}

variable "role_arn" {
  sensitive = false
  type      = string
}

variable "runtime" {
  default   = null
  sensitive = false
  type      = string
}

variable "s3_bucket" {
  default   = null
  sensitive = false
  type      = string
}

variable "s3_key" {
  default   = null
  sensitive = false
  type      = string
}

variable "s3_object_version" {
  default   = null
  sensitive = false
  type      = string
}

variable "skip_destroy" {
  default   = true
  sensitive = false
  type      = bool
}

variable "source_code_hash" {
  default   = null
  sensitive = false
  type      = string
}

variable "snap_start" {
  default   = null
  sensitive = false
  type = object({
    apply_on = string
  })
}

variable "tags" {
  default   = null
  sensitive = false
  type      = map(string)
}

variable "timeout" {
  default   = 1
  sensitive = false
  type      = number
}

variable "tracing_config" {
  default   = null
  sensitive = false
  type = object({
    mode = string
  })
}

variable "vpc_config" {
  default   = null
  sensitive = false
  type = object({
    security_group_ids = list(string)
    subnet_ids         = list(string)
  })
}
