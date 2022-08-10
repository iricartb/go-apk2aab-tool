package hfiles

import (
   "os"
)

const FILE_EXTENSION_APK = ".apk"
const FILE_EXTENSION_JAR = ".jar"
const FILE_EXTENSION_EXE = ".exe"
const FILE_EXTENSION_ZIP = ".zip"
const FILE_EXTENSION_XML = ".xml"
const FILE_EXTENSION_DEX = ".dex"
const FILE_EXTENSION_AAB = ".aab"

func FileExists(sFilename string) bool {
   oInfo, oError := os.Stat(sFilename)

   if os.IsNotExist(oError) {
      return false
   }

   return !oInfo.IsDir()
}

func FolderExists(sFolder string) bool {
   oInfo, oError := os.Stat(sFolder)

   if os.IsNotExist(oError) {
      return false
   }

   return oInfo.IsDir()
}
