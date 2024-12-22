---
title: Edit captions
source: Remotion Documentation
last_updated: 2024-12-22
---

# Edit captions

- [Home page](/)
- Recorder
- Editing
- Captions

On this page

# Edit captions

## Correct [​](\#correct "Direct link to Correct")

If the AI has made a mistake in the caption, there are three different ways to correct them:

[1](#1)

Inside the Remotion Studio you can correct it by clicking on the
faulty caption in the video. This opens a text editor containing the captions,
from where you can directly edit them. Once you are done with your corrections,
click on "Done" in the right bottom corner of the editor. This will save the
changes you made into the JSON file containing your captions.

[2](#2)

You can directly edit your JSON files located at `
public/[your-composition-id]
` in your project with an editor of your choice.

[3](#3)

If the AI makes the same mistake multiple times, you can also remap
words programatically inside the `config/autocorrect.ts` file by adding
logic inside the `autocorrectWord` function. For example:

```

tsx

const autocorrectWord = (word: Word): Word => {
  // Replace a single word with another one
  if (word.word === " github") {
    return {
      ...word,
      word: word.word.replace("github", " GitHub"),
    };
  }

  return word;
};
```

```

tsx

const autocorrectWord = (word: Word): Word => {
  // Replace a single word with another one
  if (word.word === " github") {
    return {
      ...word,
      word: word.word.replace("github", " GitHub"),
    };
  }

  return word;
};
```

## Highlight Words [​](\#highlight-words "Direct link to Highlight Words")

To highlight specific words, click on the desired word within the captions. In the opening editor, choose the `monospace` option to apply a monospaced font and default blue color to the selected word.
If you wish to change the highlighting color, modify the `WORD_HIGHLIGHT_COLOR` constant in the `config/themes.ts` file.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/editing/captions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Silence removal](/docs/recorder/editing/silence-removal) [Next\
\
Layout](/docs/recorder/editing/layout)

- [Correct](#correct)
- [Highlight Words](#highlight-words)
