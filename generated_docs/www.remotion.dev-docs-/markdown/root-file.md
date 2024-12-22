---
title: Root file
source: Remotion Documentation
last_updated: 2024-12-22
---

# Root file

- [Home page](/)
- [Terminology](/docs/terminology)
- Root file

# Root file

The Root file is the file that exports the Root component, which renders one or more [`<Composition />`](/docs/composition) components.

By default, it is `src/Root.tsx` or `remotion/Root.tsx` in most Remotion templates, but you may reorganize the files arbitrarily.

The [entry point](/docs/terminology/entry-point) will import the Root file and register it with [`registerRoot()`](/docs/register-root)

See the [Root file of the Hello World project](https://github.com/remotion-dev/template-helloworld/blob/main/src/Root.tsx) as an example.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/terminology/root-file.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Entry point](/docs/terminology/entry-point) [Next\
\
Remotion Root](/docs/terminology/remotion-root)
