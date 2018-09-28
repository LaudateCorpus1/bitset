# bitset
A simple BitSet library fo Golang.

## how to use

### how to make a new bitset object

``` go
var totalSize uint64 = 99999
bs := bitset.New(totalSize)

```

### how to set value 

```go
// set true 

var pos uint64 = 555
bs.SetTrue(pos)


// set false
bs.SetFalse(pos)
```

### how to check value 

``` go

var pos uint64 = 666
isTrue := bs.IsTrue(pos)
if isTrue {
    //Have this number : pos
} else {
    //NO this number : po2
}
```

### how to search all value status.

``` go 
// search true value 

trueValueList := bs.SearchTruePos()

// search false value

falseValueList := bs.SearchFlasePos()
```


## benchmark


goos: darwin
goarch: amd64
pkg: bitset

|funcName/size of bitset | loop times | one operation time|
|:-|-:|-:|
|BenchmarkBitSet_SearchFalsePos/32768-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SearchFalsePos/65536-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SearchFalsePos/131072-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SearchFalsePos/262144-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SearchFalsePos/524288-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SearchFalsePos/1048576-4        |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SearchFalsePos/2097152-4        |2000000000|               0.01 ns/op|
|BenchmarkBitSet_SearchFalsePos/4194304-4        |2000000000|               0.01 ns/op|
|BenchmarkBitSet_SearchFalsePos/8388608-4        |2000000000|               0.02 ns/op|
|BenchmarkBitSet_SearchFalsePos/16777216-4       |2000000000|               0.03 ns/op|
|BenchmarkBitSet_SearchFalsePos/33554432-4       |2000000000|               0.06 ns/op|
|BenchmarkBitSet_SearchFalsePos/67108864-4       |2000000000|               0.21 ns/op|
|BenchmarkBitSet_SearchFalsePos/134217728-4      |1000000000|               0.52 ns/op|
|BenchmarkBitSet_SearchFalsePos/268435456-4      |       1  |      1085940148 ns/op|
|BenchmarkBitSet_SearchFalsePos/536870912-4      |       1  |      3625218790 ns/op|
|BenchmarkBitSet_SetTrue/size_:_32768-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_65536-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_131072-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_262144-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_524288-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_1048576-4        |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_2097152-4        |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_4194304-4        |2000000000|               0.00 ns/op|
|BenchmarkBitSet_SetTrue/size_:_8388608-4        |2000000000|               0.01 ns/op|
|BenchmarkBitSet_SetTrue/size_:_16777216-4       |2000000000|               0.01 ns/op|
|BenchmarkBitSet_SetTrue/size_:_33554432-4       |2000000000|               0.02 ns/op|
|BenchmarkBitSet_SetTrue/size_:_67108864-4       |1000000000|               0.08 ns/op|
|BenchmarkBitSet_SetTrue/size_:_134217728-4      |2000000000|               0.08 ns/op|
|BenchmarkBitSet_SetTrue/size_:_268435456-4      |2000000000|               0.18 ns/op|
|BenchmarkBitSet_SetTrue/size_:_536870912-4      |1000000000|               0.62 ns/op|
|BenchmarkBitSet_IsTrue/size_:_32768-4           |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_65536-4           |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_131072-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_262144-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_524288-4          |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_1048576-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_2097152-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_4194304-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_8388608-4         |2000000000|               0.00 ns/op|
|BenchmarkBitSet_IsTrue/size_:_16777216-4        |2000000000|               0.01 ns/op|
|BenchmarkBitSet_IsTrue/size_:_33554432-4        |2000000000|               0.01 ns/op|
|BenchmarkBitSet_IsTrue/size_:_67108864-4        |2000000000|               0.02 ns/op|
|BenchmarkBitSet_IsTrue/size_:_134217728-4       |2000000000|               0.04 ns/op|
|BenchmarkBitSet_IsTrue/size_:_268435456-4       |2000000000|               0.08 ns/op|
|BenchmarkBitSet_IsTrue/size_:_536870912-4       |2000000000|               0.18 ns/op|
