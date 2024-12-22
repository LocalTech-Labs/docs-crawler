---
title: Can I embed the Remotion Studio?
source: Remotion Documentation
last_updated: 2024-12-22
---

# Can I embed the Remotion Studio?

- [Home page](/)
- FAQ
- Embedding Remotion Studio

On this page

# Can I embed the Remotion Studio?

We currently don't allow embedding the Remotion Studio as a React component into your own app. This is because:

[1](#1)

The desire to embed the Studio is usually tied to customizing the appearance or the behavior of the Studio. It is not currently designed to be customizable.

[2](#2)

The Studio is connected to a backend, it is not a frontend-only app.

Consider the following alternatives:

## Use the Remotion Player [​](\#use-the-remotion-player "Direct link to Use the Remotion Player")

The Remotion Studio uses the [Remotion Player](/player) and builds a UI around it. You may use the Player yourself and create custom UI around it that aligns with your brand.

## Deploy the Remotion Studio [​](\#deploy-the-remotion-studio "Direct link to Deploy the Remotion Studio")

You can deploy the Remotion Studio to a long-running server in the cloud and make it accessible to your team. [Read here how to do so](/docs/studio/deploy-server).

## Wait for Remotion Bundles to include the Studio [​](\#wait-for-remotion-bundles-to-include-the-studio "Direct link to Wait for Remotion Bundles to include the Studio")

If you are creating a [bundle](/docs/bundle), you will get a static folder that includes HTML, CSS and JavaScript files. Right now, if you open this folder in the browser or deploy it as a static site, you cannot preview the video. We plan to in the foreseeable future include the Remotion Studio in the bundle, so you can preview the video in the browser.

## See also [​](\#see-also "Direct link to See also")

- [Deploy the Remotion Studio](/docs/studio/deploy-server)
- [Remotion Player](/docs/player)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/embed-remotion-studio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Rendering on the edge](/docs/miscellaneous/render-on-edge) [Next\
\
Absolute paths](/docs/miscellaneous/absolute-paths)

- [Use the Remotion Player](#use-the-remotion-player)
- [Deploy the Remotion Studio](#deploy-the-remotion-studio)
- [Wait for Remotion Bundles to include the Studio](#wait-for-remotion-bundles-to-include-the-studio)
- [See also](#see-also)
