---
title: validateWebhookSignature()
source: Remotion Documentation
last_updated: 2024-12-22
---

# validateWebhookSignature()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- validateWebhookSignature()

On this page

# validateWebhookSignature()

_Available from v3.2.30_

Validates that the signature that was received by a [webhook](/docs/lambda/webhooks) endpoint is authentic. If the validation fails, an error is thrown.

## API [​](\#api "Direct link to API")

The function accepts an object with three key-value pairs:

### `secret` [​](\#secret "Direct link to secret")

The same webhook secret that was passed to [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda)'s webhook options.

### `body` [​](\#body "Direct link to body")

The body that was received by the endpoint - takes a parsed JSON object, not a `string`.

### `signatureHeader` [​](\#signatureheader "Direct link to signatureheader")

The `X-Remotion-Signature` header from the request that was received by the endpoint.

## Example [​](\#example "Direct link to Example")

In the following Next.JS webhook endpoint, an error gets thrown if the signature does not match the one expected one or is missing..

```

pages/api/webhook.ts
tsx

import {
  validateWebhookSignature,
  WebhookPayload,
} from "@remotion/lambda/client";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse,
) {
  validateWebhookSignature({
    secret: process.env.WEBHOOK_SECRET as string,
    body: req.body,
    signatureHeader: req.headers["x-remotion-signature"] as string,
  });

  // If code reaches this path, the webhook is authentic.
  const payload = req.body as WebhookPayload;
  if (payload.type === "success") {
    // ...
  } else if (payload.type === "timeout") {
    // ...
  }

  res.status(200).json({
    success: true,
  });
}
```

```

pages/api/webhook.ts
tsx

import {
  validateWebhookSignature,
  WebhookPayload,
} from "@remotion/lambda/client";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse,
) {
  validateWebhookSignature({
    secret: process.env.WEBHOOK_SECRET as string,
    body: req.body,
    signatureHeader: req.headers["x-remotion-signature"] as string,
  });

  // If code reaches this path, the webhook is authentic.
  const payload = req.body as WebhookPayload;
  if (payload.type === "success") {
    // ...
  } else if (payload.type === "timeout") {
    // ...
  }

  res.status(200).json({
    success: true,
  });
}
```

See [Webhooks](/docs/lambda/webhooks) for an Express example.

## See also [​](\#see-also "Direct link to See also")

- [Webhooks](/docs/lambda/webhooks)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/validate-webhook-signature.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/validatewebhooksignature.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
speculateFunctionName()](/docs/lambda/speculatefunctionname) [Next\
\
@remotion/cloudrun](/docs/cloudrun/api)

- [API](#api)
  - [`secret`](#secret)
  - [`body`](#body)
  - [`signatureHeader`](#signatureheader)
- [Example](#example)
- [See also](#see-also)
