GoDecode2
========
[![Build Status](https://travis-ci.org/xuender/godecode2.png)](https://travis-ci.org/xuender/godecode2)

ASCII transliterations of Unicode text for go lang.

## How to Use ##

### Install

```shell
go get github.com/xuender/godecode2
```

### Decode
```go
import (
    "fmt"
    gd "github.com/xuender/godecode2"
)
fmt.Println(gd.Decode("南无阿弥陀佛"))
Nan Wu A Mi Tuo Fo

fmt.Println(gd.Decode("一条会走路的鱼"));
Yi Tiao Hui Zou Lu De Yu

fmt.Println(gd.Decode("あみだにょらい"));
amidaniyorai
```
### Initials
```go
fmt.Println(gd.Initials("南无阿弥陀佛"));
NWAMTF

fmt.Println(gd.Initials("不怨人就是成佛的大道根"));
BYRJSCFDDDG

fmt.Println(gd.Initials("Κνωσός"));
K
```

## Reference ##

[GoDecode](https://github.com/xuender/godecode)

[Unidecode for Java](https://github.com/xuender/unidecode)

[Unidecode for Python](https://pypi.python.org/pypi/Unidecode)
