---
title: Customizing Lambda output destination
source: Remotion Documentation
last_updated: 2024-12-22
---

# Customizing Lambda output destination

- [Home page](/)
- [Lambda](/docs/lambda)
- Custom output destination

On this page

# Customizing Lambda output destination

By default a render artifact is saved into the same S3 bucket as where the site is located under the key `renders/${renderId}/out.{extension}` (for example: `renders/hy0k2siao8/out.mp4`)

You can modify the output destination by passing a different filename, writing it into a different bucket or even upload it to a different S3-compatible provider.

## Customizing the output name [​](\#customizing-the-output-name "Direct link to Customizing the output name")

To customize the output filename, pass `outName: "my-filename.mp4"` to [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#outname) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#outname).

On the CLI, use the [`--out-name`](/docs/lambda/cli/render#--out-name) flag.

The output name must match `/^([0-9a-zA-Z-!_.*'()/]+)$/g`.

## Customizing the output bucket [​](\#customizing-the-output-bucket "Direct link to Customizing the output bucket")

To render into a different bucket, specify the `outName` option to [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda) and pass an object with the `key` and `bucketName` values:

```

tsx

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  codec: "h264",
  imageFormat: "jpeg",
  maxRetries: 1,
  privacy: "public",
  outName: {
    key: "my-output",
    bucketName: "output-bucket",
  },
});
```

```

tsx

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  codec: "h264",
  imageFormat: "jpeg",
  maxRetries: 1,
  privacy: "public",
  outName: {
    key: "my-output",
    bucketName: "output-bucket",
  },
});
```

If you like to use this feature:

- You must extend the [default Remotion policy](/docs/lambda/permissions) to allow read and write access to that bucket.
- The bucket must be in the same region.
- When calling APIs such as [`downloadMedia()`](/docs/lambda/downloadmedia) or [`getRenderProgress()`](/docs/lambda/getrenderprogress), you must pass the `bucketName` where the site resides in, not the bucket where the video gets saved.
- The `key` must match `/^([0-9a-zA-Z-!_.*'()/]+)$/g`
- The bucketName must match `/^(?=^.{3,63}$)(?!^(\d+\.)+\d+$)(^(([a-z0-9]|[a-z0-9][a-z0-9-]*[a-z0-9])\.)*([a-z0-9]|[a-z0-9][a-z0-9-]*[a-z0-9])$)/`.

This feature is not supported from the CLI.

## Saving to another cloud [​](\#saving-to-another-cloud "Direct link to Saving to another cloud")

_Available from v3.2.23_

You can upload the file to another S3-compatible provider.

- List of supported providers (non-exhaustive): Cloudflare, DigitalOcean Spaces
- List of unsupported providers (non-exhaustive): Azure Blob Storage (not S3 compatible)

You must pass an `outName` [as specified above](#customizing-the-output-bucket) and also provide an `s3OutputProvider` like in the example below.

```

tsx

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  codec: "h264",
  imageFormat: "jpeg",
  maxRetries: 1,
  privacy: "no-acl",
  outName: {
    key: "my-output",
    bucketName: "output-bucket",
    s3OutputProvider: {
      endpoint: "https://fra1.digitaloceanspaces.com",
      accessKeyId: "<DIGITAL_OCEAN_ACCESS_KEY_ID>",
      secretAccessKey: "<DIGITAL_OCEAN_SECRET_ACCESS_KEY>",
    },
  },
});
```

```

tsx

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  codec: "h264",
  imageFormat: "jpeg",
  maxRetries: 1,
  privacy: "no-acl",
  outName: {
    key: "my-output",
    bucketName: "output-bucket",
    s3OutputProvider: {
      endpoint: "https://fra1.digitaloceanspaces.com",
      accessKeyId: "<DIGITAL_OCEAN_ACCESS_KEY_ID>",
      secretAccessKey: "<DIGITAL_OCEAN_SECRET_ACCESS_KEY>",
    },
  },
});
```

In this example, the output file will be uploaded to DigitalOcean Spaces. The cloud provider will give you the endpoint and credentials.

If you want to use this feature, note the following:

- When calling [`downloadMedia()`](/docs/lambda/downloadmedia#bucketname) or [`getRenderProgress()`](/docs/lambda/getrenderprogress#bucketname), you must pass the AWS `bucketName` where the site resides in, not the bucket name of the foreign cloud.
- When calling [`downloadMedia()`](/docs/lambda/downloadmedia) or [`getRenderProgress()`](/docs/lambda/getrenderprogress), you must provide the `s3OutputProvider` option with the same credentials again.
- By default, Remotion [assumes you use ACL](/docs/lambda/troubleshooting/bucket-disallows-acl) which is less common on other clouds. You need to set `privacy: "no-acl"` if you don't want to use ACL.

This feature is not supported from the CLI.

## Saving to another AWS region [v4.0.112](https://github.com/remotion-dev/remotion/releases/v4.0.112) [​](\#saving-to-another-aws-region "Direct link to saving-to-another-aws-region")

If you plan on saving to another bucket on AWS S3 and would like to use different credentials, you can specify the `region` in the `s3OutputProvider` object.

```

json

{
  "s3OutputProvider": {
    "endpoint": "https://s3.us-west-1.amazonaws.com",
    "accessKeyId": "<DIGITAL_OCEAN_ACCESS_KEY_ID>",
    "secretAccessKey": "<DIGITAL_OCEAN_SECRET_ACCESS_KEY>",
    "region": "us-west-1"
  }
}
```

```

json

{
  "s3OutputProvider": {
    "endpoint": "https://s3.us-west-1.amazonaws.com",
    "accessKeyId": "<DIGITAL_OCEAN_ACCESS_KEY_ID>",
    "secretAccessKey": "<DIGITAL_OCEAN_SECRET_ACCESS_KEY>",
    "region": "us-west-1"
  }
}
```

Note that it is not necessary to provide a custom `s3OutputProvider` if you want to use the same role as you already gave to the Lambda.

You may need to extend your [role policy](/docs/lambda/setup#2-create-role-policy) to allow writing to this bucket.

## See also [​](\#see-also "Direct link to See also")

- Customizing the filename when a file is downloaded using `downloadBehavior`: For [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#downloadbehavior) and [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#downloadbehavior)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/custom-destination.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Custom Layers](/docs/lambda/custom-layers) [Next\
\
Production Checklist](/docs/lambda/checklist)

- [Customizing the output name](#customizing-the-output-name)
- [Customizing the output bucket](#customizing-the-output-bucket)
- [Saving to another cloud](#saving-to-another-cloud)
- [Saving to another AWS region](#saving-to-another-aws-region)
- [See also](#see-also)
