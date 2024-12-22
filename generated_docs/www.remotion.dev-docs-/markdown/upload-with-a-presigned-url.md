---
title: Upload with a presigned URL
source: Remotion Documentation
last_updated: 2024-12-22
---

# Upload with a presigned URL

- [Home page](/)
- Building apps
- Upload with a presigned URL

On this page

# Upload with a presigned URL

This article provides guidance for webapps wanting to allow users to upload videos and other assets. We recommend to generate a presigned URL server-side that allows a user to directly upload a file into your cloud storage without having to pass the file through your server.

You can set constraints such as maximal file size and file type, apply rate limiting, require authentication, and predefine the storage location.

## Why use presigned URL? [​](\#why-use-presigned-url "Direct link to Why use presigned URL?")

The traditional way of implementing a file upload would be to let the client upload the file onto a server, which then stores the file on disk or forwards the upload to cloud storage. While this approach works, it's not ideal due to several reasons.

- **Reduce load**: If many clients happen to upload big files on the same server, this server can get slow or even break down under the load. With the presign workflow, the server only needs to create presign URLs, which reduces server load than handling file transfers.
- **Reduce spam**: To prevent your users using your upload feature as free hosting space, you can deny them a presigned URL if they step over your allowance.
- **Data safety**: Since a lot of hosting solutions today are ephemeral or serverless, files should not be stored on them. There is no guarantee the files will still exist after a server restart and you might run out of disk space.

## AWS Example [​](\#aws-example "Direct link to AWS Example")

Here is an example for storing user uploads are stored in AWS S3.

### Permissions [​](\#permissions "Direct link to Permissions")

In your bucket on the AWS console, go to Permissions and allow PUT requests via CORS:

```

Cross-origin resource sharing (CORS) policy
json

[
  {
    "AllowedHeaders": ["*"],
    "AllowedMethods": ["PUT"],
    "AllowedOrigins": ["*"],
    "ExposeHeaders": [],
    "MaxAgeSeconds": 3000
  }
]
```

```

Cross-origin resource sharing (CORS) policy
json

[
  {
    "AllowedHeaders": ["*"],
    "AllowedMethods": ["PUT"],
    "AllowedOrigins": ["*"],
    "ExposeHeaders": [],
    "MaxAgeSeconds": 3000
  }
]
```

note

It may prove useful to also allow the `GET` method via CORS so you can fetch the assets after uploading.

Your AWS user policy must at least have the ability to put an object and make it public:

```

User role policy
json

{
  "Sid": "Presign",
  "Effect": "Allow",
  "Action": ["s3:PutObject", "s3:PutObjectAcl"],
  "Resource": ["arn:aws:s3:::{YOUR_BUCKET_NAME}/*"]
}
```

```

User role policy
json

{
  "Sid": "Presign",
  "Effect": "Allow",
  "Action": ["s3:PutObject", "s3:PutObjectAcl"],
  "Resource": ["arn:aws:s3:::{YOUR_BUCKET_NAME}/*"]
}
```

### Presigning URLs [​](\#presigning-urls "Direct link to Presigning URLs")

First, accept a file in your frontend, for example using `<input type="file">`. You should get a `File`, from which you can determine the content type and content length:

```

App.tsx
ts

const contentType = file.type || 'application/octet-stream';
const arrayBuffer = await file.arrayBuffer();
const contentLength = arrayBuffer.byteLength;
```

```

App.tsx
ts

const contentType = file.type || 'application/octet-stream';
const arrayBuffer = await file.arrayBuffer();
const contentLength = arrayBuffer.byteLength;
```

