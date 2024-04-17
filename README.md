<div align=center><div style="width:60%;">
<img src="sample_text/coverimg.jpg" alt="Logo Image" width="40%"/>
<h1>gOwO</h1><p>
<i>OwO whats this?</i> An advanced owofier written in Go featuring syllable based substitutions for more natural owofication. 
</p></div></div>


### CLI Usage

``` 
Usage: owo [OPTION]... [TEXT|FILE]...

Owoify text passed as an argument.
       Example: owo "text to owoify"

Or as part of a pipeline operation:
        Example: owo --help | owo

Options:
        -h, --help              show this message
        -s, --stats             dry run showing replacement statistics
        -i, --infile [PATH]     load text from file
        -o, --outfile [PATH]    output to file
If no text is provided as an argument read from standard input.
Exit status is 0 if no error occurs, 1 otherwise.
```

### Go Module Usage
#### Minimal Example
```go 
package main

import (
	gowo "github.com/RootInit/gOwO"
)

func main() {
	// Initialize Owofier
	owo := gowo.DefaultOwofier()
    // Convert string
    result, err = owo.TranslateString("hello world")
    if err != nil {
        panic()
    }
    // Display string
    println(result)
}
```
#### Complete Example
See `./cmd` for a complete program using the gOwO package.

### Sample Output
     
An sample markdown document and epub of the converted complete works of William Shakespeare is provided in `./sample_text` This was obtained from [Project Gutenberg](https://www.gutenberg.org/ebooks/100).

> Now, I am become gOwO, the destroyer of classic literature.
