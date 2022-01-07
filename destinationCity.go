func destCity(paths [][]string) string {
  m := map[string]bool {}

  for _, v := range paths {
    m[v[0]] = true
  }

  for _, v := range paths {
    _, ok := m[v[1]]
    if !ok {
      return v[1]
    }
  }

  return ""
}
