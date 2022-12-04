package main

import "fmt"
import "bufio"
import "net"
import "os"
import "strings" 
import "strconv"

func check(e error) {
   	if e != nil {
		panic(e) 
	}
}

func main() {

   	fmt.Println("Launching server...")
   	ln, _ := net.Listen("tcp", ":12003") 
   	conn, _ := ln.Accept()
   	defer ln.Close()
   	defer conn.Close()
	
   	//read file size from socket(string)
   	reader1 := bufio.NewReader(conn) 
   	m_size, err1 := reader1.ReadString('\n')  
   	check(err1)
	fmt.Printf("Upload file size: %s", m_size)

   	//Remove '\n' in string, then convert string to int
	m_size = strings.Replace(m_size, "\n", "", -1)
   	size_int, err2 := strconv.Atoi(m_size)
	check(err2)
	//fmt.Print(size_int)

   	//create new file
   	out_file := "whatever.txt"
	newfile, err3 := os.Create(out_file) 
   	check(err3)
   	defer newfile.Close()


   	//read and write line to new file
   	writer1 := bufio.NewWriter(newfile)
   	i := 1
   	tmp_size := 0
   	length := 0

   	for {
       	line, errr := reader1.ReadString('\n')
       	check(errr)

		writer1.WriteString(fmt.Sprintf("%d %s", i, line))
		length = len(line) //length of each line(string)
		tmp_size += length
      		i += 1
	
		if (tmp_size >= size_int) {
	     		break
		}
	}
   	writer1.Flush()
   
   	// get the new file's size 
   	stats, err4 := os.Stat(out_file)
   	check(err4)
   	fmt.Printf("Output file size: %d\n", stats.Size())

   	// send message back to client
   	writer2 := bufio.NewWriter(conn)
   	newline := fmt.Sprintf("%s bytes received, %d bytes file generated",m_size,stats.Size()) 
   	_, errw := writer2.WriteString(newline)
   	check(errw)
   	writer2.Flush()
}
