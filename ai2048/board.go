package ai2048

import (
    //"log"
    "math/rand"
    "fmt"
)

type Board struct {
    Vals [4][4]int
    Score int
    Turns int
    TileCount int
    rng *rand.Rand
}

func CreateBoard(rng *rand.Rand) Board {
    b := Board{}
    b.rng = rng
    b.insertRandomTile()
    return b
}

func CreateBoardWithLayout(layout [4][4]int, rng *rand.Rand) Board {
    b := Board{layout, 0, 0, 0, rng}
    b.TileCount = b.countTiles()
    return b
}

func (b *Board) PlayTurn(d Direction) bool {
    ret := b.moveTiles(d)
    if !ret {
        return false
    }

    b.insertRandomTile()
    b.Turns++

    return true
}

func (b *Board) PlayUnluckyTurnLite(d Direction) bool {
    ret := b.moveTiles(d)
    if !ret {
        return false
    }

    b.insertUnluckyTileLite()
    b.Turns++

    return true
}

func (b *Board) PlayUnluckyTurn(d Direction, br *BoardRater) bool {
    ret := b.moveTiles(d)
    if !ret {
        return false
    }

    b.insertUnluckyTile(br)
    b.Turns++

    return true
}

func (b *Board) String() string {
    return (
        fmt.Sprint(b.Vals[0]) + "\n" +
        fmt.Sprint(b.Vals[1]) + "\n" +
        fmt.Sprint(b.Vals[2]) + "\n" +
        fmt.Sprint(b.Vals[3]) + "\n" +
        "Score: " + fmt.Sprint(b.Score) + "\n" +
        "Turns: " + fmt.Sprint(b.Turns) + "\n")
}

func (b *Board) countTiles() int {
    count := 0
    for i := 0; i != 4; i++ {
        for j := 0; j != 4; j++ {
            if b.Vals[i][j] != 0 {
                count++
            }
        }
    }
    return count
}

func (b *Board) moveTiles(d Direction) bool {

    moved := false

    if d == Up {
        for j := 0; j != 4; j++ {
            ii := 0
            for i := 0; i != 4; i++ {
                for ii < 4 && b.Vals[ii][j] == 0 {
                    ii++
                }

                for iii := ii + 1; iii < 4; iii++ {
                    if b.Vals[iii][j] != 0 {
                        if b.Vals[iii][j] == b.Vals[ii][j] {
                            b.Vals[iii][j] = 0
                            b.Vals[ii][j] *= 2
                            b.Score += b.Vals[ii][j]
                            b.TileCount--
                            moved = true
                        }
                        break
                    }
                }

                if ii < 4 {
                    if ii != i {
                        moved = true
                    }
                    b.Vals[i][j] = b.Vals[ii][j]
                } else {
                    b.Vals[i][j] = 0
                }

                ii++
            }
        }
    } else if d == Down {
        for j := 0; j != 4; j++ {
            ii := 3
            for i := 3; i >= 0; i-- {
                for ii >= 0 && b.Vals[ii][j] == 0 {
                    ii--
                }

                for iii := ii - 1; iii >= 0; iii-- {
                    if b.Vals[iii][j] != 0 {
                        if b.Vals[iii][j] == b.Vals[ii][j] {
                            b.Vals[iii][j] = 0
                            b.Vals[ii][j] *= 2
                            b.Score += b.Vals[ii][j]
                            b.TileCount--
                            moved = true
                        }
                        break
                    }
                }

                if ii >= 0 {
                    if ii != i {
                        moved = true
                    }
                    b.Vals[i][j] = b.Vals[ii][j]
                } else {
                    b.Vals[i][j] = 0
                }

                ii--
            }
        }
    } else if d == Left {
        for i := 0; i != 4; i++ {
            jj := 0
            for j := 0; j != 4; j++ {
                for jj < 4 && b.Vals[i][jj] == 0 {
                    jj++
                }

                for jjj := jj + 1; jjj < 4; jjj++ {
                    if b.Vals[i][jjj] != 0 {
                        if b.Vals[i][jjj] == b.Vals[i][jj] {
                            b.Vals[i][jjj] = 0
                            b.Vals[i][jj] *= 2
                            b.Score += b.Vals[i][jj]
                            b.TileCount--
                            moved = true
                        }
                        break
                    }
                }

                if jj < 4 {
                    if jj != j {
                        moved = true
                    }
                    b.Vals[i][j] = b.Vals[i][jj]
                } else {
                    b.Vals[i][j] = 0
                }

                jj++
            }
        }
    } else if d == Right {
        for i := 0; i != 4; i++ {
            jj := 3
            for j := 3; j >= 0; j-- {
                for jj >= 0 && b.Vals[i][jj] == 0 {
                    jj--
                }

                for jjj := jj - 1; jjj >= 0; jjj-- {
                    if b.Vals[i][jjj] != 0 {
                        if b.Vals[i][jjj] == b.Vals[i][jj] {
                            b.Vals[i][jjj] = 0
                            b.Vals[i][jj] *= 2
                            b.Score += b.Vals[i][jj]
                            b.TileCount--
                            moved = true
                        }
                        break
                    }
                }

                if jj >= 0 {
                    if jj != j {
                        moved = true
                    }
                    b.Vals[i][j] = b.Vals[i][jj]
                } else {
                    b.Vals[i][j] = 0
                }

                jj--
            }
        }
    }

    return moved
}

