# NOTIFICATION RECEIVER LAMBDA

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

- Build project:
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap main.go

```

- Zip function:
```
zip function.zip bootstrap
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

- Test it !
```
aws lambda invoke --function-name LambdaNotificationReceiver --payload '{"name": "Go Developer"}' response.json --cli-binary-format raw-in-base64-out
cat response.json
```

- Update function:
```
aws lambda update-function-code --function-name LambdaNotificationReceiver --zip-file fileb://function.zip
```

