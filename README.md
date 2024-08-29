# LuhnVerifier
Microservice that will receive a CCN and return if it is valid via the Luhn Algorhithm

# Notes
- Written using JetBrains GoLand IDE: https://www.jetbrains.com/go/
- Targets GoLang 1.23.0
- Written by Kyle Romero, 2024
- Inspired By: https://zerotomastery.io/blog/golang-practice-projects/

# Usage
- Run application
- HTTP Server will expose the following GET Endpoint: http://localhost:4000/Verify
- Body should be JSON in the following format:
  ```
  {
    "CCN" : "5333-6195-0371-5702"
  }
- Endpoint will return JSON response in following format:
  ```
  {
    "valid": true
  }

# Luhn Algorithm
- Details available here: https://en.wikipedia.org/wiki/Luhn_algorithm
- Implementation taken from:
  - https://github.com/durango/go-credit-card/blob/master/creditcard.go
  - The above repo also has advanced functionality and is a good reference for further enhancing this project.
  ```
  function isValid(cardNumber[1..length])
      sum := 0
      parity := length mod 2
      for i from 1 to length do
          if i mod 2 != parity then
              sum := sum + cardNumber[i]
          elseif cardNumber[i] > 4 then
              sum := sum + 2 * cardNumber[i] - 9
          else
              sum := sum + 2 * cardNumber[i]
          end if
      end for
      return cardNumber[length] == (10 - (sum mod 10))
  end function
  
