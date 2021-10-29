# dyqual
Gomega Equal with dyff

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
