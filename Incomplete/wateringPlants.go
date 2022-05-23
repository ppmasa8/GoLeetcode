func wateringPlants(plants []int, capacity int) int {
    water := capacity
    steps := 0
    
    for i := 0; i < len(plants); i++ {
        if plants[i] > water {
            steps += i * 2
            water = capacity
        }
        steps++
        water -= plants[i]
    }
    return steps
}
