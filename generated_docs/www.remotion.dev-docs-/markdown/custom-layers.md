---
title: Custom Layers
source: Remotion Documentation
last_updated: 2024-12-22
---

# Custom Layers

- [Home page](/)
- [Lambda](/docs/lambda)
- Custom Layers

On this page

# Custom Layers

The Lambda function [includes Chrome and base fonts by default](/docs/lambda/runtime).

In some advanced use cases, you want to replace certain parts of the stack:

- Use another Chrome version (you may need to [build it yourself](https://github.com/remotion-dev/chrome-build-instructions/tree/main))
- Replace default fonts or emojis

Before you create a custom stack, feel free to contact us to see if Remotion can provide the changes upstream.

Also consider that the AWS Lambda layers cumulatively may not exceed 250MB extracted, so you need to sacrifice an equal amount of other files.

If you just want to add fonts, we recommend to use [Web fonts](/docs/fonts) instead.

## Default layers [v4.0.205](https://github.com/remotion-dev/remotion/releases/v4.0.205) [​](\#default-layers "Direct link to default-layers")

Consider using our layer options before creating a custom stack:

With the default option ( `--runtime-preference=cjk` or no option):

ItemSizechromium188 MBemoji-google9.9 MBfonts1.9 MBcjk16 MB**Total**215.8 MB

With the option [`--runtime-preference=apple-emojis`](/docs/lambda/deployfunction#runtimepreference):

ItemSizechromium188 MBemoji-apple45 MBfonts1.9 MB**Total**235.9 MB

note

Note that by enabling the Apple Emojis, to stay under the 250 MB limit, support for CJK (Chinese, Japanese, Korean) characters will be removed.

Google Emojis are also removed.

## Ensure Remotion version [​](\#ensure-remotion-version "Direct link to Ensure Remotion version")

Customizing Remotion Lambda Layers is possible from v3.0.17.

Lambda binaries might change in minor Remotion versions, it is your responsibility to keep your versions up to date.

## Creating a custom binary [​](\#creating-a-custom-binary "Direct link to Creating a custom binary")

Go to the [`remotion-dev/lambda-binaries`](https://github.com/remotion-dev/lambda-binaries) repository and clone it.

The folders `chromium` and `fonts` contain the binaries for the ARM version. The x64 version has been discontinued.

Put the files that you want in the corresponding folders - for example, add the Apple Emoji Font `AppleColorEmoji.ttf` into the `fonts/.fonts/NotoSans/` folder.

Since the AWS Lambda layers may not exceed 250MB extracted, we need to sacrifice an equal amount of other files - for example by removing `fonts/.fonts/NotoColorEmoji.ttf` and `fonts/.fonts/NotoSansCJKjp-Regular.otf`.

You can see the size of the folders by running:

```

bash

sh size.sh
```

```

bash

sh size.sh
```

If you are done with your changes, run:

```

bash

sh make.sh
```

```

bash

sh make.sh
```

This will zip the layers and put them as artifacts in the `out` directory.

## Creating a Lambda layer [​](\#creating-a-lambda-layer "Direct link to Creating a Lambda layer")

- Go to your AWS console, select the Lambda product, then select "Layers":

![](/img/lambda-layers-console.png)

- Select "Create layer" and fill out the name. Upload the created layer from the `out` folder into the form. The fields "Compatible architectures", "Compatible runtimes" and "License" are optional.

- Once the layer is created, you will get a version ARN (example: `arn:aws:lambda:us-east-1:123456789012:layer:apple-emoji:1`). Copy it.

note

You need to do this for every AWS region you want to use Remotion Lambda in.

## Update the Lambda function [​](\#update-the-lambda-function "Direct link to Update the Lambda function")

Before you continue, make sure a Remotion Lambda function is deployed.

To switch out the layer of a deployed Lambda function, you can either do it via the AWS console or use a Node.JS script that you can run after every function deploy.

- Node.JS
- AWS console

- Go to your AWS console, select the Lambda product, then select the Lambda function you would like to update.
- Click on "Layers":

![](/img/lambda-single-layer.png)

- Click on "Edit", then select the layer you'd like to replace, then click "Remove", then click "Save".
- Once returned, click "Add a layer", then click "Specify an ARN", then paste in the layer version ARN you retrieved from the previous step.
- Click "Verify" and then "Add".

Congrats! You customized your Lambda function.

Before you can update a function via the Node.JS APIs, you need to [add another rule to your user role](/docs/lambda/setup#5-create-an-access-key-for-the-user):

```

json

[
  {
    "Sid": "UpdateFunction",
    "Effect": "Allow",
    "Action": [
      "lambda:GetFunctionConfiguration",
      "lambda:UpdateFunctionConfiguration"
    ],
    "Resource": ["arn:aws:lambda:*:*:function:remotion-render-*"]
  },
  {
    "Sid": "GetOwnLayerVersion",
    "Effect": "Allow",
    "Action": ["lambda:GetLayerVersion"],
    "Resource": ["*"]
  }
]
```

```

json

[
  {
    "Sid": "UpdateFunction",
    "Effect": "Allow",
    "Action": [
      "lambda:GetFunctionConfiguration",
      "lambda:UpdateFunctionConfiguration"
    ],
    "Resource": ["arn:aws:lambda:*:*:function:remotion-render-*"]
  },
  {
    "Sid": "GetOwnLayerVersion",
    "Effect": "Allow",
    "Action": ["lambda:GetLayerVersion"],
    "Resource": ["*"]
  }
]
```

Given a region, function name, layer to remove and layer to add, you can use the following snippet to update a function with custom layers

```

ts

import { AwsRegion, getAwsClient } from "@remotion/lambda";

// Customize these parameters
const REGION: AwsRegion = "us-east-1";
const FUNCTION_NAME = "remotion-render-2022-06-02-mem3000mb-disk2048mb-120sec";
const LAYER_TO_REMOVE = /fonts/;
const LAYER_TO_ADD = "arn:aws:lambda:us-east-1:1234567891:layer:apple-emoji:1";

const { client, sdk } = getAwsClient({
  region: REGION,
  service: "lambda",
});

const fnConfig = await client.send(
  new sdk.GetFunctionConfigurationCommand({
    FunctionName: FUNCTION_NAME,
  }),
);

if (!fnConfig) {
  throw new Error(`Function ${FUNCTION_NAME} not deployed`);
}

await client.send(
  new sdk.UpdateFunctionConfigurationCommand({
    FunctionName: FUNCTION_NAME,
    Layers: [
      ...(fnConfig.Layers ?? [])
        .filter((l) => !l.Arn?.match(LAYER_TO_REMOVE))
        .map((l) => l.Arn as string),
      LAYER_TO_ADD,
    ],
  }),
);
```

```

ts

import { AwsRegion, getAwsClient } from "@remotion/lambda";

// Customize these parameters
const REGION: AwsRegion = "us-east-1";
const FUNCTION_NAME = "remotion-render-2022-06-02-mem3000mb-disk2048mb-120sec";
const LAYER_TO_REMOVE = /fonts/;
const LAYER_TO_ADD = "arn:aws:lambda:us-east-1:1234567891:layer:apple-emoji:1";

const { client, sdk } = getAwsClient({
  region: REGION,
  service: "lambda",
});

const fnConfig = await client.send(
  new sdk.GetFunctionConfigurationCommand({
    FunctionName: FUNCTION_NAME,
  }),
);

if (!fnConfig) {
  throw new Error(`Function ${FUNCTION_NAME} not deployed`);
}

await client.send(
  new sdk.UpdateFunctionConfigurationCommand({
    FunctionName: FUNCTION_NAME,
    Layers: [
      ...(fnConfig.Layers ?? [])
        .filter((l) => !l.Arn?.match(LAYER_TO_REMOVE))
        .map((l) => l.Arn as string),
      LAYER_TO_ADD,
    ],
  }),
);
```

## See also [​](\#see-also "Direct link to See also")

- [Lambda runtime](/docs/lambda/runtime)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/custom-layers.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Light client](/docs/lambda/light-client) [Next\
\
Custom output destination](/docs/lambda/custom-destination)

- [Default layers](#default-layers)
- [Ensure Remotion version](#ensure-remotion-version)
- [Creating a custom binary](#creating-a-custom-binary)
- [Creating a Lambda layer](#creating-a-lambda-layer)
- [Update the Lambda function](#update-the-lambda-function)
- [See also](#see-also)
