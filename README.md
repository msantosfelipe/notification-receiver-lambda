# NOTIFICATION RECEIVER LAMBDA

This is a new version of the problem described the article [A personalized notification system with Macrodroid, Golang and OneSignal](https://msantosfelipe.medium.com/how-do-i-stay-informed-and-keep-my-data-safe-fe2d20c52247)

This Lambda Golang app consumes events triggered by [Macrodoid](https://www.macrodroid.com/). They are processed and sent to [Pushover app](https://pushover.net/)

## Build & Deploy
#### Github actions is available, the action starts on merge to branch `master`

#### - To deploy with AWS CLI the commands are in Makefile:
```
make lambda
```

- Test it !
```
aws lambda invoke --function-name LambdaNotificationReceiver --payload '{"title":"Tenha o Gemini na sua tela inicial","body":"Instale o app Gemini para acessar todos os recursos pela tela inicial","app":"Google"} output.json --cli-binary-format raw-in-base64-out
cat output.json
```

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

- Create API (copy APP ID in response)
```
aws apigateway create-rest-api \
  --name "NotificationAPI" \
  --description "API for sending notifications" \
  --region us-east-1
```

- Get Resource ID (copy in response)
```
aws apigateway get-resources \
  --rest-api-id <APP_ID> \
  --region us-east-1

```

- Create resource (Copy the new resource ID)
```
aws apigateway create-resource \
  --rest-api-id <APP_ID> \
  --parent-id <RESOURCE_ID> \
  --path-part notifications \
  --region us-east-1
```

- Create POST method
```
aws apigateway put-method \
  --rest-api-id <APP_ID> \
  --resource-id <NEW_RESOURCE_ID> \
  --http-method POST \
  --request-parameters "method.request.header.apikey=true" \
  --region us-east-1 \
  --authorization-type "NONE"
```

- Integrate resource to call Lambda
```
aws apigateway put-integration \
  --rest-api-id <APP_ID> \
  --resource-id <NEW_RESOURCE_ID> \
  --http-method POST \
  --type AWS_PROXY \
  --integration-http-method POST \
  --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:<ACCOUNT_NAME>:function:LambdaNotificationReceiver/invocations \
  --request-parameters "integration.request.header.apikey=method.request.header.apikey" \
  --region us-east-1
```

- Add permission for API Gateway to invoke Lambda:
```
aws lambda add-permission \
  --function-name LambdaNotificationReceiver \
  --statement-id apigateway-invoke \
  --action lambda:InvokeFunction \
  --principal apigateway.amazonaws.com \
  --source-arn "arn:aws:execute-api:us-east-1:<ACCOUNT_NAME>:<APP_ID>/*/POST/notifications" \
  --region us-east-1
```

- Deploy API
```
aws apigateway create-deployment \
  --rest-api-id <APP_ID> \
  --stage-name prod \
  --region us-east-1
```

- Test it:
curl -X POST \
  https://<APP_ID>.execute-api.us-east-1.amazonaws.com/prod/notifications \
  -H "Content-Type: application/json" \
  -d '{"title":"Tenha o Gemini na sua tela inicial","body":"Instale o app Gemini para acessar todos os recursos pela tela inicial","app":"Google"}'
