lambda: build_lambda zip_lambda update_lambda clear

build_lambda:
	cd cmd/ && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ../build/bootstrap main.go

zip_lambda:
	cd build/ && zip function.zip bootstrap

clear:
	cd build/ && rm bootstrap && rm function.zip

update_lambda:
	aws lambda update-function-code --function-name LambdaNotificationReceiver --zip-file fileb://build/function.zip

local:
	cd app/ && go run .