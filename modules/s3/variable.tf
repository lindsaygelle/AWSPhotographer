// aws_s3_bucket
variable "bucket_name" {
  sensitive = false
  type      = string
}

variable "bucket_prefix" {
  default   = null
  sensitive = false
  type      = string
}

variable "tags" {
  default   = null
  sensitive = false
  type      = map(string)
}

// aws_s3_bucket_intelligent_tiering_configuration
variable "intelligent_tiering_configuration" {
  default = {
    archive_access = {
      days = 125
    }
    deep_archive_access = {
      days = 180
    }
  }
  sensitive = false
  type = object({
    archive_access = object({
      days = number
    })
    deep_archive_access = object({
      days = number
    })
  })
}
