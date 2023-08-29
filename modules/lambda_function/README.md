# Lambda Function

<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_lambda_alias.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_alias) | resource |
| [aws_lambda_function.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |
| [aws_lambda_function_event_invoke_config.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function_event_invoke_config) | resource |
| [aws_lambda_permission.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission) | resource |
| [aws_lambda_provisioned_concurrency_config.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_provisioned_concurrency_config) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_architectures"></a> [architectures](#input\_architectures) | aws\_lambda\_function | `list(string)` | `null` | no |
| <a name="input_code_signing_config_arn"></a> [code\_signing\_config\_arn](#input\_code\_signing\_config\_arn) | n/a | `string` | `null` | no |
| <a name="input_dead_letter_config"></a> [dead\_letter\_config](#input\_dead\_letter\_config) | n/a | <pre>object({<br>    target_arn = string<br>  })</pre> | `null` | no |
| <a name="input_description"></a> [description](#input\_description) | n/a | `string` | `null` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `map(string)` | `null` | no |
| <a name="input_ephemeral_storage"></a> [ephemeral\_storage](#input\_ephemeral\_storage) | n/a | <pre>object({<br>    size = number<br>  })</pre> | `null` | no |
| <a name="input_event_invoke_destination_config"></a> [event\_invoke\_destination\_config](#input\_event\_invoke\_destination\_config) | aws\_lambda\_function\_event\_invoke\_config | <pre>object({<br>    on_failure = object({<br>      destination_arn = string<br>    })<br>    on_success = object({<br>      destination_arn = string<br>    })<br>  })</pre> | `null` | no |
| <a name="input_file_system_config"></a> [file\_system\_config](#input\_file\_system\_config) | n/a | <pre>object({<br>    arn              = string<br>    local_mount_path = string<br>  })</pre> | `null` | no |
| <a name="input_filename"></a> [filename](#input\_filename) | n/a | `string` | `null` | no |
| <a name="input_function_name"></a> [function\_name](#input\_function\_name) | n/a | `string` | n/a | yes |
| <a name="input_handler"></a> [handler](#input\_handler) | n/a | `string` | `null` | no |
| <a name="input_image_config"></a> [image\_config](#input\_image\_config) | n/a | <pre>object({<br>    command           = string<br>    entry_point       = string<br>    working_directory = string<br>  })</pre> | `null` | no |
| <a name="input_image_uri"></a> [image\_uri](#input\_image\_uri) | n/a | `string` | `null` | no |
| <a name="input_kms_key_arn"></a> [kms\_key\_arn](#input\_kms\_key\_arn) | n/a | `string` | `null` | no |
| <a name="input_layers"></a> [layers](#input\_layers) | n/a | `list(string)` | `null` | no |
| <a name="input_maximum_event_age_in_seconds"></a> [maximum\_event\_age\_in\_seconds](#input\_maximum\_event\_age\_in\_seconds) | n/a | `number` | `null` | no |
| <a name="input_maximum_retry_attempts"></a> [maximum\_retry\_attempts](#input\_maximum\_retry\_attempts) | n/a | `number` | `null` | no |
| <a name="input_memory_size"></a> [memory\_size](#input\_memory\_size) | n/a | `number` | `null` | no |
| <a name="input_package_type"></a> [package\_type](#input\_package\_type) | n/a | `string` | `null` | no |
| <a name="input_permission_action"></a> [permission\_action](#input\_permission\_action) | aws\_lambda\_permission | `string` | `null` | no |
| <a name="input_permission_event_source_token"></a> [permission\_event\_source\_token](#input\_permission\_event\_source\_token) | n/a | `string` | `null` | no |
| <a name="input_permission_function_url_auth_type"></a> [permission\_function\_url\_auth\_type](#input\_permission\_function\_url\_auth\_type) | n/a | `string` | `null` | no |
| <a name="input_permission_principal"></a> [permission\_principal](#input\_permission\_principal) | n/a | `string` | `null` | no |
| <a name="input_permission_principal_org_id"></a> [permission\_principal\_org\_id](#input\_permission\_principal\_org\_id) | n/a | `string` | `null` | no |
| <a name="input_permission_source_account"></a> [permission\_source\_account](#input\_permission\_source\_account) | n/a | `string` | `null` | no |
| <a name="input_permission_statement_id"></a> [permission\_statement\_id](#input\_permission\_statement\_id) | n/a | `string` | `null` | no |
| <a name="input_provisioned_concurrent_executions"></a> [provisioned\_concurrent\_executions](#input\_provisioned\_concurrent\_executions) | aws\_lambda\_provisioned\_concurrency\_config | `number` | `null` | no |
| <a name="input_reserved_concurrent_executions"></a> [reserved\_concurrent\_executions](#input\_reserved\_concurrent\_executions) | n/a | `number` | `null` | no |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | n/a | `string` | n/a | yes |
| <a name="input_runtime"></a> [runtime](#input\_runtime) | n/a | `string` | `null` | no |
| <a name="input_s3_bucket"></a> [s3\_bucket](#input\_s3\_bucket) | n/a | `string` | `null` | no |
| <a name="input_s3_key"></a> [s3\_key](#input\_s3\_key) | n/a | `string` | `null` | no |
| <a name="input_s3_object_version"></a> [s3\_object\_version](#input\_s3\_object\_version) | n/a | `string` | `null` | no |
| <a name="input_skip_destroy"></a> [skip\_destroy](#input\_skip\_destroy) | n/a | `bool` | `true` | no |
| <a name="input_snap_start"></a> [snap\_start](#input\_snap\_start) | n/a | <pre>object({<br>    apply_on = string<br>  })</pre> | `null` | no |
| <a name="input_source_code_hash"></a> [source\_code\_hash](#input\_source\_code\_hash) | n/a | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `map(string)` | `null` | no |
| <a name="input_timeout"></a> [timeout](#input\_timeout) | n/a | `number` | `1` | no |
| <a name="input_tracing_config"></a> [tracing\_config](#input\_tracing\_config) | n/a | <pre>object({<br>    mode = string<br>  })</pre> | `null` | no |
| <a name="input_vpc_config"></a> [vpc\_config](#input\_vpc\_config) | n/a | <pre>object({<br>    security_group_ids = list(string)<br>    subnet_ids         = list(string)<br>  })</pre> | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | n/a |
| <a name="output_function_name"></a> [function\_name](#output\_function\_name) | n/a |
| <a name="output_invoke_arn"></a> [invoke\_arn](#output\_invoke\_arn) | n/a |
| <a name="output_last_modified"></a> [last\_modified](#output\_last\_modified) | n/a |
| <a name="output_source_code_size"></a> [source\_code\_size](#output\_source\_code\_size) | n/a |
<!-- END_TF_DOCS -->
