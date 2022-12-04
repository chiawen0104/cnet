# Introduction
Programming assignments(coded by GO) of Introduction to Computer Networks, NTUEE, 2021 fall.

## Programming Assignments
### PA1: unix command
basic unix command practice (`ls`, `mv`, `mkdir`, `rm`, etc.)

### PA2: file access
1. prompts the user for the input and output filenames 
2. reads from the input file one line at a time
3. prepends the line count to each line
4. writes the line into the output file

### PA3: upload client
1. connects to the server that polly implements and runs already on the workstation at port 12000
2. prompts the user for the upload filename
3. sends first the file size (just the number in a single line)
4. sends next the file content (the entire file)
5. receives a message back from the server
6. prints what the server says
7. closes the connection and terminates the program

### PA4: upload server
1. listens at <your port#> until there’s an upload request
2. reads from the socket first the file size (just the number in a single line)
3. reads from the socket one line at a time
4. prepend the line count to each line and store the new line into a new file: whatever.txt
5. repeats (3) and (4) until all lines in the file is processed
6. sends a message back that tells the client the original file and the new file size
7. closes the connections and terminates the program

### PA5: looping server
1. listens at <your port#> until there’s an upload request
2. reads from the socket first the file size (just the number in a single line)
3. reads from the socket one line at a time
4. prepend the line count to each line and store the new line into a new file: whatever.txt
5. repeats (3) and (4) until all lines in the file is processed
6. sends a message back that tells the client the original file and the new file size
7. closes the connection and goes back to (1)

### PA6: concurrent server
1. Listens at <your port#> until there’s an upload request
2. reads from the socket first the file size (just the number in a single line)
3. reads from the socket one line at a time
4. prepend the line count to each line and store the new line into a new file: whatever.txt
5. repeats (3) and (4) until all lines in the file is processed
6. sends a message back that tells the client the original file and the new file size
7. closes the connections and goes back to (1)

### PA7: web request interpreter
1. Listens at <your port#> until there’s an HTTP request
2. reads from the socket
3. finds the path and name of the text/html file requested
4. In case the file exists, prints on server screen the file size
5. In case the file doesn’t exist, prints on server screen “File not found”
6. closes the connections and goes back to (1)

