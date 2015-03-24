package ai2048

func emptySquares(b *Board) float64 {
    return float64(16 - b.TileCount)
}

var boardIndices [4][4]int = [4][4]int{
    {2, 1, 1, 2},
    {1, 0, 0, 1},
    {1, 0, 0, 1},
    {2, 1, 1, 2},
}

func tilePlacement(b *Board, weights []float64) float64 {
    sum := 0.0
    k := 0
    for i := 0; i != 4; i++ {
        for j := 0; j != 4; j++ {
            sum += weights[boardIndices[i][j]] * float64(b.Vals[i][j])
            k++
        }
    }
    return sum
}

func goodNeighbours(b *Board, weights []float64) float64 {

    scorePair := func(a, b int) float64 {

        if a == 0 || b == 0 {
            return weights[0] * float64(a + b)
        }

        ratio := a / b + b / a // one of these will be zero, the other will be max / min

        if ratio == 1 {
            return weights[1] * float64(a + b)
        } else if ratio == 2 {
            return weights[2] * float64(a + b)
        } else if ratio == 4 {
            return weights[3] * float64(a + b)
        } else if ratio == 8 {
            return weights[4] * float64(a + b)
        } else if ratio == 16 {
            return weights[5] * float64(a + b)
        } else if ratio == 32 {
            return weights[6] * float64(a + b)
        }

        return weights[7] * float64(a + b)
    }

    sum := 0.0

    for i := 0; i != 3; i++ {
        for j := 0; j != 4; j++ {
            sum += scorePair(b.Vals[i][j], b.Vals[i + 1][j])
            sum += scorePair(b.Vals[j][i], b.Vals[j][i + 1])
        }
    }

    return sum
}

func monotonicity(b *Board, weights []float64) float64 {
    scoreLine := func(a, b, c, d int) float64 {
        if (a >= b && b >= c && c >= d) || (a <= b && b <= c && c <= d) {
            return (weights[0] + weights[1] * float64(a + b + c + d))
        }
        return 0
    }

    sum := 0.0
    for i := 0; i != 4; i++ {
        sum += scoreLine(b.Vals[i][0], b.Vals[i][1], b.Vals[i][2], b.Vals[i][3])
        sum += scoreLine(b.Vals[0][i], b.Vals[1][i], b.Vals[2][i], b.Vals[3][i])
    }

    return sum
}

func biggestBlockInCorner(b *Board) float64 {
    biggest := -1

    for i := 0; i != 4; i++ {
        for j := 0; j != 4; j++ {
            if b.Vals[i][j] > biggest {
                biggest = b.Vals[i][j]
            }
        }
    }

    if b.Vals[0][0] == biggest || b.Vals[0][3] == biggest || b.Vals[3][0] == biggest || b.Vals[3][3] == biggest {
        return float64(biggest)
    }

    return float64(0.0)
}