func (b* Board) insertRandomTile() {
    //log.Println("insertRandomTile()")

    if b.TileCount == 16 {
        return
    }

    newTile := 2
    if b.rng.Intn(10) == 0 {
        newTile = 4
    }

    tileIndex := b.rng.Intn(16 - b.TileCount)
    emptyCount := 0
    for i := 0; i != 4; i++ {
        for j := 0; j != 4; j++ {
            if b.Vals[i][j] == 0 {
                if emptyCount == tileIndex {
                    b.Vals[i][j] = newTile
                    b.TileCount++
                    return
                }
                emptyCount++
            }
        }
    }

    panic("Couldn't insert random tile")
}

func (b* Board) insertUnluckyTile(br *BoardRater) {
    //log.Println("insertRandomTile()")

    if b.TileCount == 16 {
        return
    }

    worstI := -1
    worstJ := -1
    worstTile := -1
    worstRating := 0.0

    for i := 0; i != 4; i++ {
        for j := 0; j != 4; j++ {
            if b.Vals[i][j] == 0 {
                b.Vals[i][j] = 2
                rating := br.Rate(b)
                if worstI == -1 || rating < worstRating {
                    worstI = i
                    worstJ = j
                    worstTile = 2
                    worstRating = rating
                }

                b.Vals[i][j] = 4
                rating = br.Rate(b)
                if worstI == -1 || rating < worstRating {
                    worstI = i
                    worstJ = j
                    worstTile = 4
                    worstRating = rating
                }

                b.Vals[i][j] = 0
            }
        }
    }

    if worstI == -1 {
        panic("Couldn't insert random tile")
    }

    b.Vals[worstI][worstJ] = worstTile
}

func (b* Board) insertUnluckyTileLite() {
    //log.Println("insertRandomTile()")

    if b.TileCount == 16 {
        return
    }

    bestI := -1
    bestJ := -1
    bestRating := 0

    for i := 0; i != 4; i++ {
        for j := 0; j != 4; j++ {
            if b.Vals[i][j] == 0 {
                sum := 0
                biggest := 0
                edges := 0

                if i != 0 {
                    v := b.Vals[i - 1][j]
                    if v > biggest {
                        biggest = v
                    }
                    sum += v
                } else {
                    edges++
                }

                if i != 3 {
                    v := b.Vals[i + 1][j]
                    if v > biggest {
                        biggest = v
                    }
                    sum += v
                } else {
                    edges++
                }

                if j != 0 {
                    v := b.Vals[i][j - 1]
                    if v > biggest {
                        biggest = v
                    }
                    sum += v
                } else {
                    edges++
                }

                if j != 3 {
                    v := b.Vals[i][j + 1]
                    if v > biggest {
                        biggest = v
                    }
                    sum += v
                } else {
                    edges++
                }

                sum += edges * biggest

                if sum > bestRating {
                    bestRating = sum
                    bestI = i
                    bestJ = j
                }
            }
        }
    }

    if bestI == -1 {
        panic("Couldn't insert random tile")
    }

    twoNeighbour := (
        (bestI != 0 && b.Vals[bestI - 1][bestJ] == 2) ||
        (bestI != 3 && b.Vals[bestI + 1][bestJ] == 2) ||
        (bestJ != 0 && b.Vals[bestI][bestJ - 1] == 2) ||
        (bestJ != 3 && b.Vals[bestI][bestJ + 1] == 2))

    if twoNeighbour {
        b.Vals[bestI][bestJ] = 4
    } else {
        b.Vals[bestI][bestJ] = 2
    }
}
