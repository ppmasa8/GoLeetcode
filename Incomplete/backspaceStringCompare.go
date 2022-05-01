func backspaceCompare(s string, t string) bool {
    listS := []string{}
    listT := []string{}
    for _, v := range s {
        if string(v) == "#" {
            if len(listS) != 0 {
                listS = listS[:len(listS)-1]
            }
        } else {
            listS = append(listS, string(v))
        }
    }
    
    for _, v := range t {
        if string(v) == "#" {
            if len(listT) != 0 {
                listT = listT[:len(listT)-1]
            }
        } else {
            listT = append(listT, string(v))
        }
    }
    
    return reflect.DeepEqual(listS, listT)
}
