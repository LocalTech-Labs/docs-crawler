---
title: Converting a Remotion project to an app
source: Remotion Documentation
last_updated: 2024-12-22
---

# Converting a Remotion project to an app

- [Home page](/)
- Building apps
- Convert Studio to an app

On this page

# Converting a Remotion project to an app

If you have started from any [Remotion template](/templates) except the Next.js and Remix templates, you currently have an app where you can start the [Remotion Studio](/docs/terminology/studio).

If you would like to now turn your project into an app, follow these steps.

[1](#1)

Scaffold a new project using one of those templates:

- [Next.js (App dir)](/templates/next)
- [Next.js (App dir + TailwindCSS)](/templates/next-tailwind)
- [Next.js (Pages dir)](/templates/next-pages-dir)
- [Remix](/templates/remix)

[2](#2)

In the `remotion/Root.tsx` file of the new project, replace
the sample code with the code from `src/Root.tsx` from your existing project.
Include all other components recursively, until `npx tsc -w` does give
no errors anymore.

[3](#3)

Copy the contents of the `public` folder into the new project.

[4](#4)

Find where the `<Player>` is rendered:

- **In the Next.js Pages Dir template:** Under `pages/index.tsx`.

- **In the Next.js App Dir template:** Under `app/page.tsx`.

- **In the Remix template:** Under `app/routes/_index.tsx`.

Choose the main [composition](/docs/terminology/composition) that you want to render in the Player and include it's `component` and other metadata in the `<Player>`. You may import the component from the `remotion` folder so it is imported in both the Studio and the app.

If you want to remove the duplication of metadata, you can define constants such as `DURATION_IN_FRAMES` and `FPS` in a separate file and import the constants in both the Studio and the app.

If you use [`calculateMetadata()`](/docs/dynamic-metadata#with-the-player), see [here](/docs/dynamic-metadata#with-the-player) how you can use it with the Player.

[5](#5)

If you have any Webpack override defined in `remotion.config.ts`
, look up how to apply it to your framework as well.

## Next steps [​](\#next-steps "Direct link to Next steps")

Good luck with your new app!

- Chat with other builders on [Discord](https://remotion.dev/discord).
- Post your app on [`#showcase`](https://remotion.dev/discord) and tag [`@remotion`](https://x.com/remotion) on x.com if you build something!
- If you build together with others, ensure you have the proper [Remotion license](/docs/license).

## See also [​](\#see-also "Direct link to See also")

- [Installation in an existing project](/docs/brownfield)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio-app.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Installation in existing project](/docs/brownfield) [Next\
\
Turn a <Player> into a Remotion project](/docs/player-into-remotion-project)

- [Next steps](#next-steps)
- [See also](#see-also)
