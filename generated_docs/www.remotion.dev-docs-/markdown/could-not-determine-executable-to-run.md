---
title: Could not determine executable to run
source: Remotion Documentation
last_updated: 2024-12-22
---

# Could not determine executable to run

- [Home page](/)
- Troubleshooting
- Could not determine executable to run

On this page

# Could not determine executable to run

If you receive the following error:

```

npm ERR! could not determine executable to run
```

```

npm ERR! could not determine executable to run
```

It may be because you are using the `npx` command but the project is configured to use a different package manager.

If [Corepack](https://www.totaltypescript.com/how-to-use-corepack) is enabled, the `packageManager` property in `package.json` matters:

```

package.json
json

{
  "packageManager": "pnpm@7.1.0"
}
```

```

package.json
json

{
  "packageManager": "pnpm@7.1.0"
}
```

If the package manager is set to something else than `npm`, then the `npx` command may give this error message.

## Resolution [​](\#resolution "Direct link to Resolution")

Use the correct command runner for your project:

- If the package manager is set to `pnpm`, use `pnpm exec` instead of `npx`.
- If the package manager is set to `yarn`, use `yarn` instead of `npx`.
- If the package manager is set to `bun`, use `bunx` instead of `npx`.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/could-not-find-executable-to-run.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
SIGKILL](/docs/troubleshooting/sigkill) [Next\
\
Stuck renders](/docs/troubleshooting/stuck-render)

- [Resolution](#resolution)
