output "arn" {
  value = aws_lambda_function.main.arn
}

output "function_name" {
  value = aws_lambda_function.main.function_name
}

output "invoke_arn" {
  value = aws_lambda_function.main.invoke_arn
}

output "last_modified" {
  value = aws_lambda_function.main.last_modified
}

output "source_code_size" {
  value = aws_lambda_function.main.source_code_size
}
