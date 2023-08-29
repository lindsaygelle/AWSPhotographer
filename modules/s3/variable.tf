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
