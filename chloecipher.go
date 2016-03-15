package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cipher interface {
  Encode(s string)
  Decode(s string)
}

type ChloeCipherV1 struct {
  runemap map[rune]rune
}

func NewChloeCipherV1() *ChloeCipherV1 {
  ans := new(ChloeCipherV1)
  runemap := make(map[rune]rune)

  runemap[rune('a')] = rune('z')
  runemap[rune('b')] = rune('y')
  runemap[rune('c')] = rune('x')
  runemap[rune('d')] = rune('w')
  runemap[rune('e')] = rune('v')
  runemap[rune('f')] = rune('u')
  runemap[rune('g')] = rune('t')
  runemap[rune('h')] = rune('s')
  runemap[rune('i')] = rune('r')
  runemap[rune('j')] = rune('q')
  runemap[rune('k')] = rune('p')
  runemap[rune('l')] = rune('o')
  runemap[rune('m')] = rune('n')
  runemap[rune('n')] = rune('m')
  runemap[rune('o')] = rune('l')
  runemap[rune('p')] = rune('k')
  runemap[rune('q')] = rune('j')
  runemap[rune('r')] = rune('i')
  runemap[rune('s')] = rune('h')
  runemap[rune('t')] = rune('g')
  runemap[rune('u')] = rune('f')
  runemap[rune('v')] = rune('e')
  runemap[rune('w')] = rune('d')
  runemap[rune('x')] = rune('c')
  runemap[rune('y')] = rune('b')
  runemap[rune('z')] = rune('a')

  runemap[rune('A')] = rune('Z')
  runemap[rune('B')] = rune('Y')
  runemap[rune('C')] = rune('X')
  runemap[rune('D')] = rune('W')
  runemap[rune('E')] = rune('V')
  runemap[rune('F')] = rune('U')
  runemap[rune('G')] = rune('T')
  runemap[rune('H')] = rune('S')
  runemap[rune('I')] = rune('R')
  runemap[rune('J')] = rune('Q')
  runemap[rune('K')] = rune('P')
  runemap[rune('L')] = rune('O')
  runemap[rune('M')] = rune('N')
  runemap[rune('N')] = rune('M')
  runemap[rune('O')] = rune('L')
  runemap[rune('P')] = rune('K')
  runemap[rune('Q')] = rune('J')
  runemap[rune('R')] = rune('I')
  runemap[rune('S')] = rune('H')
  runemap[rune('T')] = rune('G')
  runemap[rune('U')] = rune('F')
  runemap[rune('V')] = rune('E')
  runemap[rune('W')] = rune('D')
  runemap[rune('X')] = rune('C')
  runemap[rune('Y')] = rune('B')
  runemap[rune('Z')] = rune('A')

  ans.runemap = runemap

  return ans
}

func (c ChloeCipherV1) swapRune(r rune) rune {
  _, ok := c.runemap[r]
  if !ok {
    return r // character not in the map, so return as-is
  }
	return c.runemap[r]
}

func (c ChloeCipherV1) Encode(s string) string {
  return c.Decode(s)
}

func (c ChloeCipherV1) Decode(s string) string {
	var answer string

  for _, runeValue := range s {
    answer += fmt.Sprintf("%c", c.swapRune(runeValue))
  }

	return answer
}

func main() {
  c := NewChloeCipherV1()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("(Press Ctrl-D to quit)")
	for scanner.Scan() {
		fmt.Println(c.Decode(scanner.Text()))
    fmt.Println()
	}
}
