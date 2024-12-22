RecorderIntroductionOn this page
The Remotion Recorder is a video production tool built entirely in JavaScript.
Recording​
Record your facecam and screen in sync but independently. This allows you to create a dynamic layout later.
Record webcam and display as separate streams. Up to 4 sources can be recorded at the same time!Ensure you are in the center of the video and you are not too close to the microphone.
Different platforms require different formats​
Each platform has its own conventions for how videos should look.
WIPTwitter, LinkedIn1:1Muted by defaultCaptions burned inYouTube16:9Audio by defaultCaptions as .srt file Endcard with related videosTikTok, Reels, Shorts planned9:16Safe space at bottom requiredWord-by-word captions
The goal of the Recorder: Edit once and feel native on each platform.
This has a measurable impact on the performance of your post: For example, a 1:1 video on X results in 30-35% higher video views and a 80-100% higher engagement rate.
Captions are essential​
Generate captions locally using Whisper.cpp. The Recorder will automatically install it for you.
AI will make mistakes - so we make it easy to correct them.
Captions have word-level timingsOrphan words are avoided by balancing the text layoutUse backticks to highlight technical termsFix captions quickly by clicking on a wordWrite a JavaScript function to fix common misspellingsCaptions are stored as editable JSON
Agnostic endcards​
Your viewer watched the entire video! Now don't make them forget about you.
If you post on Twitter, the call-to-action is "Follow" and your other platforms
are linked.If you post on YouTube, the call to action is "Subscribe" and cutouts for related
videos are made.
The endcard is written with React and CSS, so can be easily customized.
Customization​
Having access to the entire source code, you can use CSS to customize the look of the Recorder.
Look of the built-in light themeLook of the built-in dark themeUse CSS to make it your own!
The Recorder is built with TypeScript and React, and you can fully customize it's behavior.
It's a hackable video editor.
Silence removal​
Automatically remove the silence at the start and end of your take.
Music​
Add background music to your video and automatically fade between the different sections.
Three exclusive tracks are included that you may use for your video.
Layouting​
Whether you only record your facecam, or also your display, the Recorder always finds the ideal layout - even if you mix landscape and portrait videos.
Transitions​
Enable clean transitions between scenes - no matter if you change the layout during the video.
Chapters​
Mark a scene as the beginning of a new chapter and automatically generate animated chapters.
B-Roll​
Add images and videos on top of your facecam or display while your voiceover is playing to give the video more depth.
Exporting​
Export your video with much control over the encoding options.
Render a video with a single click on a button, the CLI, Node.JS APIs, GitHub Actions, or even on AWS Lambda!
Customization​
The Remotion Recorder ships source code.
You can customize it to your liking if you have TypeScript and React knowledge.

1) 9:16 format is not yet implemented.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024NextBefore you buyRecordingDifferent platforms require different formatsCaptions are essentialAgnostic endcardsCustomizationSilence removalMusicLayoutingTransitionsChaptersB-RollExportingCustomization