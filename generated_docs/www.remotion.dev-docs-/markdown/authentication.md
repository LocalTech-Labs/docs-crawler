---
title: Authentication
source: Remotion Documentation
last_updated: 2024-12-22
---

# Authentication

- [Home page](/)
- [Lambda](/docs/lambda)
- Authentication

On this page

# Authentication

You can authenticate with the `@remotion/lambda` package either using:

- an `REMOTION_AWS_PROFILE` or `AWS_PROFILE` environment variable pointing to a file
- `REMOTION_AWS_ACCESS_KEY_ID` and `REMOTION_AWS_SECRET_ACCESS_KEY` environment variables
- `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables

Environment variables sitting in a `.env` file are automatically picked up if you use the Remotion CLI, but not if you use the Node.JS APIs. If multiple ways are provided, Remotion will use the order above and use the first credentials found.

We recommend using the environment variable variants prefixed with `REMOTION_` because:

- On some environments, the unprefixed variants may be reserved (e.g. Vercel deployments)
- Confusing conflicts between Remotion and the AWS CLI may be caused if you use the unprefixed versions.

## Rotating credentials [​](\#rotating-credentials "Direct link to Rotating credentials")

Using more than one AWS account can be a viable scaling strategy to increase the [concurrency limit](/docs/lambda/concurrency). To do so, you can set new values for the `REMOTION_AWS_ACCESS_KEY_ID` and `REMOTION_AWS_SECRET_ACCESS_KEY` or other environment variables before making an operation using `@remotion/lambda`. Below is an implementation example.

```

.env
ini

# Account 1
AWS_KEY_1=AK...
AWS_SECRET_=M/

# Account 2
AWS_KEY_2=AK...
AWS_SECRET_2=M/
```

```

.env
ini

# Account 1
AWS_KEY_1=AK...
AWS_SECRET_=M/

# Account 2
AWS_KEY_2=AK...
AWS_SECRET_2=M/
```

note

You need to read the `.env` file yourself using the [`dotenv`](https://npmjs.org/packages/dotenv) package.

```

rotate-credentials.ts
tsx

const getAccountCount = () => {
  let count = 0;
  while (
    process.env["AWS_KEY_" + (count + 1)] &&
    process.env["AWS_SECRET_" + (count + 1)]
  ) {
    count++;
  }

  return count;
};

const getRandomAwsAccount = () => {
  return Math.ceil(Math.random() * getAccountCount());
};

const setEnvForKey = (key: number) => {
  process.env.REMOTION_AWS_ACCESS_KEY_ID = process.env[`AWS_KEY_${key}`];
  process.env.REMOTION_AWS_SECRET_ACCESS_KEY = process.env[`AWS_SECRET_${key}`];
};

// Set random account credentials
setEnvForKey(getRandomAwsAccount());
```

```

rotate-credentials.ts
tsx

const getAccountCount = () => {
  let count = 0;
  while (
    process.env["AWS_KEY_" + (count + 1)] &&
    process.env["AWS_SECRET_" + (count + 1)]
  ) {
    count++;
  }

  return count;
};

const getRandomAwsAccount = () => {
  return Math.ceil(Math.random() * getAccountCount());
};

const setEnvForKey = (key: number) => {
  process.env.REMOTION_AWS_ACCESS_KEY_ID = process.env[`AWS_KEY_${key}`];
  process.env.REMOTION_AWS_SECRET_ACCESS_KEY = process.env[`AWS_SECRET_${key}`];
};

// Set random account credentials
setEnvForKey(getRandomAwsAccount());
```

## Using an AWS profile [​](\#using-an-aws-profile "Direct link to Using an AWS profile")

_available from v3.3.9_

If you prefer [AWS profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html), you may use them. The list of profiles is located at `~/.aws/credentials` on macOS and Linux and has the following format:

```

~/.aws/credentials
ini

[default]

# ...

[remotion]
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
```

```

~/.aws/credentials
ini

[default]

# ...

[remotion]
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
```

In this example, we added a `remotion` profile. Now, by setting `REMOTION_AWS_PROFILE=remotion`, you can select the profile and don't need to pass each environment variable separately anymore.

## Skipping the credentials check [v4.0.160](https://github.com/remotion-dev/remotion/releases/v4.0.160) [​](\#skipping-the-credentials-check "Direct link to skipping-the-credentials-check")

There are other ways the S3 client can authenticate, like from EC2 instance metadata.

If you have set this up, you may set the `REMOTION_SKIP_AWS_CREDENTIALS_CHECK` environment variable to any value.

```

ts

process.env.REMOTION_SKIP_AWS_CREDENTIALS_CHECK = "1";
```

```

ts

process.env.REMOTION_SKIP_AWS_CREDENTIALS_CHECK = "1";
```

Remotion will not check the credentials and will not throw an error if they are not set.

However, if there is a misconfiguration, you may still get an error from the AWS SDK.

Also read the note about [caching clients](#disable-caching).

## Disable caching [v4.0.160](https://github.com/remotion-dev/remotion/releases/v4.0.160) [​](\#disable-caching "Direct link to disable-caching")

AWS clients are cached to save memory and speed up initialization.

The cache key is based on the credentials and the region.

If you opted out of [the credentials check](#skipping-the-credentials-check), the instance is cached for the lifetime of the process.

If you want to disable the cache, set the `REMOTION_SKIP_AWS_CREDENTIALS_CHECK` environment variable to any value.

It's unlikely you need to set this value. This is only if you change the way you authenticate between API calls.

## See also [​](\#see-also "Direct link to See also")

- [Region selection](/docs/lambda/region-selection)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/authentication.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Setup](/docs/lambda/setup) [Next\
\
Permissions](/docs/lambda/permissions)

- [Rotating credentials](#rotating-credentials)
- [Using an AWS profile](#using-an-aws-profile)
- [Skipping the credentials check](#skipping-the-credentials-check)
- [Disable caching](#disable-caching)
- [See also](#see-also)
