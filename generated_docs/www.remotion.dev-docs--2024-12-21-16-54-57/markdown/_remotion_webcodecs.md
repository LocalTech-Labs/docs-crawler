WebCodecsOn this page@remotion/webcodecsavailable from v4.0.229
This package provides APIs for converting videos in the browser.
It leverages @remotion/media-parser to parse the video and audio data, and then uses the WebCodecs API to encode the video.

What can you to with this package?​
In browsers that implement WebCodecs, you can use this package to:

Convert videos from one format to another (.mp4 and .webm bidirectional, .avi, .mkv, .mov as import)
Rotate videos
Extract audio from a video
Manipulate the pixels of a video
Fix videos that were recorded with MediaRecorder
Soon: Compress, trim, crop videos

💼 License Disclaimer​
This package is licensed under the Remotion License.We consider a team of 4 or more people a "company".For "companies": A Remotion Company license needs to be obtained to use this package. In a future version of @remotion/webcodecs, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". Get in touch with us if you are planning to use this package.For individuals and teams up to 3: You can use this package for free.This is a short, non-binding explanation of our license. See the License itself for more details.
🚧 Unstable API Warning​
This package is experimental.We might change the API at any time, until we remove this notice.
Installation​
npmyarnpnpmbunnpm i --save-exact @remotion/webcodecs@4.0.241Copynpm i --save-exact @remotion/webcodecs@4.0.241Copypnpm i @remotion/webcodecs@4.0.241Copypnpm i @remotion/webcodecs@4.0.241Copybun i @remotion/webcodecs@4.0.241Copybun i @remotion/webcodecs@4.0.241Copyyarn --exact add @remotion/webcodecs@4.0.241Copyyarn --exact add @remotion/webcodecs@4.0.241CopyThis assumes you are currently using v4.0.241 of Remotion.Also update remotion and all `@remotion/*` packages to the same version.Remove all ^ character in front of the version numbers of it as it can lead to a version conflict.
Guide​
Convert a videofrom one format to anotherRotate a videoFix bad orientationFix a MediaRecorder videoFix missing video duration and poor seeking performance
APIs​
The following APIs are available:
convertMedia()Converts a video using WebCodecs and Media ParsergetAvailableContainers()Get a list of containers @remotion/webcodecs supports.canReencodeVideoTrack()Determine if a video track can be re-encodedcanReencodeAudioTrack()Determine if a audio track can be re-encodedcanCopyVideoTrack()Determine if a video track can be copied without re-encodingcanCopyAudioTrack()Determine if a audio track can be copied without re-encodinggetDefaultAudioCodec()Gets the default audio codec for a container if no other audio codec is specified.getDefaultVideoCodec()Gets the default video codec for a container if no other audio codec is specified.defaultOnAudioTrackHandler()The default track transformation function for audio tracks.defaultOnVideoTrackHandler()The default track transformation function for video tracks.getAvailableAudioCodecs()Get the audio codecs that can fit in a container.getAvailableVideoCodecs()Get the video codecs that can fit in a container.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousForeign file typesNextConvert a videoWhat can you to with this package?💼 License Disclaimer🚧 Unstable API WarningInstallationGuideAPIs