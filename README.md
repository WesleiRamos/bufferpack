# bufferpack

Module to pack values to buffers

### Install package

``` bash
> go get github.com/WesleiRamos/bufferpack
```

### Usage:

Functions ends with `BE` = Big endian<br>
Funtcions ends with `LE` = Little endian

NOTICE: Tested only with small values

``` go

package main

import "fmt"
import "github.com/WesleiRamos/bufferpack"

func main() {
	f := "kk eae man"

	b := bufferpack.NewBuffer([]byte{})
	b.WriteShortBE(len(f))
	b.WriteString(f)

	fmt.Printf("%s\n", b.ReadStringBE()) // kk eae man

	d := bufferpack.NewBuffer(b.Bytes())
	fmt.Printf("%s\n", d.ReadStringBE()) // kk eae man

	fmt.Printf("%q\n", d.Bytes()) // "\x00\nkk eae man"
}
```
