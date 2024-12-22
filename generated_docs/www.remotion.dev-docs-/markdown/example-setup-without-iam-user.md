---
title: Example setup without IAM user
source: Remotion Documentation
last_updated: 2024-12-22
---

# Example setup without IAM user

- [Home page](/)
- [Lambda](/docs/lambda)
- [Using without IAM users](/docs/lambda/without-iam/)
- IAM roles example

On this page

# Example setup without IAM user

This is a write up of how to use the [example](https://github.com/alexfernandez803/example-lambda) for the technique described under ["Using Lambda without IAM roles"](/docs/lambda/without-iam).

## Prequisites [​](\#prequisites "Direct link to Prequisites")

- Ensure that your local AWS profile is able to deploy to AWS.

## Setup [​](\#setup "Direct link to Setup")

### 1\. Clone or download the project [​](\#1-clone-or-download-the-project "Direct link to 1. Clone or download the project")

The project can be found at [`reference project`](https://github.com/alexfernandez803/example-lambda).

### 2\. Install dependencies [​](\#2-install-dependencies "Direct link to 2. Install dependencies")

- npm
- yarn
- pnpm

```

bash

npm i
```

```

bash

npm i
```

```

bash

pnpm i
```

```

bash

pnpm i
```

```

bash

yarn install
```

```

bash

yarn install
```

### 3\. Create the CDK Stack [​](\#3-create-the-cdk-stack "Direct link to 3. Create the CDK Stack")

This command will deploy the Lambda function and any other resources in the stack.

```

bash

npx aws-cdk deploy \
  --outputs-file ./cdk-outputs.json
```

```

bash

npx aws-cdk deploy \
  --outputs-file ./cdk-outputs.json
```

The Remotion packages are also bundled into the stack, these ensures that [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) can be executed by the [Lambda function](https://github.com/alexfernandez803/example-lambda/blob/main/src/render-function/index.ts).

```

package.json
json

{
  "dependencies": {
    ...
    "remotion": "^3.3.33",
    "@remotion/lambda": "^3.3.33",
  }
}

```

```

package.json
json

{
  "dependencies": {
    ...
    "remotion": "^3.3.33",
    "@remotion/lambda": "^3.3.33",
  }
}

```

The full dependencies are included in the [reference project](https://github.com/alexfernandez803/example-lambda/blob/main/package.json).

### 4\. After deployment [​](\#4-after-deployment "Direct link to 4. After deployment")

```

bash

npx aws-cdk deploy \
  --outputs-file ./cdk-outputs.json
```

```

bash

npx aws-cdk deploy \
  --outputs-file ./cdk-outputs.json
```

```

Deployment progress
bash

Bundling asset cdk-stack/render-function/Code/Stage...

  cdk.out/bundling-temp-5e88d0b45626d59e8e8ddce3b05a886b0e1b381df6e5bbbea1dc2727080641a8/index.js  6.3mb ⚠️

⚡ Done in 295ms

✨  Synthesis time: 4.29s

cdk-stack: building assets...

[0%] start: Building 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1
[0%] start: Building 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1
[50%] success: Built 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1
[100%] success: Built 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1

cdk-stack: assets built

cdk-stack: deploying... [1/1]
[0%] start: Publishing 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1
[0%] start: Publishing 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1
[50%] success: Published 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1
[100%] success: Published 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1

 ✅  cdk-stack (no changes)

✨  Deployment time: 1.39s
```

```

Deployment progress
bash

Bundling asset cdk-stack/render-function/Code/Stage...

  cdk.out/bundling-temp-5e88d0b45626d59e8e8ddce3b05a886b0e1b381df6e5bbbea1dc2727080641a8/index.js  6.3mb ⚠️

⚡ Done in 295ms

✨  Synthesis time: 4.29s

cdk-stack: building assets...

[0%] start: Building 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1
[0%] start: Building 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1
[50%] success: Built 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1
[100%] success: Built 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1

cdk-stack: assets built

cdk-stack: deploying... [1/1]
[0%] start: Publishing 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1
[0%] start: Publishing 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1
[50%] success: Published 779e5babb0ddf0d17c0faebbe5596b03bcf13785f0b20c4cd0fe0c5e616d5593:XXXXXXXXXX-us-east-1
[100%] success: Published 87d5e793cbd198c73c05023515153b142eb2f559e7150579cd2db53362c19b6e:XXXXXXXXXX-us-east-1

 ✅  cdk-stack (no changes)

✨  Deployment time: 1.39s
```

```

Output
bash

Outputs:
cdk-stack.apiUrl = https://du7jfr.execute-api.us-east-1.amazonaws.com/
cdk-stack.region = us-east-1
cdk-stack.userPoolClientId = 4l5tsda2iu8lqugl73m8hgeb83
cdk-stack.userPoolId = us-east-1_bVwFsBUGO
Stack ARN:
arn:aws:cloudformation:us-east-1:XXXXXXXXXX:stack/cdk-stack/faf43800-9878-11ed-a070-0aacc64c8662

```

```

Output
bash

Outputs:
cdk-stack.apiUrl = https://du7jfr.execute-api.us-east-1.amazonaws.com/
cdk-stack.region = us-east-1
cdk-stack.userPoolClientId = 4l5tsda2iu8lqugl73m8hgeb83
cdk-stack.userPoolId = us-east-1_bVwFsBUGO
Stack ARN:
arn:aws:cloudformation:us-east-1:XXXXXXXXXX:stack/cdk-stack/faf43800-9878-11ed-a070-0aacc64c8662

```

The output contains the API Gateway base URL, region and Cognito client ID and user pool ID, which are used for authentication.

### 5\. Cleanup [​](\#5-cleanup "Direct link to 5. Cleanup")

The following will delete the function, in case it's not needed anymore.

```

bash

npx aws-cdk destroy
```

```

bash

npx aws-cdk destroy
```

## Lambda role [​](\#lambda-role "Direct link to Lambda role")

The CDK creates an IAM role named `remotionLambdaServerlessRole` which needs the Remotion policy [setup](/docs/lambda/without-iam).

## Test your endpoint [​](\#test-your-endpoint "Direct link to Test your endpoint")

The API is secured by Cognito which requires an authorization token.

In order to test, you need to do the steps below, just in case you still don't have frontend.

### 1\. Create a Cognito User [​](\#1-create-a-cognito-user "Direct link to 1. Create a Cognito User")

```

bash

aws cognito-idp sign-up \
  --client-id YOUR_USER_POOL_CLIENT_ID \
  --username "sample@test.com" \
  --password "compLicat3d123"
```

```

bash

aws cognito-idp sign-up \
  --client-id YOUR_USER_POOL_CLIENT_ID \
  --username "sample@test.com" \
  --password "compLicat3d123"
```

### 2\. Confirm the user so they can sign in [​](\#2-confirm-the-user-so-they-can-sign-in "Direct link to 2. Confirm the user so they can sign in")

```

bash

aws cognito-idp admin-confirm-sign-up \
  --user-pool-id YOUR_USER_POOL_ID \
  --username "sample@test.com"
```

```

bash

aws cognito-idp admin-confirm-sign-up \
  --user-pool-id YOUR_USER_POOL_ID \
  --username "sample@test.com"
```

### 3\. Log the user to retrieve an identity JWT token [​](\#3-log-the-user-to-retrieve-an-identity-jwt-token "Direct link to 3. Log the user to retrieve an identity JWT token")

```

bash

aws cognito-idp initiate-auth \
  --auth-flow USER_PASSWORD_AUTH \
  --auth-parameters \
  USERNAME="sample@test.com",PASSWORD="compLicat3d123" \
  --client-id YOUR_USER_POOL_CLIENT_ID

```

```

bash

aws cognito-idp initiate-auth \
  --auth-flow USER_PASSWORD_AUTH \
  --auth-parameters \
  USERNAME="sample@test.com",PASSWORD="compLicat3d123" \
  --client-id YOUR_USER_POOL_CLIENT_ID

```

`YOUR_USER_POOL_CLIENT_ID` and `YOUR_USER_POOL_ID` are part of the CDK output.

```

Output
bash

{
    "ChallengeParameters": {},
    "AuthenticationResult": {
        "AccessToken": "eyJraWQiOiJGcUJ....",
        "ExpiresIn": 3600,
        "TokenType": "Bearer",
        "RefreshToken": "eyJjdHkiOiJKV1QiLCJlbm...",
        "IdToken": "eyJraWQiOiJCcjY3Rk5WdzRpYVVYVlpNdF..."
    }
}
```

```

Output
bash

{
    "ChallengeParameters": {},
    "AuthenticationResult": {
        "AccessToken": "eyJraWQiOiJGcUJ....",
        "ExpiresIn": 3600,
        "TokenType": "Bearer",
        "RefreshToken": "eyJjdHkiOiJKV1QiLCJlbm...",
        "IdToken": "eyJraWQiOiJCcjY3Rk5WdzRpYVVYVlpNdF..."
    }
}
```

The API will give you a verbose response but will only use the `IdToken`.

### 4\. Use the token to invoke a request to the endpoint using curl. [​](\#4-use-the-token-to-invoke-a-request-to-the-endpoint-using-curl "Direct link to 4. Use the token to invoke a request to the endpoint using curl.")

#### Request [​](\#request "Direct link to Request")

```

bash

curl --location --request POST 'https://du7jfr6.execute-api.us-east-1.amazonaws.com/render' \
--header 'Authorization: Bearer eyJraWQiOiJGcUJFV1B1cHhxM0NXRko0RVN2..........'
```

```

bash

curl --location --request POST 'https://du7jfr6.execute-api.us-east-1.amazonaws.com/render' \
--header 'Authorization: Bearer eyJraWQiOiJGcUJFV1B1cHhxM0NXRko0RVN2..........'
```

#### Response [​](\#response "Direct link to Response")

```

bash

{"message":"SUCCESS","bucketName":"remotionlambda-apsoutheast2-5essis84y1","renderId":"1pwhfhh11z"}
```

```

bash

{"message":"SUCCESS","bucketName":"remotionlambda-apsoutheast2-5essis84y1","renderId":"1pwhfhh11z"}
```

That's it! You now have an API that you can use to invoke the rendering of a video.

warning

It is important to note that the Lambda function should not be accessible to unauthenticated users.

The function uses version 2 of the CDK, which is still being actively developed.

## Next Steps [​](\#next-steps "Direct link to Next Steps")

- [Customize the Lambda function](https://github.com/alexfernandez803/example-lambda/blob/main/src/render-function/index.ts) so that the rendered video will be moved to another directory.
- Try assigning the Remotion [role](/docs/lambda/without-iam#1-create-role-policy) via CDK [`code`](https://github.com/alexfernandez803/example-lambda/blob/main/lib/remotion-cdk-starter-stack.ts).
- Add request parameters to the Lambda function as input parameters for [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda).

## See also [​](\#see-also "Direct link to See also")

- [Using Lambda without IAM user](/docs/lambda/without-iam)
- [Permissions](/docs/lambda/permissions)
- Some code is borrowed from [bobbyhadz.com](https://bobbyhadz.com/blog/aws-cdk-api-authorizer)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/without-iam/example.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
With IAM Roles](/docs/lambda/without-iam/) [Next\
\
With EC2 + STS](/docs/lambda/ec2)

- [Prequisites](#prequisites)
- [Setup](#setup)
  - [1\. Clone or download the project](#1-clone-or-download-the-project)
  - [2\. Install dependencies](#2-install-dependencies)
  - [3\. Create the CDK Stack](#3-create-the-cdk-stack)
  - [4\. After deployment](#4-after-deployment)
  - [5\. Cleanup](#5-cleanup)
- [Lambda role](#lambda-role)
- [Test your endpoint](#test-your-endpoint)
  - [1\. Create a Cognito User](#1-create-a-cognito-user)
  - [2\. Confirm the user so they can sign in](#2-confirm-the-user-so-they-can-sign-in)
  - [3\. Log the user to retrieve an identity JWT token](#3-log-the-user-to-retrieve-an-identity-jwt-token)
  - [4\. Use the token to invoke a request to the endpoint using curl.](#4-use-the-token-to-invoke-a-request-to-the-endpoint-using-curl)
- [Next Steps](#next-steps)
- [See also](#see-also)
