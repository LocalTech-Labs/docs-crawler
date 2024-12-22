---
title: Uninstall Cloud Run
source: Remotion Documentation
last_updated: 2024-12-22
---

# Uninstall Cloud Run

- [Home page](/)
- [Cloud Run](/docs/cloudrun-alpha)
- Uninstall Cloud Run

On this page

# Uninstall Cloud Run

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Of course we are bummed if you decided not to use Remotion Cloud Run anymore. We are very receptive to feedback if you'd like to share it with us.

If you would like to remove **all Remotion Cloud Run related objects** from your infrastructure, you can follow these steps.

warning

This will remove all videos already rendered, and break your programs that are using Remotion Cloud Run to render videos.

## Delete Cloud Run Services [​](\#delete-cloud-run-services "Direct link to Delete Cloud Run Services")

You can delete all services using the following command. The `yes` flag is already included, if you run this, it will delete all services without confirmation.

```

npx remotion cloudRun services rmall -y
```

```

npx remotion cloudRun services rmall -y
```

## Delete projects [​](\#delete-projects "Direct link to Delete projects")

```

npx remotion cloudRun sites rmall -y
```

```

npx remotion cloudRun sites rmall -y
```

## Delete renders and artifacts [​](\#delete-renders-and-artifacts "Direct link to Delete renders and artifacts")

Delete all Cloud Storage buckets starting with `remotioncloudrun-` from your GCP Project.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/uninstall.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Upgrading Cloud Run](/docs/cloudrun/upgrading) [Next\
\
Overview](/docs/media-parser/)

- [Delete Cloud Run Services](#delete-cloud-run-services)
- [Delete projects](#delete-projects)
- [Delete renders and artifacts](#delete-renders-and-artifacts)
