# Description
- Does the same as in 1.0 with a hash table storing keys with values byte offset in the data file. 
- If the DBMS crashes and restarts, it regenerates the Hash Table first then take instructions.
- When a key is added to the file, it should be updated in the hash map.
- *This approach is used in **Bitcask** (**Riak** storage engine)
## Concurrency Controler
- When multiple threads write concurrently to the database, they access the same HashMap and modify the values based on the offset of the new key in the file. 
- Example:
  - Thread 1 opens the file 
  - Thread 2 opens the file
  - Thread 1 reads the current fileSize (fs0)
  - Thread 2 reads the current fileSize (fs1) => In this case fs0 = fs1
  - Thread 1 writes and stores the offset in the HashMap with fs0
  - Thread 2 writes and stores the offset in the HashMap with fs1
  - When reading the thread 2's value, it reads the value thread 1 wrote
- Solution:
  - Use a lock before opening the file
- If multiple threads modify the same key, Last Write Wins (LWW) approach is used to decide which is the most up to date value which in this case is the last written key-value pair

# Improvements
- Same Improvements as previous version
- Reuse the repeated code from v1.1

# How to run the tests and benchmarks?
- Go inside modules
- Run `go test` to run the tests
- Run `go test -bench=.` to run the benchmarks

# Comparison example:
- BenchmarkDBGetWithHashMap-4          170           7063325 ns/op
- BenchmarkDBGetNoHashMap-4              5         229147551 ns/op

*This shows significant improvement when the HashMap is being used*
