# 📌 Go Standard Library Cheatsheet

A collection of the most popular and frequently useful functions/methods from Go's standard library.  

---

## ✅ `strings` package

```go
strings.ToUpper("go")              // "GO"
strings.ToLower("HELLO")           // "hello"
strings.TrimSpace("  hi  ")        // "hi"
strings.HasPrefix("gopher", "go")  // true
strings.HasSuffix("gopher", "er")  // true
strings.Contains("hello", "ll")    // true
strings.Count("banana", "a")       // 3
strings.Repeat("na", 3)            // "nanana"
strings.ReplaceAll("foo bar foo", "foo", "baz") // "baz bar baz"
strings.Split("a,b,c", ",")        // []string{"a","b","c"}
strings.Join([]string{"a","b"}, "-") // "a-b"
strings.Fields("a b   c")          // []string{"a","b","c"}
```


## ✅ `strconv` package
```go
strconv.Itoa(123)                  // "123"
strconv.Atoi("456")                // 456
strconv.FormatInt(15, 2)           // "1111"
strconv.ParseInt("1111", 2, 64)    // 15
```

## ✅ `fmt` package
```go
fmt.Println("Hello", "Go")         // Hello Go
fmt.Printf("Pi: %.2f\n", 3.14159)  // Pi: 3.14
msg := fmt.Sprintf("Hi %s", "Bob") // "Hi Bob"
```

## ✅ `math` package
```go
math.Max(10, 20)   // 20
math.Min(10, 20)   // 10
math.Abs(-42)      // 42
math.Pow(2, 3)     // 8
math.Sqrt(16)      // 4
math.Round(2.7)    // 3
math.Floor(2.9)    // 2
math.Ceil(2.1)     // 3
```

## ✅ `time` package
```go
now := time.Now()                  // current time
time.Sleep(2 * time.Second)        // pause
elapsed := time.Since(now)         // duration since
future := time.Until(now.Add(1*time.Hour))
t, _ := time.Parse("2006-01-02", "2025-09-09")
fmt.Println(now.Format("2006-01-02")) // format date
```

## ✅ `sort` package
```go
nums := []int{3, 1, 4, 2}
sort.Ints(nums)                     // [1,2,3,4]

words := []string{"banana","apple"}
sort.Strings(words)                 // [apple, banana]

sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })   // custom sort (descending)
```

## ✅ `bytes` package
```go
bytes.Contains([]byte("hello"), []byte("he"))   // true
bytes.Equal([]byte("a"), []byte("a"))           // true
bytes.Join([][]byte{[]byte("a"), []byte("b")}, []byte("-")) // "a-b"
bytes.Split([]byte("a,b"), []byte(","))         // [["a"],["b"]]
bytes.TrimSpace([]byte(" hi "))                 // "hi"
```

## ✅ `io\os` package
```go
data, _ := os.ReadFile("file.txt") // read file
os.WriteFile("out.txt", data, 0644) // write file
f, _ := os.Open("file.txt")        // open file
defer f.Close()

os.Create("new.txt")               // create file
os.Remove("old.txt")               // delete file

io.Copy(os.Stdout, f)              // copy file content to stdout
```


## ✅ `bufio` package
```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    fmt.Println(scanner.Text())     // read line by line
}
```


## ✅ `unicode` package
```go
unicode.ToUpper('a')                // 'A'
unicode.ToLower('A')                // 'a'
unicode.IsDigit('5')                // true
unicode.IsLetter('x')               // true
unicode.IsSpace(' ')                // true
```


## common Go format specifiers:


    %d → Integer (decimal)

    %f → Floating point

    %s → String

    %v → Default format (any value)

    %t → Boolean

    %x → Hexadecimal

    %T → Type of value
