package main

import "fmt"

/*
func DNAtoRNA(dna string) string {

	rna := dna
	runes := []rune(rna)
	box := []rune{}

	for i := 0; i < len(runes); i++ {
		if i != 'T' {
			box = append(box, runes[i])
		}
		if i == 'U' {
			box = append(box, 'U')
		}
	}
	dna = string(box[:])
	return dna

}
*/

/*

  retval := ""
  for _, ch := range dna {
    switch ch {
      case 'G':
      retval += "G"
      case 'C':
      retval += "C"
      case 'A':
      retval += "A"
      case 'T':
      retval += "U"
      case 'U':
      retval += "U"
    }
  }
  return retval
}
 * /






func DNAtoRNA(dna string) string {
	runes := []rune(dna)
	box := []rune{}

	for i := 0; i < len(runes); i++ {
		if i != 'T' {
			box = append(box, runes[i])
		}
		if i == 'U' {
			box = append(box, 'U')
		}
	}
	dna = string(box[:])
	return dna
}

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}
