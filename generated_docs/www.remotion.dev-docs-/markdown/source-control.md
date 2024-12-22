---
title: Source Control
source: Remotion Documentation
last_updated: 2024-12-22
---

# Source Control

- [Home page](/)
- Recorder
- Source Control

On this page

# Source Control

By default the `public` folder is added to the `.gitignore` file. This means your recordings are not stored in Git.

The videos are being kept locally and only changes to the code can be staged.

We set this as default because GitHub does not allow repositores bigger than 1GB and you may run into an error quickly.

## Enable Git LFS storage [​](\#enable-git-lfs-storage "Direct link to Enable Git LFS storage")

note

We use Git LFS ourselves and recommend setting this up.

If you want to add all your recordings to your repository, you can use Git LFS.

Your Git provider might charge you for storage.

A GitHub account includes 1 GiB of free storage and 1 GiB a month of free bandwidth.

[Afterwards, you can buy a 50GB data pack for $5](https://docs.github.com/en/billing/managing-billing-for-git-large-file-storage/upgrading-git-large-file-storage).

### 1\. Track public folder [​](\#1-track-public-folder "Direct link to 1. Track public folder")

Comment the line in `.gitignore` that ignores the `public` folder:

```

Comment out this line
diff

- public/**/*.{mp4,webm,mov}
+ # public/**/*.{mp4,webm,mov}
```

```

Comment out this line
diff

- public/**/*.{mp4,webm,mov}
+ # public/**/*.{mp4,webm,mov}
```

### 2\. Setup Git LFS [​](\#2-setup-git-lfs "Direct link to 2. Setup Git LFS")

1. [Download Git LFS](https://git-lfs.com/).
2. Initialize Git LFS in your repository:

```

bash

git lfs install
```

```

bash

git lfs install
```

3. Configure Git LFS to track all .mp4 files located in the /public and all its subfolders:

```

bash

git lfs track public/**/*.mp4
```

```

bash

git lfs track public/**/*.mp4
```

### 3\. Commit and push [​](\#3-commit-and-push "Direct link to 3. Commit and push")

```

bash

git add .gitattributes
git add public
git commit -m "Enable Git LFS"
git push
```

```

bash

git add .gitattributes
git add public
git commit -m "Enable Git LFS"
git push
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/source-control.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Subtitles](/docs/recorder/exporting-subtitles) [Next\
\
External recordings](/docs/recorder/external-recordings)

- [Enable Git LFS storage](#enable-git-lfs-storage)
  - [1\. Track public folder](#1-track-public-folder)
  - [2\. Setup Git LFS](#2-setup-git-lfs)
  - [3\. Commit and push](#3-commit-and-push)
