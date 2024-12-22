---
title: Can I render videos using Vercel Serverless functions?
source: Remotion Documentation
last_updated: 2024-12-22
---

# Can I render videos using Vercel Serverless functions?

- [Home page](/)
- FAQ
- Vercel Functions / Next.js

On this page

# Can I render videos using Vercel Serverless functions?

It is currently not possible to render videos or stills on Vercel Serverless functions due to the 50MB maximum function size. Since Chromium is a dependency of Remotion, it alone almost entirely fills the quota available.

However, you can trigger a Remotion Lambda render through a Vercel Serverless function – our preferred choice for building applications. See our [Next template](https://github.com/remotion-dev/template-next) or [GitHub Unwrapped project](https://github.com/remotion-dev/github-unwrapped-2022) for an example of how to do so.

## Can I render videos in Next.js? [​](\#can-i-render-videos-in-nextjs "Direct link to Can I render videos in Next.js?")

If you don't deploy to Vercel, it might be possible to render videos in API routes using the [server-side rendering](/docs/ssr) primitives.

Check if your provider has enough disk space for Chromium and Remotion and has sufficient CPU and RAM to render videos.

⚠️ You need to be aware of the following:

- You cannot use `@remotion/bundler` inside an API route, because it includes Webpack, and the API route itself is being bundled with Webpack. It is not possible to bundle Webpack with Webpack. Instead, use `@remotion/bundler` to create a bundle outside the API route and then use the resulting folder in your API route.
- The `@remotion/renderer` package requires an FFmpeg binary from `node_modules`. Because an API route is being bundled, you might need to override the Webpack configuration to ensure the `remotion` binary is being included in it.
- Error have been reported with the Next.js App Router because the `remotion` package exports client-side components. Those are false positives because no React components are being rendered.

**Recommendation:**

- If you plan to deploy to Vercel, we recommend to trigger Remotion Lambda renders from Vercel Serverless functions. See our [Next template](https://github.com/remotion-dev/template-next) or [GitHub Unwrapped project](https://github.com/remotion-dev/github-unwrapped-2022) for an example of how to do so.
- If you plan to deploy to a [long-running server, we recommend rendering in a pure Node.js environment](/docs/ssr-node) outside of Next.js API routes to avoid the complications of bundling server-side code.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/vercel-functions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Automatic duration](/docs/miscellaneous/automatic-duration) [Next\
\
Rendering on the edge](/docs/miscellaneous/render-on-edge)

- [Can I render videos in Next.js?](#can-i-render-videos-in-nextjs)
