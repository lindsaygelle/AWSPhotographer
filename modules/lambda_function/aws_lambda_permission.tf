resource "aws_lambda_permission" "main" {
  action                 = var.permission_action
  count                  = !(var.permission_action == null) ? 1 : 0
  event_source_token     = var.permission_event_source_token
  function_name          = aws_lambda_alias.main.function_name
  function_url_auth_type = var.function_url_auth_type
  principal              = var.permission_principal
  principal_org_id       = var.permission_principal_org_id
  qualifier              = aws_lambda_alias.main.function_version
  source_account         = var.permission_source_account
  statement_id           = var.permission_statement_id
  statement_id_prefix    = var.permission_statement_id_prefix
}
