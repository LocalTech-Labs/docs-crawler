---
title: npx remotion lambda sites
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda sites

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- sites

On this page

# npx remotion lambda sites

The `npx remotion lambda sites` command allows to create, view and delete Remotion projects in your S3 bucket.

- [`create`](#create)
- [`ls`](#ls)
- [`rm`](#rm)
- [`rmall`](#rmall)

## create [​](\#create "Direct link to create")

```

npx remotion lambda sites create <entry-point>?
```

```

npx remotion lambda sites create <entry-point>?
```

You may pass an [entry point](/docs/terminology/entry-point) as the first argument, otherwise the entry point will be [determined](/docs/terminology/entry-point#which-entry-point-is-being-used).

Bundle and upload a Remotion video to an S3 bucket.

The result will be a URL such as `https://remotionlambda-12345.s3.eu-central-1.amazonaws.com/sites/abcdef/index.html`.

note

If you make changes locally, you need to redeploy the site. Use [`--site-name`](#--site-name) to overwrite an existing site.

You can use this "Serve URL" to render a video on Remotion Lambda using:

- The [`npx remotion lambda render`](/docs/lambda/cli/render) and [`npx remotion lambda still`](/docs/lambda/cli/still) commands
- The [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) and [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda) functions.
- Locally using the [`renderMedia()`](/docs/renderer/render-media) and [`renderStill()`](/docs/renderer/render-still) functions.
- Locally using the [`npx remotion render`](/docs/cli) and [`npx remotion still`](/docs/cli) commands

If you are rendering on Lambda, you can also pass the site Name (in this case `abcdef`) as an abbreviation.

Example output

```
(1/3) [====================] Bundled video 3975ms
(2/3) [====================] Created bucket 457ms
(3/3) [====================] Uploaded to S3 25118ms

Deployed to S3!
Serve URL: https://remotionlambda-12345.s3.eu-central-1.amazonaws.com/sites/abcdef/index.html
Site Name: abcdef

```

### `--region` [​](\#--region "Direct link to --region")

The [AWS region](/docs/lambda/region-selection) to select. Both project and function should be in this region.

### `--site-name` [​](\#--site-name "Direct link to --site-name")

Uploads the project to a specific directory and returns a deterministic URL. If a site already existed under this name, it will be overwritten. Can only contain the following characters: `0-9`, `a-z`, `A-Z`, `-`, `!`, `_`, `.`, `*`, `'`, `(`, `)`

```

npx remotion lambda sites create src/index.ts --site-name=my-project
```

```

npx remotion lambda sites create src/index.ts --site-name=my-project
```

### `--force-bucket-name` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--force-bucket-name "Direct link to --force-bucket-name")

Specify a specific bucket name to be used. [This is not recommended](/docs/lambda/multiple-buckets), instead let Remotion discover the right bucket automatically.

### `--privacy` [v3.3.97](https://github.com/remotion-dev/remotion/releases/v3.3.97) [​](\#--privacy "Direct link to --privacy")

Either `public` (default) or `no-acl` if you are not using ACL. Sites must have a public URL to be able to be rendered on Lambda, since the headless browser opens that URL.

### `--public-dir` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#--public-dir "Direct link to --public-dir")

Define the location of the [`public/ directory`](/docs/terminology/public-dir). If not defined, Remotion will assume the location is the \`public\` folder in your Remotion root.

### `--enable-folder-expiry` [v4.0.32](https://github.com/remotion-dev/remotion/releases/v4.0.32) [​](\#--enable-folder-expiry "Direct link to --enable-folder-expiry")

When deploying sites, enable or disable S3 Lifecycle policies which allow for renders to auto-delete after a certain time. Default is `null`, which does not change any lifecycle policies of the S3 bucket. See: [Lambda autodelete](/docs/lambda/autodelete).

### `--throw-if-site-exists` [v4.0.141](https://github.com/remotion-dev/remotion/releases/v4.0.141) [​](\#--throw-if-site-exists "Direct link to --throw-if-site-exists")

Prevents accidential update of an existing site. If there are any files in the subfolder where the site should be placed, the function will throw.

### `--disable-git-source` [v4.0.182](https://github.com/remotion-dev/remotion/releases/v4.0.182) [​](\#--disable-git-source "Direct link to --disable-git-source")

Disables the Git Source being connected to the Remotion Studio. Clicking on stack traces and certain menu items will be disabled.

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style "Direct link to --force-path-style")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

* * *

## ls [​](\#ls "Direct link to ls")

```

npx remotion lambda sites ls
```

```

npx remotion lambda sites ls
```

Get a list of sites. The URL that is printed can be passed to the `render` command to render a video.

Example output

```
Site Name Bucket Size Last updated
pr6fwglz05 remotionlambda-abcdefg 14.7 MB 2021-12-02
https://remotionlambda-abcdefg.s3.eu-central-1.amazonaws.com/sites/pr6fwglz05/index.html

testbed remotionlambda-abcdefg 14.7 MB 2021-12-02

https://remotionlambda-abcdefg.s3.eu-central-1.amazonaws.com/sites/testbed/index.html

```

### `--region` [​](\#--region-1 "Direct link to --region-1")

The [AWS region](/docs/lambda/region-selection) to select. Both project and function should be in this region.

### `--quiet`, `-q` [​](\#--quiet--q "Direct link to --quiet--q")

Returns only a list of space-separated sites.

```

npx remotion lambda sites ls -q
```

```

npx remotion lambda sites ls -q
```

Example output

```
pr6fwglz05 testbed

```

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style-1 "Direct link to --force-path-style-1")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

## rm [​](\#rm "Direct link to rm")

Removes a site (or multiple) from S3 by its ID.

```

bash

npx remotion lambda sites rm abcdef
npx remotion lambda sites rm abcdef my-project # multiple at once
```

```

bash

npx remotion lambda sites rm abcdef
npx remotion lambda sites rm abcdef my-project # multiple at once
```

Example output

```
Site abcdef in bucket remotionlambda-gc1w0xbfzl (14.7 MB): Delete? (Y/n): Y
Deleted sites/abcdef/052787b08233d85edebfc4ce4610944e.mp4
Deleted sites/abcdef/258.bundle.js
Deleted sites/abcdef/15.bundle.js
Deleted sites/abcdef/249.bundle.js.map
Deleted sites/abcdef/263.bundle.js
Deleted sites/abcdef/143.bundle.js
Deleted sites/abcdef/258.bundle.js.map
Deleted sites/abcdef/15.bundle.js.map
Deleted sites/abcdef/185.bundle.js.map
Deleted sites/abcdef/249.bundle.js
Deleted sites/abcdef/143.bundle.js.map
Deleted sites/abcdef/185.bundle.js
Deleted sites/abcdef/1f2d09019ff34eed846a5151b8561d5b.mp4
Deleted sites/abcdef/263.bundle.js.map
Deleted sites/abcdef/268.bundle.js
Deleted sites/abcdef/378.bundle.js.map
Deleted sites/abcdef/268.bundle.js.map
Deleted sites/abcdef/378.bundle.js
Deleted sites/abcdef/2b91c5234e41d3c36d4bf6df37876958.webm
Deleted sites/abcdef/450.bundle.js
Deleted sites/abcdef/46.bundle.js.map
Deleted sites/abcdef/46.bundle.js
Deleted sites/abcdef/450.bundle.js.map
Deleted sites/abcdef/534.bundle.js.map
Deleted sites/abcdef/569.bundle.js
Deleted sites/abcdef/3577958454aa99ad707b596f65151746.webm
Deleted sites/abcdef/534.bundle.js
Deleted sites/abcdef/575.bundle.js.map
Deleted sites/abcdef/575.bundle.js
Deleted sites/abcdef/569.bundle.js.map
Deleted sites/abcdef/801.bundle.js
Deleted sites/abcdef/7badbf53d3130d91b90c46181a2ecdc4.webm
Deleted sites/abcdef/801.bundle.js.map
Deleted sites/abcdef/873.bundle.js
Deleted sites/abcdef/98.bundle.js.map
Deleted sites/abcdef/bff822b868a2b87b31877f3606c9cc13.mp3
Deleted sites/abcdef/873.bundle.js.map
Deleted sites/abcdef/98.bundle.js
Deleted sites/abcdef/a2f36e3a48b4989e0da1fea9959fb35f.mp3
Deleted sites/abcdef/bundle.js
Deleted sites/abcdef/bundle.js.map
Deleted sites/abcdef/a7d87d9934059032eebb9c1536378a2a.webm
Deleted sites/abcdef/index.html
Deleted site abcdef and freed up 14.7 MB.

```

### `--region` [​](\#--region-2 "Direct link to --region-2")

The [AWS region](/docs/lambda/region-selection) to select. Both project and function should be in this region.

### `--yes`, `-y` [​](\#--yes--y "Direct link to --yes--y")

Removes a site without asking for confirmation.

```

npx remotion lambda sites rm abcdef -y
```

```

npx remotion lambda sites rm abcdef -y
```

### `--force-bucket-name` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--force-bucket-name-1 "Direct link to --force-bucket-name-1")

Specify a specific bucket name to be used. [This is not recommended](/docs/lambda/multiple-buckets), instead let Remotion discover the right bucket automatically.

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style-2 "Direct link to --force-path-style-2")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

## rmall [​](\#rmall "Direct link to rmall")

Remove all sites in the selected AWS region.

```

bash

npx remotion lambda sites rmall
```

```

bash

npx remotion lambda sites rmall
```

Example output

```
Site abcdef in bucket remotionlambda-gc1w0xbfzl (14.7 MB): Delete? (Y/n): Y
Deleted sites/abcdef/052787b08233d85edebfc4ce4610944e.mp4
Deleted sites/abcdef/258.bundle.js
Deleted sites/abcdef/15.bundle.js
Deleted sites/abcdef/249.bundle.js.map
Deleted sites/abcdef/263.bundle.js
Deleted sites/abcdef/143.bundle.js
Deleted sites/abcdef/258.bundle.js.map
Deleted sites/abcdef/15.bundle.js.map
Deleted sites/abcdef/185.bundle.js.map
Deleted sites/abcdef/249.bundle.js
Deleted sites/abcdef/143.bundle.js.map
Deleted sites/abcdef/185.bundle.js
Deleted sites/abcdef/1f2d09019ff34eed846a5151b8561d5b.mp4
Deleted sites/abcdef/263.bundle.js.map
Deleted sites/abcdef/268.bundle.js
Deleted sites/abcdef/378.bundle.js.map
Deleted sites/abcdef/268.bundle.js.map
Deleted sites/abcdef/378.bundle.js
Deleted sites/abcdef/2b91c5234e41d3c36d4bf6df37876958.webm
Deleted sites/abcdef/450.bundle.js
Deleted sites/abcdef/46.bundle.js.map
Deleted sites/abcdef/46.bundle.js
Deleted sites/abcdef/450.bundle.js.map
Deleted sites/abcdef/534.bundle.js.map
Deleted sites/abcdef/569.bundle.js
Deleted sites/abcdef/3577958454aa99ad707b596f65151746.webm
Deleted sites/abcdef/534.bundle.js
Deleted sites/abcdef/575.bundle.js.map
Deleted sites/abcdef/575.bundle.js
Deleted sites/abcdef/569.bundle.js.map
Deleted sites/abcdef/801.bundle.js
Deleted sites/abcdef/7badbf53d3130d91b90c46181a2ecdc4.webm
Deleted sites/abcdef/801.bundle.js.map
Deleted sites/abcdef/873.bundle.js
Deleted sites/abcdef/98.bundle.js.map
Deleted sites/abcdef/bff822b868a2b87b31877f3606c9cc13.mp3
Deleted sites/abcdef/873.bundle.js.map
Deleted sites/abcdef/98.bundle.js
Deleted sites/abcdef/a2f36e3a48b4989e0da1fea9959fb35f.mp3
Deleted sites/abcdef/bundle.js
Deleted sites/abcdef/bundle.js.map
Deleted sites/abcdef/a7d87d9934059032eebb9c1536378a2a.webm
Deleted sites/abcdef/index.html
Deleted site abcdef and freed up 14.7 MB.

```

### `--region` [​](\#--region-3 "Direct link to --region-3")

The [AWS region](/docs/lambda/region-selection) to select. Both project and function should be in this region.

### `--yes`, `-y` [​](\#--yes--y-1 "Direct link to --yes--y-1")

Removes all sites without asking for confirmation.

```

npx remotion lambda sites rmall -y
```

```

npx remotion lambda sites rmall -y
```

### `--force-bucket-name` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--force-bucket-name-2 "Direct link to --force-bucket-name-2")

Specify a specific bucket name to be used. [This is not recommended](/docs/lambda/multiple-buckets), instead let Remotion discover the right bucket automatically.

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style-3 "Direct link to --force-path-style-3")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/sites.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/lambda/cli) [Next\
\
functions](/docs/lambda/cli/functions)

- [create](#create)
  - [`--region`](#--region)
  - [`--site-name`](#--site-name)
  - [`--force-bucket-name`](#--force-bucket-name)
  - [`--privacy`](#--privacy)
  - [`--public-dir`](#--public-dir)
  - [`--enable-folder-expiry`](#--enable-folder-expiry)
  - [`--throw-if-site-exists`](#--throw-if-site-exists)
  - [`--disable-git-source`](#--disable-git-source)
  - [`--force-path-style`](#--force-path-style)
- [ls](#ls)
  - [`--region`](#--region-1)
  - [`--quiet`, `-q`](#--quiet--q)
  - [`--force-path-style`](#--force-path-style-1)
- [rm](#rm)
  - [`--region`](#--region-2)
  - [`--yes`, `-y`](#--yes--y)
  - [`--force-bucket-name`](#--force-bucket-name-1)
  - [`--force-path-style`](#--force-path-style-2)
- [rmall](#rmall)
  - [`--region`](#--region-3)
  - [`--yes`, `-y`](#--yes--y-1)
  - [`--force-bucket-name`](#--force-bucket-name-2)
  - [`--force-path-style`](#--force-path-style-3)
