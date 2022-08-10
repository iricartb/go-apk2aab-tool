package hcompressions

import (
   "archive/zip"
   "fmt"
   "io"
   "os"
   "path/filepath"
   "strings"
)

const OS_PLATFORM_WINDOWS = "windows"

func ZipCompression(sSource string, sTarget string, bNeedBaseDir bool) error {
   zipfile, oError := os.Create(sTarget)
   if oError != nil {
      return oError
   }
   defer zipfile.Close()

   oArchive := zip.NewWriter(zipfile)
   defer oArchive.Close()

   oInfo, oError := os.Stat(sSource)
   if oError != nil {
      return oError
   }

   var sBaseDir string
   if oInfo.IsDir() {
      sBaseDir = filepath.Base(sSource)
   }

   filepath.Walk(sSource, func(sPath string, oInfo os.FileInfo, oError error) error {
      if oError != nil {
         return oError
      }

      oHeader, oError := zip.FileInfoHeader(oInfo)
      if oError != nil {
         return oError
      }

      if sBaseDir != "" {
         if bNeedBaseDir {
            oHeader.Name = filepath.Join(sBaseDir, strings.TrimPrefix(sPath, sSource))
         } else {
            sPath := strings.TrimPrefix(sPath, sSource)
            if len(sPath) > 0 && (sPath[0] == '/' || sPath[0] == '\\') {
               sPath = sPath[1:]
            }
            if len(sPath) == 0 {
               return nil
            }

            oHeader.Name = sPath
         }
      }

      if oInfo.IsDir() {
         oHeader.Name += "/"
      } else {
         oHeader.Method = zip.Deflate
      }

      oHeader.Name = strings.ReplaceAll(oHeader.Name, "\\", "/")

      oWriter, oError := oArchive.CreateHeader(oHeader)
      if oError != nil {
         return oError
      }

      if oInfo.IsDir() {
         return nil
      }

      oFile, oError := os.Open(sPath)
      if oError != nil {
         return oError
      }
      defer oFile.Close()

      _, oError = io.Copy(oWriter, oFile)
      return oError
   })

   return oError
}

func ZipDecompression(sSource string, sDestination string) error {
   // 1. Open the zip file
   oReader, oError := zip.OpenReader(sSource)
   if oError != nil {
      return oError
   }
   defer oReader.Close()

   // 2. Get the absolute destination path
   sDestination, oError = filepath.Abs(sDestination)
   if oError != nil {
      return oError
   }

   // 3. Iterate over zip files inside the archive and unzip each of them
   for _, oFile := range oReader.File {
      oError := unzipFile(oFile, sDestination)
      if oError != nil {
         return oError
      }
   }

   return nil
}

func unzipFile(oFile *zip.File, sDestination string) error {
   // 4. Check if file paths are not vulnerable to Zip Slip
   sFilePath := filepath.Join(sDestination, oFile.Name)
   if !strings.HasPrefix(sFilePath, filepath.Clean(sDestination)+string(os.PathSeparator)) {
      return fmt.Errorf("invalid file path: %s", sFilePath)
   }

   // 5. Create directory tree
   if oFile.FileInfo().IsDir() {
      if oError := os.MkdirAll(sFilePath, os.ModePerm); oError != nil {
         return oError
      }
      return nil
   }

   if oError := os.MkdirAll(filepath.Dir(sFilePath), os.ModePerm); oError != nil {
      return oError
   }

   // 6. Create a destination file for unzipped content
   oDestinationFile, oError := os.OpenFile(sFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, oFile.Mode())
   if oError != nil {
      return oError
   }
   defer oDestinationFile.Close()

   // 7. Unzip the content of a file and copy it to the destination file
   oZippedFile, oError := oFile.Open()
   if oError != nil {
      return oError
   }
   defer oZippedFile.Close()

   if _, oError := io.Copy(oDestinationFile, oZippedFile); oError != nil {
      return oError
   }

   return nil
}
