---
title: Root component Timeout
source: Remotion Documentation
last_updated: 2024-12-22
---

# Root component Timeout

- [Home page](/)
- Troubleshooting
- Root component Timeout

On this page

# Root component Timeout

If you get an error message:

```

A delayRender() "Loading root component" was called but not cleared after 28000ms
```

```

A delayRender() "Loading root component" was called but not cleared after 28000ms
```

You have specified an [entry point](/docs/terminology/entry-point) that does not call [`registerRoot()`](/docs/register-root).

In most of the templates, the entry point is `src/index.ts`.

- Maybe you have specified the list of compositions ( `src/Root.tsx` in most templates) instead.
- Maybe you have specified the filename of a component that you want to render.

Ensure that you are passing the file that calls [`registerRoot()`](/docs/register-root) as the entry point.

In the CLI, the entry point is passed as an argument to the render command:

```

npx remotion render [entry-point]
```

```

npx remotion render [entry-point]
```

In the [`bundle()`](/docs/bundle) and [`deploySite()`](/docs/lambda/deploysite) Node.JS apis, you pass the entry point via the [`entryPoint`](/docs/terminology/entry-point) property.

## See also [​](\#see-also "Direct link to See also")

- [Timeout](/docs/timeout)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/loading-root-component.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Apple Silicon detected](/docs/troubleshooting/rosetta) [Next\
\
defaultProps too big](/docs/troubleshooting/defaultprops-too-big)

- [See also](#see-also)
