# Description
Has 2 instructions:

```powershell
$ db_set key value

$ db_get key

$ db_del key
```

Setting a value for a key appends the key value pair in a text file

Reading a value requires reading from the end to the beginning of the file to get the latest value for the key. When reading a key’s value, check that it is not the tombstone value.

Deleting a value happens by adding a tombstone value to the key such as `status: deleted` 

If a user sets the value manually to this tombstone value, it will be treated as deleted.

Best file format is a binary format that encodes the length of a string in bytes followed by the raw string.

Allows only 1 writer and multiple readers.

# Improvements
- Read file in chunks instead of reading all of it to fit into memory
- Make chunk sizes the biggest to fit into memory to reduce the number of I/O's
- Parsing strings ('' => gets parsed to empty string for example and commas sometimes are stored as empty strings as well)
- Every key should be mapped to a byte offset in the data file with a file containing keys mapping to byte offset in another file containing values
