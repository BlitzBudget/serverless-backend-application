
Deploy
---------------------

### Creating Environments

`aws cloudformation create-stack --stack-name production-user-authentication-stack --template-body file://prod-user-authentication-stack.json  --region eu-west-1 --profile blitzbudget --capabilities CAPABILITY_NAMED_IAM`
`aws cloudformation update-stack --stack-name production-serverless-api-dashboard-backend-stack --template-body file://prod-serverless-api-backend-stack.json  --region eu-west-1 --profile blitzbudget --capabilities CAPABILITY_NAMED_IAM`

### Update CloudFormation Stack

`aws cloudformation update-stack --stack-name production-user-authentication-stack --template-body file://prod-user-authentication-stack.json  --region eu-west-1 --profile blitzbudget --capabilities CAPABILITY_NAMED_IAM`
`aws cloudformation update-stack --stack-name production-serverless-api-dashboard-backend-stack --template-body file://prod-serverless-api-backend-stack.json  --region eu-west-1 --profile blitzbudget --capabilities CAPABILITY_NAMED_IAM`