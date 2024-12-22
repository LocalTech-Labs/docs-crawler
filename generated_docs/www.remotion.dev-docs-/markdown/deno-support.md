---
title: Deno support
source: Remotion Documentation
last_updated: 2024-12-22
---

# Deno support

- [Home page](/)
- Miscellaneous
- Deno support

On this page

# Deno support

As of right now, Remotion **does not support Deno**.

If you would like to experiment and run Remotion using Deno, you can do so.

For this, we created a way to run the Remotion CLI using Deno (available from v2.0.227):

```

npx remotiond
```

```

npx remotiond
```

Use `remotiond` instead of the `remotion` command.

It does not matter if you use `npx` or something else as a task runner.

## Permission flags [​](\#permission-flags "Direct link to Permission flags")

Remotion does not work with Deno is not run with the following permissions:

```

--allow-env --allow-read --allow-write --allow-net --allow-run --allow-sys
```

```

--allow-env --allow-read --allow-write --allow-net --allow-run --allow-sys
```

## See also [​](\#see-also "Direct link to See also")

- [Bun support](/docs/bun)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/deno.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Bun support](/docs/bun) [Next\
\
Standalone functions](/docs/standalone)

- [Permission flags](#permission-flags)
- [See also](#see-also)
