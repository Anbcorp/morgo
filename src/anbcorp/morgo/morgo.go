package main

import (
    "fmt"
    "math/rand"
    "time"
)

func initScreen() {
    var player int;
    var player_mark string;

    rand.Seed(time.Now().UnixNano())
    for j :=0; j<12;j+=1 {
        if j%4 == 0 {
            fmt.Println("+---+---+---+")
        } else if j%4 == 1 || j%4 == 3 {
            fmt.Println("|   |   |   |")
        } else {
            for i :=0; i<3; i+=1 {
                fmt.Print("| ")
                player = rand.Intn(2)
                if player == 0 {
                    player_mark = "X"
                } else {
                    player_mark = "O"
                }
                fmt.Print(player_mark)
                fmt.Print(" ")
            }
            fmt.Println("|")
        }
    }
    fmt.Println("+---+---+---+")
}

func mainLoop() {
}

func main() {
  initScreen()
  mainLoop()
}