This example uses [`@aws-sdk/s3-request-presigner`](https://github.com/aws/aws-sdk-js-v3/tree/main/packages/s3-request-presigner) and [the AWS SDK imported from `@remotion/lambda`](/docs/lambda/getawsclient). By calling the function below, two URLs are generated:

- `presignedUrl` is a URL to which the file can be uploaded to
- `readUrl` is the URL from which the file can be read from.

```

generate-presigned-url.ts
tsx

import {getSignedUrl} from '@aws-sdk/s3-request-presigner';
import {AwsRegion, getAwsClient} from '@remotion/lambda/client';

export const generatePresignedUrl = async (
  contentType: string,
  contentLength: number,
  expiresIn: number,
  bucketName: string,
  region: AwsRegion,
): Promise<{presignedUrl: string; readUrl: string}> => {
  if (contentLength > 1024 * 1024 * 200) {
    throw new Error(
      `File may not be over 200MB. Yours is ${contentLength} bytes.`,
    );
  }

  const {client, sdk} = getAwsClient({
    region: process.env.REMOTION_AWS_REGION as AwsRegion,
    service: 's3',
  });

  const key = crypto.randomUUID();

  const command = new sdk.PutObjectCommand({
    Bucket: bucketName,
    Key: key,
    ACL: 'public-read',
    ContentLength: contentLength,
    ContentType: contentType,
  });

  const presignedUrl = await getSignedUrl(client, command, {
    expiresIn,
  });

  // The location of the asset after the upload
  const readUrl = `https://${bucketName}.s3.${region}.amazonaws.com/${key}`;

  return {presignedUrl, readUrl};
};
```

```

generate-presigned-url.ts
tsx

import {getSignedUrl} from '@aws-sdk/s3-request-presigner';
import {AwsRegion, getAwsClient} from '@remotion/lambda/client';

export const generatePresignedUrl = async (
  contentType: string,
  contentLength: number,
  expiresIn: number,
  bucketName: string,
  region: AwsRegion,
): Promise<{presignedUrl: string; readUrl: string}> => {
  if (contentLength > 1024 * 1024 * 200) {
    throw new Error(
      `File may not be over 200MB. Yours is ${contentLength} bytes.`,
    );
  }

  const {client, sdk} = getAwsClient({
    region: process.env.REMOTION_AWS_REGION as AwsRegion,
    service: 's3',
  });

  const key = crypto.randomUUID();

  const command = new sdk.PutObjectCommand({
    Bucket: bucketName,
    Key: key,
    ACL: 'public-read',
    ContentLength: contentLength,
    ContentType: contentType,
  });

  const presignedUrl = await getSignedUrl(client, command, {
    expiresIn,
  });

  // The location of the asset after the upload
  const readUrl = `https://${bucketName}.s3.${region}.amazonaws.com/${key}`;

  return {presignedUrl, readUrl};
};
```

Explanation:

- First, the upload request gets checked for constraints. In this example, we reject uploads that are over 200MB. You could add more constraints or add rate-limiting.
- The AWS SDK gets imported using [getAwsClient()](/docs/lambda/getawsclient). If you don't use Remotion Lambda, install the [`@aws-sdk/client-s3`](https://github.com/aws/aws-sdk-js-v3/tree/main/clients/client-s3) package directly.
- A [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) gets used as the filename to avoid name clashes.
- Finally, the presigned URL and output URL get calculated and returned.

### Next.js example code [​](\#nextjs-example-code "Direct link to Next.js example code")

Here is a sample snippet for the Next.js App Router.

The endpoint is available under `api/upload/route.ts`.

```

app/api/upload/route.ts
tsx

import {NextResponse} from 'next/server';
import {getSignedUrl} from '@aws-sdk/s3-request-presigner';
import {AwsRegion, getAwsClient} from '@remotion/lambda/client';

const generatePresignedUrl = async ({
  contentType,
  contentLength,
  expiresIn,
  bucketName,
  region,
}: {
  contentType: string;
  contentLength: number;
  expiresIn: number;
  bucketName: string;
  region: AwsRegion;
}): Promise<{presignedUrl: string; readUrl: string}> => {
  if (contentLength > 1024 * 1024 * 200) {
    throw new Error(
      `File may not be over 200MB. Yours is ${contentLength} bytes.`,
    );
  }

  const {client, sdk} = getAwsClient({
    region: process.env.REMOTION_AWS_REGION as AwsRegion,
    service: 's3',
  });

  const key = crypto.randomUUID();

  const command = new sdk.PutObjectCommand({
    Bucket: bucketName,
    Key: key,
    ACL: 'public-read',
    ContentLength: contentLength,
    ContentType: contentType,
  });

  const presignedUrl = await getSignedUrl(client, command, {
    expiresIn,
  });

  // The location of the asset after the upload
  const readUrl = `https://${bucketName}.s3.${region}.amazonaws.com/${key}`;

  return {presignedUrl, readUrl};
};

