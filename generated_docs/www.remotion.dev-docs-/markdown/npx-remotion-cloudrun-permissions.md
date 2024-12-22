---
title: npx remotion cloudrun permissions
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion cloudrun permissions

- [Home page](/)
- [Command line](/docs/cli/)
- [cloudrun](/docs/cloudrun/cli)
- permissions

On this page

# npx remotion cloudrun permissions

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Prints and validates the permissions that should be on the IAM role that is attached to the Service Account in GCP, as per the [setup steps](/docs/cloudrun/setup).

```

npx remotion cloudrun permissions
```

```

npx remotion cloudrun permissions
```

Show output

```
✅ iam.serviceAccounts.actAs
✅ run.operations.get
✅ run.routes.invoke
✅ run.services.create
✅ run.services.get
✅ run.services.delete
✅ run.services.list
✅ run.services.update
✅ storage.buckets.create
✅ storage.buckets.get
✅ storage.buckets.list
✅ storage.objects.create
✅ storage.objects.delete
✅ storage.objects.list
✅ logging.logEntries.list

```

## See also [​](\#see-also "Direct link to See also")

- [Setup guide](/docs/cloudrun/setup)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/cli/permissions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
still](/docs/cloudrun/cli/still) [Next\
\
regions](/docs/cloudrun/cli/regions)

- [See also](#see-also)
