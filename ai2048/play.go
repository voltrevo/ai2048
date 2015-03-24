package ai2048

import "fmt"
import "math/rand"

var best int = 0

func PlayGame(rng *rand.Rand, br *BoardRater) Board {
    b := CreateBoard(rng)

    for {
        bestDir := Direction(-1)
        bestScore := -1.0

        depth := 5
        for bestDir == Direction(-1) && depth != -1 {
            for d := Direction(0); d != Direction(4); d++ {
                bb := b
                if success := bb.PlayUnluckyTurnLite(d); success {
                    if score := recursiveRating(&bb, br, depth); score > bestScore {
                        bestDir = d
                        bestScore = score
                    }
                }
            }

            depth--
        }

        if bestDir == Direction(-1) {
            break
        }

        b.PlayTurn(bestDir)
    }

    if b.Score > best {
        best = b.Score
        fmt.Println("New best score")
        fmt.Println(&b)
    }

    return b
}

func recursiveRating(b *Board, br *BoardRater, depth int) float64 {
    if depth == 0 {
        return br.Rate(b)
    }

    bestScore := -1.0
    for d := Direction(0); d != Direction(4); d++ {
        bb := *b
        if success := bb.PlayUnluckyTurnLite(d); success {
            if score := recursiveRating(&bb, br, depth - 1); score > bestScore {
                bestScore = score
            }
        }
    }

    return bestScore
}

func shuffleInts(x []int, rng *rand.Rand) {
    limit := len(x) - 1
    length := len(x)
    for i := 0; i != limit; i++ {
        j := i + rng.Intn(length - i)
        x[i], x[j] = x[j], x[i]
    }
}

func PlayGameNaive(rng *rand.Rand) Board {
    b := CreateBoard(rng)

    for {
        dirs := [...]int{3, 0, 1, 2}
        moved := false

        for i := 0; i != len(dirs); i++ {
            if b.PlayTurn(Direction(dirs[i])) {
                moved = true
                break
            }
        }

        if !moved {
            break
        }

        dirs = [...]int{3, 1, 0, 2}
        moved = false

        for i := 0; i != len(dirs); i++ {
            if b.PlayTurn(Direction(dirs[i])) {
                moved = true
                break
            }
        }

        if !moved {
            break
        }
    }

    return b
}
