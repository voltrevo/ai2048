package main

import (
    "fmt"
    "math/rand"
)

type Organism interface {
    Breed(Organism) Organism
    Extrapolate(Organism) Organism
    Measure() float64
}

func Compete(adam Organism, trials, generations int, rng *rand.Rand) {
    leader := adam
    sum := 0.0
    for i := 0; i != trials; i++ {
        sum += leader.Measure()
        fmt.Print(".")
    }
    fmt.Println()
    leaderAvg := sum / float64(trials)
    fmt.Println(0, leaderAvg)
    fmt.Println(leader)
    fmt.Println()

    for gen := 0; gen != generations; gen++ {
        challenger := leader.Breed(leader)
        challengerSum := 0.0
        fmt.Print(gen)
        for i := 1; i <= trials; i++ {
            challengerSum += challenger.Measure()
            fmt.Print(".")
            if challengerSum / float64(i) < leaderAvg {
                break
            }
        }

        fmt.Println()

        challengerAvg := challengerSum / float64(trials)
        for challengerAvg > leaderAvg {
            oldLeader := leader
            leader = challenger

            sum = 0.0
            for i := 0; i != trials; i++ {
                sum += leader.Measure()
            }
            leaderAvg = sum / float64(trials)

            fmt.Println(gen, leaderAvg)
            fmt.Println(leader)
            fmt.Println()

            challenger = oldLeader.Extrapolate(leader)
            challengerSum = 0.0
            for i := 0; i != trials; i++ {
                challengerSum += challenger.Measure()
                fmt.Print(".")
            }
            fmt.Println()
            challengerAvg = challengerSum / float64(trials)
        }
    }
}
