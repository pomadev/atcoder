# 標準入力を受け取る

## 1行1つの数字を受け取る
```
1
```

```go
var n int
fmt.Scan(&n)
fmt.Println(n) // 1
```

## 1行複数の数字を受け取る
```
1 2 3
```

```go
var a, b, c int
fmt.Scan(&a, &b, &c)
fmt.Println(a, b, c) // 1 2 3
```

## 1行1つの文字列を受け取る
```
hello
```

```go
var s string
fmt.Scan(&s)
fmt.Println(s) // hello
```

## 1行複数の文字列を受け取る
```
hello world !
```

```go
var a, b, c string
fmt.Scan(&a, &b, &c)
fmt.Println(a, b, c) // hello world !
```
