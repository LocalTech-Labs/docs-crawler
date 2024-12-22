---
title: Formatting in the Remotion repo
source: Remotion Documentation
last_updated: 2024-12-22
---

# Formatting in the Remotion repo

- [Home page](/)
- Contributing
- Formatting

On this page

# Formatting in the Remotion repo

The Remotion codebase uses Prettier and ESLint for formatting and requires all code to be formatted correctly before it is merged.

## Prettier [​](\#prettier "Direct link to Prettier")

In VS Code, you can install the [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode). Extensions for other editors are available on the [Prettier website](https://prettier.io/docs/en/editors.html).

The editor should automatically pick up our `.vscode/settings.json` file and format your code when you save a file.

To manually format a whole package, run `pnpm exec prettier src --write` on a package. Example:

```

bash

cd packages/renderer
pnpm exec prettier src --write
```

```

bash

cd packages/renderer
pnpm exec prettier src --write
```

## ESLint [​](\#eslint "Direct link to ESLint")

ESLint will warn about code style issues and errors. You can install the [ESLint VS Code](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint) extension to get warnings in the editor as you write code. Extensions for other editors are available on the [ESLint website](https://eslint.org/docs/user-guide/integrations).

You can also run `pnpm run lint` in any package to check if there are any errors. Example:

```

bash

cd packages/renderer
pnpm run lint
```

```

bash

cd packages/renderer
pnpm run lint
```

## Testing everything [​](\#testing-everything "Direct link to Testing everything")

You can run

```

bash

pnpm run stylecheck
```

```

bash

pnpm run stylecheck
```

in the root to test locally if the whole repo is well-formatted and will pass the continuous integration checks.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/contributing/formatting.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Writing docs](/docs/contributing/docs) [Next\
\
Bounty issues](/docs/contributing/bounty)

- [Prettier](#prettier)
- [ESLint](#eslint)
- [Testing everything](#testing-everything)
