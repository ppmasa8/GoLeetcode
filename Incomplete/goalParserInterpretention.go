func interpret(command string) string {
    res := ""
    i := 0
    for i < len(command) {
        if string(command[i]) == "G" {
            res = res + "G"
            i++
        } else if string(command[i]) == "(" && string(command[i+1]) == ")" {
            res = res + "o"
            i+=2
        } else if string(command[i]) == "(" && string(command[i+3]) == ")" && string(command[i+1]) == "a" && string(command[i+2]) == "l" {
            res = res + "al"
            i+=4
        }
    }
    return res
}
