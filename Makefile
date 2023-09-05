.PHONY: apply plan

apply:
	terraform init -upgrade
	terraform apply -input=false

plan:
	terraform init -upgrade
	terraform fmt -check
	terraform plan -input=false
