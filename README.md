# regexop

This package `regexop` provides some set operations on **regular expressions**, such as :

-   `Union`
-   `Intersection`
-   `IsSubsetOf`
-   `isEqualTo`
-   ...



Examples:

```go
import regexop

var uRegex string = regexop.UnionRegex("a", "b", "c")
// uRegex = "a|b|c"

bool s = regexop.IsSubsetOf("a+", "a*")
// s = true

bool e = regexop.isEqualTo("a+", "aa*")
// e = true
```