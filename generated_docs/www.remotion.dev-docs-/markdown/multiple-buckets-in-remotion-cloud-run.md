---
title: Multiple buckets in Remotion Cloud Run
source: Remotion Documentation
last_updated: 2024-12-22
---

# Multiple buckets in Remotion Cloud Run

- [Home page](/)
- [Cloud Run](/docs/cloudrun-alpha)
- Multiple buckets

On this page

# Multiple buckets in Remotion Cloud Run

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Remotion Cloud Run supports **1 bucket per region and project** that you use Remotion Cloud Run in. It is not necessary to create multiple buckets because:

- A Cloud Storage bucket is an effectively infinitely scalable storage solution.
- Remotion will perfectly isolate each render so they do not interfere with another.
- You can give sites a unique identifier to separate production and development environments.
- The Remotion Cloud Run service is a binary that does not change with your codebase.

You might intuitively create multiple buckets because you have multiple environments, but it is usually not needed.

## Deleting extraneous buckets [​](\#deleting-extraneous-buckets "Direct link to Deleting extraneous buckets")

If you got an error message

```

You have multiple buckets [remotioncloudrun-1a2b3c4d, remotioncloudrun-howDidThisHappen] in your Cloud Storage region [us-east1] starting with "remotioncloudrun-".
```

```

You have multiple buckets [remotioncloudrun-1a2b3c4d, remotioncloudrun-howDidThisHappen] in your Cloud Storage region [us-east1] starting with "remotioncloudrun-".
```

delete the extraneous buckets in the GCP console to fix the error.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/multiple-buckets.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Instance Count](/docs/cloudrun/instancecount) [Next\
\
Limits](/docs/cloudrun/limits)

- [Deleting extraneous buckets](#deleting-extraneous-buckets)
