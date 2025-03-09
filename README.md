# NOTIFICATION RECEIVER LAMBDA

## Creating AWS Infra - Run just once:

### - Lambda

- Create role:
```
aws iam create-role \
  --role-name MyLambdaExecutionRole \
  --assume-role-policy-document '{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": { "Service": "lambda.amazonaws.com" },
        "Action": "sts:AssumeRole"
      }
    ]
  }'
```

- Attach role (Grant permission for logging in CloudWatch):
```
aws iam attach-role-policy \
  --role-name MyLambdaExecutionRole \
  --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

- Get role ARN:
```
aws iam get-role --role-name MyLambdaExecutionRole
```

- Create function:
```
aws lambda create-function \
    --function-name LambdaNotificationReceiver \
    --runtime provided.al2023 --handler bootstrap \
    --role arn:aws:iam::<ACCOUNT_NAME>:role/MyLambdaExecutionRole \
    --handler bootstrap \
    --zip-file fileb://function.zip
```

### - API Gateway 


## Build & Deploy
- Do it with make
```
make lambda
```

- Test it !
```
aws lambda invoke --function-name LambdaNotificationReceiver --payload '{"title":"Tenha o Gemini na sua tela inicial","body":"Instale o app Gemini para acessar todos os recursos pela tela inicial","app":"Google"} output.json --cli-binary-format raw-in-base64-out
cat output.json
```
