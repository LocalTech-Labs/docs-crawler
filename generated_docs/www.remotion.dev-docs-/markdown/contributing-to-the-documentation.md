---
title: Contributing to the documentation
source: Remotion Documentation
last_updated: 2024-12-22
---

# Contributing to the documentation

- [Home page](/)
- Contributing
- Writing docs

On this page

# Contributing to the documentation

Good documentation makes the difference between frustration and joy! We happily accept contributions to the Remotion documentation.

## Editing an existing page [​](\#editing-an-existing-page "Direct link to Editing an existing page")

At the bottom of each page, the `Improve this page` button is the easiest way to edit a page. You don't even need to setup the repository locally.

## Submitting a new page [​](\#submitting-a-new-page "Direct link to Submitting a new page")

- Basic setup (quicker)
- Full docs

This method will avoid having to compile all docs. You cannot validate if links to other docs are valid.

[1](#1)

Set up the Remotion repository [according to the instructions here](/docs/contributing).

[2](#2) Run `node new-doc.mjs` and follow the instructions.

[1](#1)

Set up the Remotion repository [according the instructions here](/docs/contributing).

[2](#2) Create a new `.md` document in the `packages/docs/docs` folder.

[3](#3) Add the document to `packages/docs/sidebars.js`.

[4](#4) Write what you have to say!

[5](#5) Run `node render-cards.mjs` in `packages/docs` to generate preview cards that will show up if the documentation page is shared on social media.

## Language guidelines [​](\#language-guidelines "Direct link to Language guidelines")

- **Keep it brief**: Developers don't like to read, so adding too much words will lead to information being lost.
- **Link to terminology**: When using a Remotion-specific term, link to the [terminology](/docs/terminology) page that explains it.
- **Avoid emotions** and filler language ("Great! Let's move on to the next step"). Usually it can be removed without losing any information.
- **Separate into multiple paragraphs** if a section is becoming too long.
- **Address the reader as "you"** and not "we".
- **Don't blame the user**: Instead of "You have provided the wrong input", say "The input is invalid" or "the wrong input was provided".
- **Don't assume it is easy**: Avoid using the words "simply" and "just" as beginners might not find it as simple as you do.

## Adding code snippets [​](\#adding-code-snippets "Direct link to Adding code snippets")

You can add codesnippets by beginning a paragraph with three backticks: ```` ``` ````. The code will be highlighted according to the language you specify after the backticks.

```` ```ts```` will highlight the code as TypeScript.

### Type safe snippets [​](\#type-safe-snippets "Direct link to Type safe snippets")

```` ```ts twoslash```` will make a snippet type-safe - it will be checked against the TypeScript compiler. This is the preferred way of writing docs, but if it is too hard, you don't have to do it.

When writing typesafe snippets, sometimes it does not make sense to list all import statements at the top.

You can add a line stating `// ---cut---` and only the content below will be displayed.

### Titles [​](\#titles "Direct link to Titles")

Add a title to a code snippet by adding a line with ```` ```ts twoslash title="file.ts"````:

```

file.ts
ts

console.log("Hello World");
```

```

file.ts
ts

console.log("Hello World");
```

## Special formatting [​](\#special-formatting "Direct link to Special formatting")

### Steps [​](\#steps "Direct link to Steps")

Use `<Step>` to create lists:

```

md

- <Step>1</Step> Step 1
- <Step>2</Step> Step 2
```

```

md

- <Step>1</Step> Step 1
- <Step>2</Step> Step 2
```

- [1](#1) Step 1
- [2](#2) Step 2

### Mark as experimental [​](\#mark-as-experimental "Direct link to Mark as experimental")

Use `<ExperimentalBadge />` to mark something as experimental:

```

md

<ExperimentalBadge>
<p>This feature is still experimental.</p>
</ExperimentalBadge>
```

```

md

<ExperimentalBadge>
<p>This feature is still experimental.</p>
</ExperimentalBadge>
```

EXPERIMENTAL

This feature is still experimental.

### Demos [​](\#demos "Direct link to Demos")

Using `<Demo type="[demo-name]" />` you can render a [Remotion Player](/docs/terminology/player) and specify props that can be updated.

```

md

<Demo type="rect"/>
```

```

md

<Demo type="rect"/>
```

width

`200`

height

`200`

cornerRadius

`0`

edgeRoundness

``

debug

The demo needs to be implemented in `packages/docs/components/demos/index.tsx`.

## See also [​](\#see-also "Direct link to See also")

- [General information](/docs/contributing)
- [Implementing a new feature](/docs/contributing/feature)
- [Implementing a new option](/docs/contributing/option)
- [How to take a bounty issue](/docs/contributing/bounty)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/contributing/docs.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Implementing a new option](/docs/contributing/option) [Next\
\
Formatting](/docs/contributing/formatting)

- [Editing an existing page](#editing-an-existing-page)
- [Submitting a new page](#submitting-a-new-page)
- [Language guidelines](#language-guidelines)
- [Adding code snippets](#adding-code-snippets)
  - [Type safe snippets](#type-safe-snippets)
  - [Titles](#titles)
- [Special formatting](#special-formatting)
  - [Steps](#steps)
  - [Mark as experimental](#mark-as-experimental)
  - [Demos](#demos)
- [See also](#see-also)
