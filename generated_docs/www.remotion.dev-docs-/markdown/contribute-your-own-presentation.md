---
title: Contribute your own presentation
source: Remotion Documentation
last_updated: 2024-12-22
---

# Contribute your own presentation

- [Home page](/)
- Contributing
- Adding new presentations

On this page

# Contribute your own presentation

Make your custom presentation accessible for others in the `@remotion/transitions` package.

## Setup the remotion project [​](\#setup-the-remotion-project "Direct link to Setup the remotion project")

If this is your first contribution, see the [CONTRIBUTING.md](https://github.com/remotion-dev/remotion/blob/main/CONTRIBUTING.md) file for information on how to contribute and instructions to set up the remotion project.

## How to proceed [​](\#how-to-proceed "Direct link to How to proceed")

[1](#1)

Create a custom transition. Loook at the [custom presentation](/docs/transitions/presentations/custom) docs to see how it's done.

[2](#2) Add your presentation to the remotion monorepository under `packages/transitions/src/presentations`.

[3](#3) In the `bundle.ts`, add your presentation to the `presentations array`.

```

tsx

  const presentations = ['slide', 'flip', 'wipe', 'fade', ..., 'yourPresentation'];
```

```

tsx

  const presentations = ['slide', 'flip', 'wipe', 'fade', ..., 'yourPresentation'];
```

[4](#4)

Add your presentation to the `exports` of the `package.json` at `packages/transition/package.json` as well as to the `typesVersions`, so it can be correctly imported in other remotion projects.

```

json

"exports": {
  "./yourPresentation": {
    "types": "./dist/presentations/yourPresentation.d.ts",
    "module": "./dist/presentations/yourPresentation.js",
    "import": "./dist/presentations/yourPresentation.js",
    "require": "./dist/cjs/presentations/yourPresentation.js"
    },
  },
"typesVersions": {
  ">=1.0": {
    "yourPresentation": [
      "dist/presentations/yourPresentation.d.ts"
      ],
    },
  }
```

```

json

"exports": {
  "./yourPresentation": {
    "types": "./dist/presentations/yourPresentation.d.ts",
    "module": "./dist/presentations/yourPresentation.js",
    "import": "./dist/presentations/yourPresentation.js",
    "require": "./dist/cjs/presentations/yourPresentation.js"
    },
  },
"typesVersions": {
  ">=1.0": {
    "yourPresentation": [
      "dist/presentations/yourPresentation.d.ts"
      ],
    },
  }
```

Make sure to `pnpm build` in `remotion/packages/transitions` so your transition gets usable in your remotion repository.

[5](#5)

Write a documentation for your presentation. Have a look at the presentations linked in the [presentation](/docs/transitions/presentations) docs for reference. The documentation should consist of the following parts:

- A `short description` of what your presentation does.
- A `demo` of your presentation. For instructions, have a look at the `demo` paragraph in the [contributing to the documentation](/docs/contributing/docs#demos) page, or have a look at the source code of other presentation documentations (\[presentationType\].mdx files).
- An `example code snippet` showing how to use your presentation . See the [type safe snippets](/docs/contributing/docs#type-safe-snippets) docs for instructions on typesafe code snippets.
- The API of your presentation

     For more help on how to write a documentation, see the [contributing to the documentation](/docs/contributing/docs) page.

[6](#6)

Add your presentation to the table of contents at [docs/transitions/presentations](/docs/transitions/presentations) by creating a `<TOCItem>` containing a link to your documentation, a `<PresentationPreview` displaying your presentation and a one-liner describing what your presentation does.

```

Example TOCItem
tsx

<TOCItem link="/docs/transitions/presentations/yourPresentation">
  <div style={row}>
    <PresentationPreview
      durationRestThreshold={0.001}
      effect={yourPresentation()}
    />
    <div style={{ flex: 1, marginLeft: 10 }}>
      <strong>
        <code>{"yourPresentation()"}</code>
      </strong>
      <div>Insert one-liner describing your presentation</div>
    </div>
  </div>
</TOCItem>
```

```

Example TOCItem
tsx

<TOCItem link="/docs/transitions/presentations/yourPresentation">
  <div style={row}>
    <PresentationPreview
      durationRestThreshold={0.001}
      effect={yourPresentation()}
    />
    <div style={{ flex: 1, marginLeft: 10 }}>
      <strong>
        <code>{"yourPresentation()"}</code>
      </strong>
      <div>Insert one-liner describing your presentation</div>
    </div>
  </div>
</TOCItem>
```

An pull request for reference containing all required steps and filechanges can be found [here](https://github.com/remotion-dev/remotion/pull/3199/files).

## See also [​](\#see-also "Direct link to See also")

- [Implementing a new feature](/docs/contributing/feature)
- [Writing documentation](/docs/contributing/docs)
- [How to take a bounty issue](/docs/contributing/bounty)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/contributing/presentation.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Rust code](/docs/contributing/rust) [Next\
\
Authoring a library](/docs/authoring-packages)

- [Setup the remotion project](#setup-the-remotion-project)
- [How to proceed](#how-to-proceed)
- [See also](#see-also)
