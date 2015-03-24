package ai2048

import "math/rand"

type BoardRater [16]float64

func (br *BoardRater) Rate(b *Board) float64 {
    return (
        br[0] * float64(b.Score) +
        br[1] * emptySquares(b) +
        br[2] * float64(b.Score) * emptySquares(b) +
        tilePlacement(b, br[3:6]) +
        goodNeighbours(b, br[6:14]) +
        monotonicity(b, br[14:16]))
}

func (br *BoardRater) Mutate(x float64, rng *rand.Rand) {
    if rng.Intn(2) == 0 {
        i := rng.Intn(len(br))
        br[i] *= 1.0 + x * rng.NormFloat64()

        if rng.Float64() < 0.1 * x {
            br[i] *= -1.0
        }
    } else {
        for i := 0; i != len(br); i++ {
            br[i] *= 1.0 + x * rng.NormFloat64()

            if rng.Float64() < 0.1 * x {
                br[i] *= -1.0
            }
        }
    }
    /*
    i := len(br) - 1
    br[i] *= 1.0 + x * rng.NormFloat64()

    if rng.Float64() < 0.1 * x {
        br[i] *= -1.0
    }
    */
}

func (br *BoardRater) Normalise() {
    firstParam := br[0]
    for i := 0; i != len(br); i++ {
        br[i] /= firstParam
    }
}
