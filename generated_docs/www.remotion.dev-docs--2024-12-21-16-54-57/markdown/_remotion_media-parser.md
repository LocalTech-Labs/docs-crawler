Media ParserOn this page@remotion/media-parseravailable from v4.0.190
This is an experimental package that parses video and other media files in order to extract relevant metadata.
warningUnstable API: This package is experimental. We might change the API at any time, until we remove this notice.
Design goals:

Support as many relevant containers as possible - currently: .mp4, .mov, .webm, .mkv, .avi, .ts
Work in the browser, and server runtimes (Node.js, Bun, Edge, etc.)
Enable powerful WebCodecs use cases
Fine-grained querying, only download as much data as necessary
Functional TypeScript API
Be useful when passing unsupported media
Solve problems that are relevant for Remotion users

Installation​
npmyarnpnpmbunnpm i --save-exact @remotion/media-parser@4.0.241Copynpm i --save-exact @remotion/media-parser@4.0.241Copypnpm i @remotion/media-parser@4.0.241Copypnpm i @remotion/media-parser@4.0.241Copybun i @remotion/media-parser@4.0.241Copybun i @remotion/media-parser@4.0.241Copyyarn --exact add @remotion/media-parser@4.0.241Copyyarn --exact add @remotion/media-parser@4.0.241CopyThis assumes you are currently using v4.0.241 of Remotion.Also update remotion and all `@remotion/*` packages to the same version.Remove all ^ character in front of the version numbers of it as it can lead to a version conflict.
Guide​
Getting video metadataSimple examples of extracting video metadataFast and slow operationsEfficently use parseMedia()Extract ID3 tags and EXIF dataGet embedded tags from video filesRuntime supportsWhere you can run Media ParserWebCodecsHow Media Parser integrates with WebCodecs
APIs​
The following APIs are available:
parseMedia()Parse a media file.nodeReaderRead a file from the local file system.fetchReaderRead a file using fetch().webFileReaderRead a file from <input type="file">.
License​
Remotion LicenseImprove this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousUninstall Cloud RunNextGetting video metadataInstallationGuideAPIsLicense