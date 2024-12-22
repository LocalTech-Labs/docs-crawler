---
title: Remotion Bundle
source: Remotion Documentation
last_updated: 2024-12-22
---

# Remotion Bundle

- [Home page](/)
- [Terminology](/docs/terminology)
- Remotion Bundle

# Remotion Bundle

Once you have written your video in React, you need to bundle it using [Webpack](https://webpack.js.org/) in order to render it.

The output of the bundling process is called the "bundle".

It is a folder containing HTML, CSS, JavaScript and your assets.

When you render using the command line, a bundle is automatically generated (this is the `Bundling video` step).

You can use the [`bundle()`](/docs/bundle) API to create a bundle programmatically.

You can create a bundle on the CLI using the [`npx remotion bundle`](/docs/cli/bundle) command.

To bundle and deploy to S3, you can use the [`npx remotion lambda sites create`](/docs/lambda/cli/sites#create) command.

If you host a bundle under a URL, it becomes a [Serve URL](/docs/terminology/serve-url), which you can specify to render videos.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/terminology/bundle.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Serve URL](/docs/terminology/serve-url) [Next\
\
Sequence](/docs/terminology/sequence)
