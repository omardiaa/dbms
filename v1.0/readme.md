# Improvements
- Read file in chunks instead of reading all of it to fit into memory
- Make chunk sizes the biggest to fit into memory to reduce the number of I/O's
- Parsing strings ('' => gets parsed to empty string for example and commas sometimes are stored as empty strings as well)
