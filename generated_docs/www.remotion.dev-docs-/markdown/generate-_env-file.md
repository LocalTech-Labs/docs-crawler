---
title: Generate .env File
source: Remotion Documentation
last_updated: 2024-12-22
---

# Generate .env File

- [Home page](/)
- [Cloud Run](/docs/cloudrun-alpha)
- Generate .env File

On this page

# Generate .env File

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

As GCP Cloud Shell has access to the project, a script can be run to generate a .env file with the required Service Account keys and Project ID. As Cloud Shell makes use of the logged-in user's credentials for permissions in the project, it is assumed you have the owner role on the project. If you were the one that created the project, you will be an owner.

It is important to have strict control over the service account keys, as they are used to authenticate the Remotion project to GCP. The .env file should not be committed to source control, and not shared with anyone. Any time you no longer require a service account key, it should be deleted.

GCP allows a maximum of 10 keys per service account. If you have already created 10 keys, you will need to delete one before you can create another. The following script provides you with an opportunity to do this from within the terminal.

1. Open the [Project Dashboard](https://console.cloud.google.com/home/dashboard) in the GCP console, and select your existing Remotion project.

2. In the top right hand corner of the screen, click the Activate Cloud Shell icon
![](/img/cloudrun/selectCloudShell.jpg)
3. Within the Cloud Shell, type the following command and follow the prompts.



```

bash

curl -L https://github.com/remotion-dev/remotion/raw/main/packages/cloudrun/src/gcpInstaller/gcpInstaller.tar | tar -x -C . && node install.mjs
```


```

bash

curl -L https://github.com/remotion-dev/remotion/raw/main/packages/cloudrun/src/gcpInstaller/gcpInstaller.tar | tar -x -C . && node install.mjs
```



_The first command downloads a tar file from the Remotion repo, and extracts it to the current directory. The second command runs the installer script._
If you want to generate a new .env file, or manage keys already created, you will want to select option 3.

If this is the first time [initialising Remotion in the GCP project](/docs/cloudrun/setup), you will want to select option 1.

If you are [updating the version of Remotion for this GCP project](/docs/cloudrun/setup), you will want to select option 1.

4. Run the following command to view the environment variables. Copy them into your local `.env` file (create it if it doesn't exist):



```

bash

cat .env
```


```

bash

cat .env
```

5. Remove the .env file from the virtual machine, using this command:



```

bash

rm .env
```


```

bash

rm .env
```

## 5\. Optional: Validate the permission setup [​](\#5-optional-validate-the-permission-setup "Direct link to 5. Optional: Validate the permission setup")

From within your code base, run the following command to validate the permissions are setup correctly in GCP. As long as your GCP project was setup with a matching Remotion version, this should pass.

```

npx remotion cloudrun policies validate
```

```

npx remotion cloudrun policies validate
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/generate-env.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Permissions](/docs/cloudrun/permissions) [Next\
\
Region selection](/docs/cloudrun/region-selection)

- [5\. Optional: Validate the permission setup](#5-optional-validate-the-permission-setup)
