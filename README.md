# regexop

This package `regexop` provides some set operations on **regular expressions**, such as :

-   `Union`
-   `Intersection`
-   `IsSubsetOf`
-   `IsEqualTo`
-   ...



--------------------

## Examples

```go
import regexop

var uRegex string = regexop.UnionRegex("a", "b", "c")
// uRegex = "a|b|c"

var s bool = regexop.IsSubsetOf("a+", "a*")
// s = true

var e bool = regexop.isEqualTo("a+", "aa*")
// e = true
```





--------------------

## References

- Regular Expression to NFA, NFA to DFA:
    -   the book, **Principle of Compiler** (Dragon Book), Alfred V. Aho / Monica S.Lam

- minimize DFA
    -   [DFA minimization:  from Brzozowski to Hopcroft](https://pdfs.semanticscholar.org/0d70/3ad6fe657bd98cf0bae1d2a883d6e950a212.pdf)


- DFA to Regex
    -   https://github.com/qntm/greenery  A python library with similar functions
    -   https://arxiv.org/pdf/1502.03573.pdf  A paper about Automata and Regular Expression
    -   the book, **Introduction to the Theory of Computation**, Michael Sipser 





----------------

## Note

Although there are many improved efficient algorithms about automata in some published papers, this package implements the basic algorithms whose efficiency are lower . 