package main

import "fmt"
import "log"
import "os"
import "bufio"
import "net"
import "strconv"

func check(e error) {
   if e != nil {
      panic(e) 
   }
}

func main () {

   conn, errc := net.Dial("tcp", "127.0.0.1:12000") 
   check(errc)
   defer conn.Close()

   fmt.Printf("Input filename: ")
   in_file := ""
   fmt.Scanf("%s", &in_file)

   f, err := os.Open(in_file) 
   check(err)
   
   stats, err := os.Stat(in_file)
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("Send the file size first: %d\n", stats.Size())
   
   writer := bufio.NewWriter(conn)
   _, errw := writer.WriteString(strconv.Itoa(int(stats.Size())) + "\n")
   check(errw)
   scanner1 := bufio.NewScanner(f)
   for scanner1.Scan() {
      writer.WriteString(fmt.Sprintln(scanner1.Text()))
   }
   writer.Flush()

   scanner2 := bufio.NewScanner(conn)
   if scanner2.Scan() {
      fmt.Printf("Server says: %s\n", scanner2.Text()) 
   }

   f.Close()
} 