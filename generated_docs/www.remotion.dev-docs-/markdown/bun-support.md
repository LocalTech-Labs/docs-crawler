---
title: Bun support
source: Remotion Documentation
last_updated: 2024-12-22
---

# Bun support

- [Home page](/)
- Miscellaneous
- Bun support

On this page

# Bun support

Remotion is excited about [Bun](https://bun.sh), and we mostly support it (from v1.0.3).

## As a package manager [​](\#as-a-package-manager "Direct link to As a package manager")

You can use `bun i` to initialize all of our Remotion templates.

To scaffold a new project with bun, use:

```

bun create video
```

```

bun create video
```

This command sets all scripts to use [`bunx remotionb`](/docs/cli/#using-bun) which will use Bun as a runtime.

Change `remotionb` to `remotion` if you want to use Node.js as a runtime.

## Remotion CLI [​](\#remotion-cli "Direct link to Remotion CLI")

If you want to run the Remotion CLI using Bun, **use `remotionb` instead of the `remotion` command**.

It doesn't matter if you prefix `remotionb` with `npx`, `bunx` or another runner command.

```

npx remotionb render
```

```

npx remotionb render
```

## As a runtime [​](\#as-a-runtime "Direct link to As a runtime")

As of Bun 1.0.24 and Remotion 4.0.88, the following issues are known:

- ⚠️ The `lazyComponent` prop on `<Composition>` and `<Player>` does not work, and this feature is automatically disabled.
- ⚠️ A server-side rendering script may not quit automatically after it is done running.

Feel free to file more issues with Remotion if you find them.

Previous issues listed here have been resolved as of Bun 1.0.24.

## For contributors [​](\#for-contributors "Direct link to For contributors")

Start the example testbed using `bun run start-bun`.

## See also [​](\#see-also "Direct link to See also")

- [Deno support](/docs/deno)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/bun.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
--gl flag](/docs/gl-options) [Next\
\
Deno support](/docs/deno)

- [As a package manager](#as-a-package-manager)
- [Remotion CLI](#remotion-cli)
- [As a runtime](#as-a-runtime)
- [For contributors](#for-contributors)
- [See also](#see-also)
