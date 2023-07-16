# Description
- Does the same as in 1.0 with a hash table storing keys with values byte offset in the data file. 
- If the DBMS crashes and restarts, it regenerates the Hash Table first then take instructions.
- When a key is added to the file, it should be updated in the hash map.
- *This approach is used in **Bitcask** (**Riak** storage engine)

# Improvements
- Same Improvements as previous version
- Reuse the repeated code from v1.1

# How to run the tests and benchmarks?
- Go inside modules
- Run `go test` to run the tests
- Run `go test -bench=.` to run the benchmarks

# Comparison example:
- BenchmarkDBGetWithHashMap-4     26.68     ns/op
- BenchmarkDBGetNoHashMap-4       158380496 ns/op

*This shows significant improvement when the HashMap is being used*
