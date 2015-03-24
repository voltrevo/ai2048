package main

import (
    "fmt"
    "math/rand"
    "sort"
)

type Organism interface {
    Breed(Organism) Organism
    Measure() float64
}

type organismWithMeasure struct {
    org     Organism
    measure float64
}

type organismCollection []organismWithMeasure

func (oc organismCollection) Less(i, j int) bool {
    return oc[i].measure > oc[j].measure
}

func (oc organismCollection) Len() int { return len(oc) }
func (oc organismCollection) Swap(i, j int) { oc[i], oc[j] = oc[j], oc[i] }

func Evolve(adam Organism, popSize, breedSize, generations int, rng *rand.Rand) Organism {
    population := make(organismCollection, 0, popSize)
    population = append(population, organismWithMeasure{adam, 0.0})

    for i := 1; i != popSize; i++ {
        population = append(population, organismWithMeasure{adam.Breed(adam), 0.0})
    }

    for i := 0; i != generations; i++ {
        sum := 0.0

        for j := 0; j != popSize; j++ {
            population[j].measure = population[j].org.Measure()
            sum += population[j].measure
        }

        sort.Sort(population)

        for k := breedSize; k < popSize; k++ {
            population[k].org = population[rng.Intn(breedSize)].org.Breed(population[rng.Intn(breedSize)].org)
        }

        fmt.Println(sum / float64(popSize))
    }

    return population[0].org
}