export const POST = async (request: Request) => {
  if (!process.env.REMOTION_AWS_BUCKET_NAME) {
    throw new Error('REMOTION_AWS_BUCKET_NAME is not set');
  }

  if (!process.env.REMOTION_AWS_REGION) {
    throw new Error('REMOTION_AWS_REGION is not set');
  }

  const json = await request.json();
  if (!Number.isFinite(json.size)) {
    throw new Error('size is not a number');
  }
  if (typeof json.contentType !== 'string') {
    throw new Error('contentType is not a string');
  }

  const {presignedUrl, readUrl} = await generatePresignedUrl({
    contentType: json.contentType,
    contentLength: json.size,
    expiresIn: 60 * 60 * 24 * 7,
    bucketName: process.env.REMOTION_AWS_BUCKET_NAME as string,
    region: process.env.REMOTION_AWS_REGION as AwsRegion,
  });

  return NextResponse.json({presignedUrl, readUrl});
};
```

```

app/api/upload/route.ts
tsx

import {NextResponse} from 'next/server';
import {getSignedUrl} from '@aws-sdk/s3-request-presigner';
import {AwsRegion, getAwsClient} from '@remotion/lambda/client';

const generatePresignedUrl = async ({
  contentType,
  contentLength,
  expiresIn,
  bucketName,
  region,
}: {
  contentType: string;
  contentLength: number;
  expiresIn: number;
  bucketName: string;
  region: AwsRegion;
}): Promise<{presignedUrl: string; readUrl: string}> => {
  if (contentLength > 1024 * 1024 * 200) {
    throw new Error(
      `File may not be over 200MB. Yours is ${contentLength} bytes.`,
    );
  }

  const {client, sdk} = getAwsClient({
    region: process.env.REMOTION_AWS_REGION as AwsRegion,
    service: 's3',
  });

  const key = crypto.randomUUID();

  const command = new sdk.PutObjectCommand({
    Bucket: bucketName,
    Key: key,
    ACL: 'public-read',
    ContentLength: contentLength,
    ContentType: contentType,
  });

  const presignedUrl = await getSignedUrl(client, command, {
    expiresIn,
  });

  // The location of the asset after the upload
  const readUrl = `https://${bucketName}.s3.${region}.amazonaws.com/${key}`;

  return {presignedUrl, readUrl};
};

export const POST = async (request: Request) => {
  if (!process.env.REMOTION_AWS_BUCKET_NAME) {
    throw new Error('REMOTION_AWS_BUCKET_NAME is not set');
  }

  if (!process.env.REMOTION_AWS_REGION) {
    throw new Error('REMOTION_AWS_REGION is not set');
  }

  const json = await request.json();
  if (!Number.isFinite(json.size)) {
    throw new Error('size is not a number');
  }
  if (typeof json.contentType !== 'string') {
    throw new Error('contentType is not a string');
  }

  const {presignedUrl, readUrl} = await generatePresignedUrl({
    contentType: json.contentType,
    contentLength: json.size,
    expiresIn: 60 * 60 * 24 * 7,
    bucketName: process.env.REMOTION_AWS_BUCKET_NAME as string,
    region: process.env.REMOTION_AWS_REGION as AwsRegion,
  });

  return NextResponse.json({presignedUrl, readUrl});
};
```

This is how you can call it in the frontend:

```

Uploader.tsx
tsx

const presignedResponse = await fetch('/api/upload', {
  method: 'POST',
  body: JSON.stringify({
    size: file.size,
    contentType: file.type,

const file: File
  }),
});

const json = (await presignedResponse.json()) as {
  presignedUrl: string;
  readUrl: string;
};
```

```

