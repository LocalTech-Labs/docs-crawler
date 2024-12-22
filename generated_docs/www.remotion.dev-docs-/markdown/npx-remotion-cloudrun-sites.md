---
title: npx remotion cloudrun sites
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion cloudrun sites

- [Home page](/)
- [Command line](/docs/cli/)
- [cloudrun](/docs/cloudrun/cli)
- sites

On this page

# npx remotion cloudrun sites

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

The `npx remotion cloudrun sites` command allows to create, view and delete Remotion projects in your Cloud Storage bucket.

- [`create`](#create)
- [`ls`](#ls)
- [`rm`](#rm)
- [`rmall`](#rmall)

## create [​](\#create "Direct link to create")

```

bash

npx remotion cloudrun sites create <entry-point>?
```

```

bash

npx remotion cloudrun sites create <entry-point>?
```

You may pass an [entry point](/docs/terminology/entry-point) as the first argument, otherwise the entry point will be [determined](/docs/terminology/entry-point#which-entry-point-is-being-used).

Bundle and upload a Remotion video to a Cloud Storage bucket.

The result will be a URL such as `https://storage.googleapis.com/remotioncloudrun-12345/sites/mySite123/index.html`.

note

If you make changes locally, you need to redeploy the site. Use [`--site-name`](#--site-name) to overwrite an existing site.

You can use this "Serve URL" to render a video on Remotion Cloud Run using:

- The [`npx remotion cloudrun render`](/docs/cloudrun/cli/render) command.
- Locally using the [`renderMedia()`](/docs/renderer/render-media) and [`renderStill()`](/docs/renderer/render-still) functions.
- Locally using the [`npx remotion render`](/docs/cli) and [`npx remotion still`](/docs/cli) commands

If you are rendering on Cloud Run, you can also pass the site Name (in this case `mySite123`) as an abbreviation.

### `--region` [​](\#--region "Direct link to --region")

The [GCP region](/docs/cloudrun/region-selection) to select. The service accessing the site should also be in this same region to minimise latency.

### `--site-name` [​](\#--site-name "Direct link to --site-name")

Uploads the project to a specific directory and returns a deterministic URL. If a site already existed under this name, in the same region, it will be overwritten. Can only contain the following characters: `0-9`, `a-z`, `A-Z`, `-`, `!`, `_`, `.`, `*`, `'`, `(`, `)`

```

npx remotion cloudrun sites create src/index.ts --site-name=another-site
```

```

npx remotion cloudrun sites create src/index.ts --site-name=another-site
```

### `--disable-git-source` [v4.0.182](https://github.com/remotion-dev/remotion/releases/v4.0.182) [​](\#--disable-git-source "Direct link to --disable-git-source")

Disables the Git Source being connected to the Remotion Studio. Clicking on stack traces and certain menu items will be disabled.

## ls [​](\#ls "Direct link to ls")

```

npx remotion cloudrun sites ls
```

```

npx remotion cloudrun sites ls
```

Get a list of sites. The URL that is printed can be passed to the `render` command to render a video.

Example output

```
2 sites in us-east1, in the remotion-example project.

Site: another-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url: https://storage.googleapis.com/remotioncloudrun-12345/sites/another-site/index.html

Site: test-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url:
https://storage.googleapis.com/remotioncloudrun-12345/sites/test-site/index.html

```

### `--region` [​](\#--region-1 "Direct link to --region-1")

The [GCP region](/docs/cloudrun/region-selection) to list sites from.

### `--all-regions`, [​](\#--all-regions "Direct link to --all-regions")

Ignores region, returning sites across all regions for the project.

```

npx remotion cloudrun sites ls --all-regions
```

```

npx remotion cloudrun sites ls --all-regions
```

Example output

```
3 sites in all regions, in the remotion-example project.

Site: another-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url: https://storage.googleapis.com/remotioncloudrun-12345/sites/another-site/index.html

Site: test-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url:
https://storage.googleapis.com/remotioncloudrun-12345/sites/test-site/index.html

Site: central-site
Bucket: remotioncloudrun-abcdefgh
Region: us-central1
Serve Url: https://storage.googleapis.com/remotioncloudrun-abcdefgh/sites/central-site/index.html
```

### `--quiet`, `-q` [​](\#--quiet--q "Direct link to --quiet--q")

Returns only a list of space-separated sites.

```

npx remotion cloudrun sites ls -q
```

```

npx remotion cloudrun sites ls -q
```

Example output

```
another-site test-site central-site

```

## rm [​](\#rm "Direct link to rm")

Removes a site (or multiple) from Cloud Storage by it's ID.

```

bash

npx remotion cloudrun sites rm central-site
npx remotion cloudrun sites rm central-site another-site # multiple at once
```

```

bash

npx remotion cloudrun sites rm central-site
npx remotion cloudrun sites rm central-site another-site # multiple at once
```

Example output

```
Site: central-site
Bucket: remotioncloudrun-abcdefgh
Region: us-central1
Serve Url:
https://storage.googleapis.com/remotioncloudrun-abcdefgh/sites/central-site/index.html

Delete? (Y/n) Y

Deleted site central-site from bucket remotioncloudrun-abcdefgh.

```

### `--region` [​](\#--region-2 "Direct link to --region-2")

The [GCP region](/docs/cloudrun/region-selection) to remove sites from.

note

The `rm` command does not support the --all-regions flag, as it is possible to have the same site name in multiple regions. This makes it difficult to remove multiple site-names from multiple regions.

### `--yes`, `-y` [​](\#--yes--y "Direct link to --yes--y")

Removes a site (or multiple) without asking for confirmation.

```

npx remotion cloudrun sites rm central-site -y
```

```

npx remotion cloudrun sites rm central-site -y
```

## rmall [​](\#rmall "Direct link to rmall")

Remove all sites in the selected GCP project.

```

bash

npx remotion cloudrun sites rmall
```

```

bash

npx remotion cloudrun sites rmall
```

Example output

```
Retrieving sites in us-east1.

Site: another-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url: https://storage.googleapis.com/remotioncloudrun-12345/sites/another-site/index.html

Delete? (Y/n) n

Skipping site - another-site.

Site: test-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url: https://storage.googleapis.com/remotioncloudrun-12345/sites/test-site/index.html

Delete? (Y/n) n

Skipping site - test-site.
```

### `--region` [​](\#--region-3 "Direct link to --region-3")

The [GCP region](/docs/cloudrun/region-selection) to remove all sites from.

### `--all-regions`, [​](\#--all-regions-1 "Direct link to --all-regions-1")

Ignores region, removing sites across all regions for the project.

```

npx remotion cloudrun sites rmall --all-regions
```

```

npx remotion cloudrun sites rmall --all-regions
```

Example output

```
Retrieving sites in all regions.

Site: another-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url: https://storage.googleapis.com/remotioncloudrun-12345/sites/another-site/index.html

Delete? (Y/n) n

Skipping site - another-site.

Site: test-site
Bucket: remotioncloudrun-12345
Region: us-east1
Serve Url: https://storage.googleapis.com/remotioncloudrun-12345/sites/test-site/index.html

Delete? (Y/n) n

Skipping site - test-site.

Site: central-site
Bucket: remotioncloudrun-abcdefgh
Region: us-central1
Serve Url: https://storage.googleapis.com/remotioncloudrun-abcdefgh/sites/central-site/index.html

Delete? (Y/n) n

Skipping site - central-site.
```

### `--yes`, `-y` [​](\#--yes--y-1 "Direct link to --yes--y-1")

Removes all sites without asking for confirmation.

```

npx remotion cloudrun sites rmall -y
```

```

npx remotion cloudrun sites rmall -y
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/cli/sites.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/cloudrun/cli) [Next\
\
services](/docs/cloudrun/cli/services)

- [create](#create)
  - [`--region`](#--region)
  - [`--site-name`](#--site-name)
  - [`--disable-git-source`](#--disable-git-source)
- [ls](#ls)
  - [`--region`](#--region-1)
  - [`--all-regions`,](#--all-regions)
  - [`--quiet`, `-q`](#--quiet--q)
- [rm](#rm)
  - [`--region`](#--region-2)
  - [`--yes`, `-y`](#--yes--y)
- [rmall](#rmall)
  - [`--region`](#--region-3)
  - [`--all-regions`,](#--all-regions-1)
  - [`--yes`, `-y`](#--yes--y-1)
