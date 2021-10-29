# dyqual
[Gomega Equal](https://onsi.github.io/gomega/) with [dyff](https://github.com/homeport/dyff)

## Usage

```golang
import (
    . "github.com/onsi/gomega"
    . "github.com/tenstad/dyqual/matchers"
)
```

```golang
Expect(ACTUAL).To(Dyqual(EXPECTED))
```

## Standad Gomega `Equal()`

```txt
 Expected
      <string>: a: a\nb: b\nc: c\nd: d\ne: e\n
  to equal
      <string>: a: a\nb: b\nc: q\nd: d\ne: \"\"\n
```

## `Dyqual()`

```txt
  c
    ± value change
      - c
      + q
  
  e
    ± value change
      - e
      +
```
