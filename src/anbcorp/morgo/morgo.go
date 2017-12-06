package main

import (
    "fmt"
    "math/rand"
    "time"
)

var board [9]int

func randomBoard() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i<9; i+=1 {
        board[i] = rand.Intn(2)
    }
}

func initScreen() {
    var player int;
    var player_mark string;

    randomBoard()
    fmt.Println(board)

    for j :=0; j<12;j+=1 {
        if j%4 == 0 {
            fmt.Println("+---+---+---+")
        } else if j%4 == 1 || j%4 == 3 {
            fmt.Println("|   |   |   |")
        } else {
            for i :=0; i<3; i+=1 {
                fmt.Print("| ")
                player = board[j/4*3+i]
                if player == 0 {
                    player_mark = "O"
                } else {
                    player_mark = "X"
                }
                fmt.Print(player_mark)
                fmt.Print(" ")
            }
            fmt.Println("|")
        }
    }
    fmt.Println("+---+---+---+")
    winningBoard()
}

func sum(slice []int) (sum int) {
    for _, v := range slice {
        sum += v
    }
    return
}
func winningBoard() {
    var set []int

    for i :=0; i<3; i+=1 {
        set = board[i*3:(i+1)*3]
        if sum(set) == 0 {
            fmt.Println("O wins on line", i)
        }
        if sum(set) == 3 {
            fmt.Println("X wins on line", i)
        }
        fmt.Println(set, sum(set)==0, sum(set)==3)
        set = nil
    }
    for i :=0; i<3; i+=1 {
        set = append(set, board[0+i], board[3+i], board[6+i])
        if sum(set) == 0 {
            fmt.Println("O wins on column", i)
        }
        if sum(set) == 3 {
            fmt.Println("X wins on column", i)
        }
        fmt.Println(set, sum(set)==0, sum(set)==3)
        set = nil
    }
    // Testing diagonals
    d1 := []int{board[0], board[4], board[8]}
    d2 := []int{board[2], board[4], board[6]}
    diags := [][]int{d1,d2}
    for i, d := range diags {
        fmt.Println("for", i, d)
        if sum(d) == 0 {
            fmt.Printf("O wins on diagonal %d,%d-%d,%d\n", 1+i,1+i,3-2*i,3-2*i)
        }
        if sum(d) == 3 {
            fmt.Printf("X wins on diagonal %d,%d-%d,%d\n", 1+i,1+i,3-2*i,3-2*i)
        }
        fmt.Println(d, sum(d)==0, sum(d)==3)
    }

}
func mainLoop() {
}

func main() {
  initScreen()
  mainLoop()
}
