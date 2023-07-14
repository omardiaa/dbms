# Description
- Does the same as in 1.0 with a hash table storing keys with values byte offset in the data file. 
- If the DBMS crashes and restarts, it regenerates the Hash Table first then take instructions.
- When a key is added to the file, it should be updated in the hash map.
- *This approach is used in **Bitcask** (**Riak** storage engine)

# Improvements
- Same Improvements as previous version
