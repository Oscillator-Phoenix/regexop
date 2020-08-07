package regexop

import (
	"fmt"
	"testing"
)

func TestSymbolSet(t *testing.T) {
	ss1 := newSymbolSetWithEpsilon()
	fmt.Println(ss1) // [ <esp> ]

	ss1.insert('p', 'q', 'r')
	fmt.Println(ss1) // [ <esp> p q r ]

	ss2 := newSymbolSet()

	fmt.Println(ss2) // [ ]

	ss2.insert('a')
	ss2.insert('狗')
	ss2.insert('の')
	ss2.insert('4')
	fmt.Println(ss2) // [ 狗 の 4 a ]

	ss2.erase('a')
	ss2.erase('b')
	fmt.Println(ss2) // [ 狗 の 4 ]

}

func TestUnionSymbolSet(t *testing.T) {
	ss1 := newSymbolSet()
	ss1.insert('a', 'b')
	ss1.insert('b')

	ss2 := newSymbolSet()
	ss2.insert('b')
	ss2.insert('c')

	ss3 := newSymbolSet()
	ss3.insert('b')
	ss3.insert('d')

	u1 := unionSymbolSet(ss1, ss2)
	fmt.Println(u1) // [ a b c ]

	u2 := unionSymbolSet(ss1, ss2, ss3)
	fmt.Println(u2) // [ a b c d ]
}

func TestIntersectionSymbolSet(t *testing.T) {
	ss1 := newSymbolSet()
	ss1.insert('a')
	ss1.insert('b')

	ss2 := newSymbolSet()
	ss2.insert('a')
	ss2.insert('b')
	ss2.insert('c')

	ss3 := newSymbolSet()
	ss3.insert('b')
	ss3.insert('d')

	i1 := intersectionSymbolSet(ss1, ss2)
	fmt.Println(i1) // [ a b ]

	i2 := intersectionSymbolSet(ss1, ss2, ss3)
	fmt.Println(i2) // [ b ]
}
