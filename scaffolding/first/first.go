package first

import second "example/second/sub"

func First() string {
   return second.Second()
}
