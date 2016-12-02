package main

import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "log"
)

// rune constants (enum-like)
const (
  Left rune = 'L'
  Right rune = 'R'
  Up rune = 'U'
  Down rune = 'D'
)

type Keypad struct {
  Keys [3][3]int
  X, Y int

}

func CurrentNumber(keypad Keypad) (num int){
  num = keypad.Keys[keypad.Y][keypad.X]
  return
}

func NewKeypad(startx int, starty int) Keypad {
  keypad := Keypad{}
  keypad.Keys = [3][3]int{}
  keypad.Keys[0] = [3]int{1, 2, 3}
  keypad.Keys[1] = [3]int{4, 5, 6}
  keypad.Keys[2] = [3]int{7, 8, 9}
  keypad.X = startx
  keypad.Y = starty
  return keypad
}

func MoveKeypad(keypad *Keypad, dir rune) {
  switch dir {
  case Left:
    if keypad.X != 0 {
      keypad.X -= 1
    }
  case Right:
    if keypad.X != 2 {
      keypad.X += 1
    }
  case Up:
    if keypad.Y != 0 {
      keypad.Y -= 1
    }
  case Down:
    if keypad.Y != 2 {
      keypad.Y += 1
    }
  }
}

func main() {
  keypad := NewKeypad(1, 1)
  keycode := ""

  file, err := os.Open("./day2input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    input := scanner.Text()

    for i, c := range input {
      _ = i // null function to avoid compiler error...
      MoveKeypad(&keypad, c)
    }

    keycode = keycode + strconv.Itoa(CurrentNumber(keypad))
  }

  fmt.Println("Keycode: " + keycode)
}


