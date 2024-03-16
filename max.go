package piscine
func Max(a []int) int {
max:=a[0]
  for i:=range a{
    if max < a[i]{
      max = a[i]
    }
  } 
  return max
}