---
title: Generate captions
source: Remotion Documentation
last_updated: 2024-12-22
---

# Generate captions

- [Home page](/)
- Recorder
- Generate captions

On this page

# Generate captions

To generate captions, the Remotion Recorder uses [Whisper.cpp](https://github.com/ggerganov/whisper.cpp) for fast and accurate transcriptions.
Each time you record a clip with the Remotion Recorder, captions are automatically generated and persisted to same folder as your recordings.

## Installing Whisper.cpp [​](\#installing-whispercpp "Direct link to Installing Whisper.cpp")

The very first time you finish recording a clip, Whisper.cpp and a 1.5GB model will be installed on your computer. This may take a few minutes.
Once installed, captions for the `webcam` clip will be generated.

note

Captions are only generated for files with the `webcam` prefix.

## Make corrections [​](\#make-corrections "Direct link to Make corrections")

If the AI has made a mistake, no problem, there are various ways to correct the transcriptions manually. [See here how to do this](/docs/recorder/editing/captions).

## Generate captions via CLI [​](\#generate-captions-via-cli "Direct link to Generate captions via CLI")

For [external recordings](/docs/recorder/external-recordings), you can also generate captions via the CLI.

```

bun sub.ts
```

```

bun sub.ts
```

Note that the names of the files you want to transcribe need to start with the prefix `webcam`, all other files will be ingored.
The JSON files containing the captions will be generated and saved under `public/<composition-id>/sub[timestamp].json`.

## Non-english languages [​](\#non-english-languages "Direct link to Non-english languages")

If you do not record in English, edit the `config/whisper.ts` file.

Set the language to a [supported value](/docs/install-whisper-cpp/transcribe#language) change change the `model` to a [supported value](/docs/install-whisper-cpp/download-whisper-model#model) that does not end in `.en`.

It is advised to choose a larger model if you are transcribing in a non-english language.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/captions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Cropping sources](/docs/recorder/record/cropping) [Next\
\
Start editing](/docs/recorder/editing/)

- [Installing Whisper.cpp](#installing-whispercpp)
- [Make corrections](#make-corrections)
- [Generate captions via CLI](#generate-captions-via-cli)
- [Non-english languages](#non-english-languages)
