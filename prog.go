// You can edit this code!
// Click here and start typing.
package main

import ("fmt"
        "strings" )

type Call struct {
 Url string
}
 
//func Call(url string) (c Call) 
//{
//  return Call{url}
//}

func (c *Call)Speak(input string) (result string) {
  result = strings.Replace(input, "s", "th", -1) 
  result = strings.Replace(result, "S", "Th", -1) 
  return result
}

func main() {
  var c = Call{"asdf"}
  var lithped = c.Speak("this is cool")
  fmt.Printf(lithped)
}

