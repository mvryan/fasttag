# fasttag
A Go implementation of a Brill Part of Speech Tagger


This Part of Speech tagger is a Go port of the [fasttag_v2](https://github.com/mark-watson/fasttag_v2) library by [Mark Watson](http://markwatson.com).


It keeping with the spirit of openness he provided this library dual licensed under both the [LGPLv3](https://www.gnu.org/licenses/lgpl-3.0.en.html) and [Apache2](http://www.apache.org/licenses/LICENSE-2.0) licenses.

## Sample.go
### Source
```go
package main

import (
        "fmt"
        "github.com/mvryan/fasttag"
)

func main() {
        fmt.Println("Hello, world")
        words := fasttag.WordsToSlice("Hello, world")
        fmt.Println("Words: ", words)
        pos := fasttag.BrillTagger(words)
        fmt.Println("Parts of Speech: ", pos)
}
```
### Output
```
Hello, world
Words:  [Hello , world]
Parts of Speech:  [UH , NN]
```
