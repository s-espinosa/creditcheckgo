package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  number    := os.Args[1]
  padded    := pad(number)
  split     := split(padded)
  converted := convert(split)
  reversed  := reverseArray(converted)
  doubled   := doubleEveryOther(reversed)
  summed    := sumGreaterThanTen(doubled)
  totalled  := totalAllDigits(summed)
  valid     := verify(totalled)
  printValidity(number, valid)
}

func pad(input string) string {
  lenRemainder := len(input) % 2
  var output string
  switch lenRemainder {
  case 0: output = input
  case 1: output = "0" + input
  }
  return output
}

func split(input string) []string {
  return strings.Split(input, "")
}

func convert(input []string) [16]int {
  var converted [16]int
  for i, num := range input {
    newNum, _ := strconv.Atoi(num)
    converted[i] = newNum
  }
  return converted
}

func reverseArray(numbers [16]int) [16]int {
  var reversed [16]int
  for i, num := range numbers {
    newPosition := 15 - i
    reversed[newPosition] = num
  }
  return reversed
}

func doubleEveryOther(numbers [16]int) [16]int {
  var doubled [16]int
  for i, num := range numbers {
    remainder := i % 2
    switch remainder {
      case 0: doubled[i] = num
      case 1: doubled[i] = num * 2
    }
  }
  return doubled
}

func sumGreaterThanTen(numbers[16]int) [16]int {
  var summed [16]int
  for i, num := range numbers {
    greaterThanNine := num > 9
    switch greaterThanNine {
      case true: summed[i] = num - 9
      case false: summed[i] = num
    }
  }
  return summed
}

func totalAllDigits(numbers[16]int) int {
  total := 0
  for _, num := range numbers {
    total += num
  }
  return total
}

func verify(sum int) bool {
  if sum % 10 == 0 {
    return true
  } else {
    return false
  }
}

func printValidity(number string, validity bool) {
  if validity {
    fmt.Println(number, "is a valid card number.")
  } else {
    fmt.Println(number, "is not a valid card number.")
  }
}
