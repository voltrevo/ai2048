package main

import (
    "./ai2048"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"
)

func setup() {
    logf, err := os.OpenFile("ai2048.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        panic("Couldn't open ai2048.log: " + fmt.Sprint(err))
    }
    log.SetOutput(logf)
    log.Println("Logging initialised")

    seed := time.Now().UTC().UnixNano()
    rand.Seed(seed)
    log.Println("rand seeded with", seed)
}

type Bot struct {
    br           ai2048.BoardRater
    rng          *rand.Rand
    mutationRate float64
}

func (b *Bot) Breed(b2o Organism) Organism {
    child := &Bot{}
    child.rng = b.rng
    child.mutationRate = b.mutationRate

    b2 := b2o.(*Bot)

    for i := 0; i != len(b.br); i++ {
        if b.rng.Intn(2) == 0 {
            child.br[i] = b.br[i]
        } else {
            child.br[i] = b2.br[i]
        }
    }

    child.br.Mutate(child.mutationRate, child.rng)

    return child
}

func (b *Bot) Extrapolate(b2o Organism) Organism {
    child := &Bot{}
    child.rng = b.rng
    child.mutationRate = b.mutationRate

    b2 := b2o.(*Bot)

    for i := 0; i != len(b.br); i++ {
        child.br[i] = b2.br[i] * b2.br[i] / b.br[i]
    }

    return child
}

func (b *Bot) Measure() float64 {
    return float64(ai2048.PlayGame(b.rng, &b.br).Score)
}

func playGameAsync(n int, br ai2048.BoardRater, rng *rand.Rand, c chan<- int) {
    testBr := br
    mutationRate := 0.01
    testBr.Mutate(mutationRate, rng)
    overallSum := 0

    for i := 0; i != n / 2000; i++ {
        sum := 0
        sumTest := 0
        for j := 0; j != 1000; j++ {
            b := ai2048.PlayGame(rng, &br)
            bTest := ai2048.PlayGame(rng, &testBr)
            sum += b.Score
            sumTest += bTest.Score
        }

        overallSum += sum + sumTest

        if sumTest > sum {
            //fmt.Println("Test params", testBr, "outperformed", br, "scoring", sumTest, "against", sum)
            br = testBr
        } else {
            //fmt.Println("Params", br, "outperformed", testBr, "scoring", sum, "against", sumTest)
            testBr = br
        }

        fmt.Println(sum)

        //fmt.Println("Mutating test params")
        //fmt.Println("Before:", testBr)
        testBr.Mutate(mutationRate, rng)
        //fmt.Println("After:", testBr)
        //fmt.Println()
    }

    br.Normalise()

    fmt.Println(br)
    fmt.Println(float64(overallSum) / float64(n))

    c <- overallSum
}

func main() {
    setup()

    //sum := 0.0
    //games := 200000

    initialRater := ai2048.BoardRater{
          /*
          1.3654369821388026,    // raw score
        114.49779070335288,      // empty squares
        113.89329471986436,      // raw score * empty squares

        // tile placement
          0.0010884043917799581, // centre weight
          0.0011352785416723196, // side weight
          0.0014107455090546657, // corner weight

        // neighbours
         -0.10097048863523984,   // zero neighbour
          0.10423399474189894,   // equal neighbour
         -0.10718112725853235,   // double neighbour
         -0.06824539779479108,   // 4x neighbour
         -0.11002006538584412,   // 8x neighbour
         -0.12172405253281565,   // 16x or more neighbour
         */
         /*
         1.0,
         1.0,
         1.0,

         1.0,
         1.0,
         1.0,

         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         1.0,
         */
        25.492481753838803,
        7.172761851594679,
        0.09503409140084336,

        -15.309597480724443,
        -1.7141242253439901,
        0.2258882543764266,

        -0.791763505641042,
        -8.344120401200911,
        0.19302060327520187,
        -0.015623732640613895,
        -0.23247295714412963,
        -0.9929199242858658,
        -2.3813238371649845,
        -5.883362015596853,
        2.897679304721464,
        0.8100697901165972,
    }

    /*
    c := make(chan int)
    go playGameAsync(games, initialRater, rand.New(rand.NewSource(rand.Int63())), c)
    //go playGameAsync(games / 2, initialRater, rand.New(rand.NewSource(rand.Int63())), c)

    sum += float64(<-c)
    //sum += float64(<-c)

    //fmt.Println(sum / float64(games))
    */

    rng := rand.New(rand.NewSource(rand.Int63()))

    adam := Bot{
        initialRater,
        rng,
        0.2}

    /*
    fmt.Println(Evolve(
        &adam,
        20,  // population size
        4,   // breeding size
        100, // generations
        rng))
    */
    Compete(&adam, 200, 50000000, rng)

    /*
    rng := rand.New(rand.NewSource(rand.Int63()))
    ai2048.PlayGame(rng, &initialRater)
    */
}
