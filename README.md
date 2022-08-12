<img src="https://linube.com/blog/wp-content/uploads/apk-aab.png" width="100%" />

<h1>APK 2 AAB TOOL</h1>

<h2>Tool that allows to transform an APK file to AAB (Android App Bundle)</h2>

According to Google play policy Requirements: from August 2021, Google play will begin to require new applications to be published using Android App bundle (hereinafter referred to as AAB). This format will replace APK as the standard release format

Under normal circumstances, AAB can be generated directly with the packaging of as to meet the requirements and uploaded to Google play

But there will be such a problem. Not all the time you can get a game project or source code. If you are given an APK package, what do you do?

Next, let’s introduce how to transform APK into AAB step by step

<hr>

<b>NEED TOOLS</b>

bundletool-all-1.6.1.jar

    bundletool. Jar is a tool provided by Google to generate & test AAB, and it is also used in gradle packaging.
    Get by GitHub:github.com/google/bundletool/releases
    Detailed documentation & Usage:developer.android.com/studio/command-line/bundletool

apktool_2.5.0.jar

    Decompile Android APK tools.
    Get by GitHub:github.com/iBotPeaches…

aapt2

    AAPT’s full name is Android asset packaging tool, which is an Android resource packaging tool.
    Obtaining method: Android SDK: $Android_ SDK/build-tools/30.0.3/aapt2
    Detailed documentation & Usage:developer.android.com/studio/command-line/aapt2

android.jar

    Android framework provides system resources and APIs.
    Obtaining method: Android SDK: $Android_ SDK/platforms/android-30/android. jar

<hr>

<b>STEPS</b>

1 Unzip the apk

    Put all tool files in one folder

    Open your command prompt in current directory

    Decompile the apk through apktool.jar

    java -jar apktool_2.5.0.jar d test.apk -s -o decompile_apk -f

2 Compile the resource

    Compile the resources using aapt2

    aapt2.exe compile --dir decompile_apk\res -o res.zip

    After that you will see a res.zip will generate in your current directory

3 Link the resources

    Execute below command in command line

    aapt2.exe link --proto-format -o base.zip -I android.jar --manifest decompile_apk\AndroidManifest.xml --min-sdk-version $version --target-sdk-version $version --version-code $version --version-name $version -R res.zip --auto-add-overlay

    $version should be replace with your apk version for example my apk have min-sdk is 7 , target-sdk is 30 , version-code is 1 , verison-name is 1.0 . So my command will be ->

    aapt2.exe link --proto-format -o base.zip -I android.jar --manifest decompile_apk\AndroidManifest.xml --min-sdk-version 7 --target-sdk-version 30 --version-code 1 --version-name 1.0 -R res.zip --auto-add-overlay

    After that you will see a base.zip will generate in your current directory

4 Unzip the base.zip

    Directory structure:

    base/
    /AndroidManifest.xml
    /res
    /resources.pb

5 Copy the files

    Take base folder as your main folder for now !

    Create a folder manifest name folder inside base folder and move your base/AndroidManifest.xml to manifest/AndroidManifest.xml

    Copy whole assets folder from decompile_apk/assets and paste to base/assets

    Copy lib folder from decompile_apk/lib and paste to base/lib

    Copy all files inside unknown folder from decompile_apk/unknown and paste to base/root

    Copy whole kotlin folder from decompile_apk/kotlin and paste to base/root/kotlin

    Final directory structure

    base/
    /assets
    /dex
    /lib
    /manifest
    /res
    /root
    /resources.pb

6 Create a zip

    Open your command prompt in /base directory

    We are going to create a zip using cmd or you can use any software like Winrar & 7zip

    Execute below command in command line

    jar cMf base.zip manifest dex res root lib assets resources.pb

    It will create base.zip in current directory now copy the base.zip and paste where you put all tool file (.jars , .exe)

7 Compile aab

    Open your command prompt in /tools directory

    Execute below command in command line

    java -jar bundletool.jar build-bundle --modules=base.zip --output=base.aab

    base.aab file will generate in current folder

    You need to sign the generated aab file with your keystore in order to publish the application in play store
