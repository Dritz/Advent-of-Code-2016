package main

import (
  "fmt"
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
  Keys [5][5]rune
  X, Y int

}

func CurrentNumber(keypad Keypad) (num rune){
  num = keypad.Keys[keypad.Y][keypad.X]
  return
}

func GetNumber(keypad *Keypad, x, y int) (num rune) {
  num = keypad.Keys[y][x]
  return
}

func NewKeypad(startx int, starty int) Keypad {
  keypad := Keypad{}
  keypad.Keys = [5][5]rune{}
  keypad.Keys[0] = [5]rune{' ', ' ', '1', ' ', ' '}
  keypad.Keys[1] = [5]rune{' ', '2', '3', '4', ' '}
  keypad.Keys[2] = [5]rune{'5', '6', '7', '8', '9'}
  keypad.Keys[3] = [5]rune{' ', 'A', 'B', 'C', ' '}
  keypad.Keys[4] = [5]rune{' ', ' ', 'D', ' ', ' '}
  keypad.X = startx
  keypad.Y = starty
  return keypad
}

func MoveKeypad(keypad *Keypad, dir rune) {
  switch dir {
  case Left:
    if keypad.X != 0 && GetNumber(keypad, keypad.X - 1, keypad.Y) != ' ' {
      keypad.X -= 1
    }
  case Right:
    if keypad.X != 4 && GetNumber(keypad, keypad.X + 1, keypad.Y) != ' ' {
      keypad.X += 1
    }
  case Up:
    if keypad.Y != 0 && GetNumber(keypad, keypad.X, keypad.Y - 1) != ' ' {
      keypad.Y -= 1
    }
  case Down:
    if keypad.Y != 4 && GetNumber(keypad, keypad.X, keypad.Y + 1) != ' ' {
      keypad.Y += 1
    }
  }
}

func main() {
  keypad := NewKeypad(0, 3)
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

    keycode = keycode + fmt.Sprintf("%c", CurrentNumber(keypad))
  }

  fmt.Println("Keycode: " + keycode)
}


