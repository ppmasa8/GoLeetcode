func destCity(paths [][]string) string {
  map := map[string]bool {}

  for _, v := range paths {
    map[v[0]] = true
  }

  for _, v := range paths {
    _, ok := map[v[1]]
    if !ok {
      return v[1]
    }
  }

  return ""
}
