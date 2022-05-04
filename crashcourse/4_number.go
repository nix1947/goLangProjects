package main
import (
"fmt"
)

func main() {
  fmt.Println("Decimal number is", 42) 
  fmt.Printf("Binary of %d is %b  \n", 42, 42)
  fmt.Printf("Hexadecmal of %d is %x\n", 42, 42)
  fmt.Printf("Octal of %d is %o\n", 42, 42)
  fmt.Printf("UTF-8 of %d is %q\n", 42, 42)

  for i:=0; i<=1000; i++ {
      if i == 0 {
        fmt.Printf("Decimal \t Binary \t HexaDecimal \t Octal \t utf-8 \n")
      }
      fmt.Printf("%d \t %b \t %x \t %o \t %q\n", i, i, i, i, i)
  }

}
