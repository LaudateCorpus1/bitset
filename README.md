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


### func SearchFalsePos
BenchmarkBitSet_SearchFalsePos/32768-4          2000000000               0.00 ns/op
2018/09/28 14:37:55 65536
BenchmarkBitSet_SearchFalsePos/65536-4          2000000000               0.00 ns/op
2018/09/28 14:37:56 131072
BenchmarkBitSet_SearchFalsePos/131072-4         2000000000               0.00 ns/op
2018/09/28 14:37:57 262144
BenchmarkBitSet_SearchFalsePos/262144-4         2000000000               0.00 ns/op
2018/09/28 14:37:58 524288
BenchmarkBitSet_SearchFalsePos/524288-4         2000000000               0.00 ns/op
2018/09/28 14:37:59 1048576
BenchmarkBitSet_SearchFalsePos/1048576-4        2000000000               0.00 ns/op
2018/09/28 14:38:00 2097152
BenchmarkBitSet_SearchFalsePos/2097152-4        2000000000               0.00 ns/op
2018/09/28 14:38:01 4194304
BenchmarkBitSet_SearchFalsePos/4194304-4        2000000000               0.01 ns/op
2018/09/28 14:38:02 8388608
BenchmarkBitSet_SearchFalsePos/8388608-4        2000000000               0.01 ns/op
2018/09/28 14:38:03 16777216
BenchmarkBitSet_SearchFalsePos/16777216-4       2000000000               0.03 ns/op
2018/09/28 14:38:05 33554432
BenchmarkBitSet_SearchFalsePos/33554432-4       2000000000               0.08 ns/op
2018/09/28 14:38:07 67108864
BenchmarkBitSet_SearchFalsePos/67108864-4       1000000000               0.30 ns/op
2018/09/28 14:38:13 134217728
BenchmarkBitSet_SearchFalsePos/134217728-4      1000000000               0.54 ns/op
2018/09/28 14:38:31 268435456
BenchmarkBitSet_SearchFalsePos/268435456-4             1        1739315899 ns/op
2018/09/28 14:38:34 536870912
BenchmarkBitSet_SearchFalsePos/536870912-4             1        4845138974 ns/op


### func SetTrue
BenchmarkBitSet_SetTrue/size_:_32768-4          2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_65536-4          2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_131072-4         2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_262144-4         2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_524288-4         2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_1048576-4        2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_2097152-4        2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_4194304-4        2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_8388608-4        2000000000               0.00 ns/op
BenchmarkBitSet_SetTrue/size_:_16777216-4       2000000000               0.01 ns/op
BenchmarkBitSet_SetTrue/size_:_33554432-4       2000000000               0.02 ns/op
BenchmarkBitSet_SetTrue/size_:_67108864-4       2000000000               0.04 ns/op
BenchmarkBitSet_SetTrue/size_:_134217728-4      2000000000               0.07 ns/op
BenchmarkBitSet_SetTrue/size_:_268435456-4      2000000000               0.15 ns/op
BenchmarkBitSet_SetTrue/size_:_536870912-4      1000000000               0.59 ns/op


### func IsTrue
BenchmarkBitSet_IsTrue/size_:_32768-4           2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_65536-4           2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_131072-4          2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_262144-4          2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_524288-4          2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_1048576-4         2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_2097152-4         2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_4194304-4         2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_8388608-4         2000000000               0.00 ns/op
BenchmarkBitSet_IsTrue/size_:_16777216-4        2000000000               0.01 ns/op
BenchmarkBitSet_IsTrue/size_:_33554432-4        2000000000               0.01 ns/op
BenchmarkBitSet_IsTrue/size_:_67108864-4        2000000000               0.02 ns/op
BenchmarkBitSet_IsTrue/size_:_134217728-4       1000000000               0.08 ns/op
BenchmarkBitSet_IsTrue/size_:_268435456-4       2000000000               0.09 ns/op
BenchmarkBitSet_IsTrue/size_:_536870912-4       2000000000               0.15 ns/op

PASS