Uploader.tsx
tsx

const presignedResponse = await fetch('/api/upload', {
  method: 'POST',
  body: JSON.stringify({
    size: file.size,
    contentType: file.type,

const file: File
  }),
});

const json = (await presignedResponse.json()) as {
  presignedUrl: string;
  readUrl: string;
};
```

note

This example does not implement any rate limiting or authentication.

## Performing the Uploading [​](\#performing-the-uploading "Direct link to Performing the Uploading")

### Using fetch() [​](\#using-fetch "Direct link to Using fetch()")

Send the presigned URL back to the client. Afterwards, you can now perform an upload using the built-in [`fetch()`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) function:

```

upload-with-fetch.ts
ts

await fetch(presignedUrl, {
  method: 'PUT',
  body: arrayBuffer,
  headers: {
    'content-type': contentType,
  },
});
```

```

upload-with-fetch.ts
ts

await fetch(presignedUrl, {
  method: 'PUT',
  body: arrayBuffer,
  headers: {
    'content-type': contentType,
  },
});
```

### Tracking the upload progress [​](\#tracking-the-upload-progress "Direct link to Tracking the upload progress")

As of October 2024, if you need to track the progress of the upload, you need to use [`XMLHTTPRequest`](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest).

```

upload-with-progress.ts
ts

export type UploadProgress = {
  progress: number;
  loadedBytes: number;
  totalBytes: number;
};

export type OnUploadProgress = (options: UploadProgress) => void;

export const uploadWithProgress = ({
  file,
  url,
  onProgress,
}: {
  file: File;
  url: string;
  onProgress: OnUploadProgress;
}): Promise<void> => {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();

    xhr.open('PUT', url);

    xhr.upload.onprogress = function (event) {
      if (event.lengthComputable) {
        onProgress({
          progress: event.loaded / event.total,
          loadedBytes: event.loaded,
          totalBytes: event.total,
        });
      }
    };

    xhr.onload = function () {
      if (xhr.status === 200) {
        resolve();
      } else {
        reject(new Error(`Upload failed with status: ${xhr.status}`));
      }
    };

    xhr.onerror = function () {
      reject(new Error('Network error occurred during upload'));
    };

    xhr.setRequestHeader('content-type', file.type);
    xhr.send(file);
  });
};
```

```

upload-with-progress.ts
ts

export type UploadProgress = {
  progress: number;
  loadedBytes: number;
  totalBytes: number;
};

export type OnUploadProgress = (options: UploadProgress) => void;

export const uploadWithProgress = ({
  file,
  url,
  onProgress,
}: {
  file: File;
  url: string;
  onProgress: OnUploadProgress;
}): Promise<void> => {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();

    xhr.open('PUT', url);

    xhr.upload.onprogress = function (event) {
      if (event.lengthComputable) {
        onProgress({
          progress: event.loaded / event.total,
          loadedBytes: event.loaded,
          totalBytes: event.total,
        });
      }
    };

    xhr.onload = function () {
      if (xhr.status === 200) {
        resolve();
      } else {
        reject(new Error(`Upload failed with status: ${xhr.status}`));
      }
    };

    xhr.onerror = function () {
      reject(new Error('Network error occurred during upload'));
    };

    xhr.setRequestHeader('content-type', file.type);
    xhr.send(file);
  });
};
```

## See also [​](\#see-also "Direct link to See also")

- [Handling user video uploads](/docs/video-uploads)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/presigned-urls.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Handling user video uploads](/docs/video-uploads) [Next\
\
Build a Google Font picker](/docs/font-picker)

- [Why use presigned URL?](#why-use-presigned-url)
- [AWS Example](#aws-example)
  - [Permissions](#permissions)
  - [Presigning URLs](#presigning-urls)
  - [Next.js example code](#nextjs-example-code)
- [Performing the Uploading](#performing-the-uploading)
  - [Using fetch()](#using-fetch)
  - [Tracking the upload progress](#tracking-the-upload-progress)
- [See also](#see-also)
