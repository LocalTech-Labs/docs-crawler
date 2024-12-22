---
title: Cloud Run Alpha
source: Remotion Documentation
last_updated: 2024-12-22
---

# Cloud Run Alpha

- [Home page](/)
- Cloud Run

On this page

# Cloud Run Alpha

Help us shape Remotion Cloud Run!

[Cloud Run](/docs/cloudrun) is an alternative option to [Remotion Lambda](/docs/lambda). Where Lambda offers a cloud-based rendering solution on AWS (Amazon Web Services), Cloud Run makes use of GCP (Google Cloud Platform).

## What to test [​](\#what-to-test "Direct link to What to test")

We are looking for feedback on the experience of setting up a GCP Project for Remotion Cloud Run, as well as the required components for rendering on the cloud:

- Deploy a rendering service (in Lambda, a service is known as a function).
- Deploy a Remotion project to GCP Cloud Storage (in Lambda, the storage solution is S3).
- Render a composition stored in Cloud Storage on a Cloud Run service.

We are welcoming any [bug reports](https://remotion.dev/issue).

## 1\. Install `@remotion/cloudrun` [​](\#1-install-remotioncloudrun "Direct link to 1-install-remotioncloudrun")

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/cloudrun@4.0.242Copy
```

```

npm i --save-exact @remotion/cloudrun@4.0.242Copy
```

```

pnpm i @remotion/cloudrun@4.0.242Copy
```

```

pnpm i @remotion/cloudrun@4.0.242Copy
```

```

bun i @remotion/cloudrun@4.0.242Copy
```

```

bun i @remotion/cloudrun@4.0.242Copy
```

```

yarn --exact add @remotion/cloudrun@4.0.242Copy
```

```

yarn --exact add @remotion/cloudrun@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

From `v4.0.18`, Cloud run is distributed together with the main release of Remotion. Before that, you had to install the alpha release (see below).

## Changelog (moved) [​](\#changelog-moved "Direct link to Changelog (moved)")

From `4.0.18` on, see changes [here](https://remotion.dev/changelog).

### `4.0.18` [​](\#4018 "Direct link to 4018")

Remotion Cloud Run is now distributed together with the main release of Remotion. You no longer need to switch to the alpha release, although Remotion Cloud Run is still alpha software. **The changelog is [now part of the main changelog](https://remotion.dev/changelog)**.

### `4.1.0-alpha12` [​](\#410-alpha12 "Direct link to 410-alpha12")

Includes features and bugfixes from `v4.0.17`.
Includes a fix for streaming progress sometimes throwing an exception.

### `4.1.0-alpha11` [​](\#410-alpha11 "Direct link to 410-alpha11")

Includes bugfixes from `v4.0.12`.

### `4.1.0-alpha9` [​](\#410-alpha9 "Direct link to 410-alpha9")

#### Known issues [​](\#known-issues "Direct link to Known issues")

- any internal errors created by Remotion from within the service are not currently sent back in the error response to the renderMediaOnCloudrun and renderStillOnCloudrun APIs (these APIs are also used within the CLI). For these errors, users will need to check the logs for now.

#### Improvements [​](\#improvements "Direct link to Improvements")

- Artifact registry, used to store versioned images for deploying services, now has two folders - production and development.
- Provide helpful response when Cloud Run crashes during render.
  - CLI alerts user there was a crash, fetches logs, determines if cause was likely memory or timeout issue.
  - API can receive a success or crash response
  - New response documented
- Default concurrency for rendering media is now 100%. This will set the concurrency equal to the number of cores the deployed service has.

### `4.1.0-alpha5` [​](\#410-alpha5 "Direct link to 410-alpha5")

- Fix input props not working for dynamic metadata
- Apply changes from `4.0.0-alpha20`.

### `4.1.0-alpha4` [​](\#410-alpha4 "Direct link to 410-alpha4")

Fixed schema error when invoking a render.

Bug fixes leading to public testing.

IssueResolutionRendering a still via CLI with defaults results in error - You can only pass the `quality` option if `imageFormat` is 'jpeg'.Migrated to V4 method, using internalRenderStill() instead of renderStill(). Noticed missing options, added them in and documented.

### `4.1.0-alpha3` [​](\#410-alpha3 "Direct link to 410-alpha3")

Bug fixes leading to public testing.

IssueResolutionWhen deploying a service, the image didn't exist in Google Artifact Registry.Added publish script that runs submit.mjs, automatically deploying the image, tagged with the version number.Functions folder wasn't included in dist folder, so no CLI commandswould work.Removed this from .npmignore, so that it is included.When using the CLI to request a render without passing a composition name, it fails to list out compositions to choose fromIssue raised, present in V4 for Lambda also.Service name structuring clips off alpha version denominator. During alpha, this will make it impossible to deploy multiple services spanning alpha versions.Create new name formatting that meets requirements. Added tests for this.CLI commands for rendering not aligned with Remotion Lambda.`npx remotion cloudrun render media` is now `npx remotion cloudrun render`.

`npx remotion cloudrun render still` is now `npx cloudrun remotion still`.

Documentation also updated.

### `4.1.0-alpha2` [​](\#410-alpha2 "Direct link to 410-alpha2")

Initial cloud run alpha release 🎉.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun-alpha.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Function naming convention](/docs/lambda/naming-convention) [Next\
\
Overview](/docs/cloudrun)

- [What to test](#what-to-test)
- [1\. Install `@remotion/cloudrun`](#1-install-remotioncloudrun)
- [Changelog (moved)](#changelog-moved)
  - [`4.0.18`](#4018)
  - [`4.1.0-alpha12`](#410-alpha12)
  - [`4.1.0-alpha11`](#410-alpha11)
  - [`4.1.0-alpha9`](#410-alpha9)
  - [`4.1.0-alpha5`](#410-alpha5)
  - [`4.1.0-alpha4`](#410-alpha4)
  - [`4.1.0-alpha3`](#410-alpha3)
  - [`4.1.0-alpha2`](#410-alpha2)
