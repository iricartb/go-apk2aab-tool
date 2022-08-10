package hmessages

import (
   "apk2aab/helpers/hcolors"
   "apk2aab/helpers/hstrings"
)

func GetSuccessMessage(sMessage string) string {
   return GetMessage(sMessage, hcolors.Green, "SUCCESS")
}

func GetErrorMessage(sMessage string) string {
   return GetMessage(sMessage, hcolors.Red, "ERROR")
}

func GetInfoMessage(sMessage string) string {
   return GetMessage(sMessage, hcolors.Yellow, "INFO")
}

func GetMessage(sMessage string, sExtraParams ...string) string {
   var sFullMessage string = ""

   if len(sExtraParams) > 0 {
      if len(sExtraParams) > 1 {
         var sTypeMessage string = ""

         if !hstrings.IsEmpty(sExtraParams[0]) {
            sTypeMessage = "[ " + sExtraParams[0] + sExtraParams[1] + hcolors.Reset + " ]"
         } else {
            sTypeMessage = "[ " + sExtraParams[1] + " ]"
         }

         sFullMessage += sTypeMessage + " "
      }
   }

   if len(sExtraParams) > 0 && !hstrings.IsEmpty(sExtraParams[0]) {
      sFullMessage += sExtraParams[0] + sMessage + hcolors.Reset
   } else {
      sFullMessage += sMessage
   }

   return sFullMessage
}
