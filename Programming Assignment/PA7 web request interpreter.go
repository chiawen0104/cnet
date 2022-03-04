package main

import "fmt"
import "bufio"
import "net"
import "strings"
import "os"

func check(e error) {
   if e != nil {
      panic(e) 
   }
}

func main() {

   fmt.Println("Launching server...")
   ln, _ := net.Listen("tcp", ":12003") 
   defer ln.Close()
   
   for {
      conn, _ := ln.Accept()
      //defer conn.Close()

      reader := bufio.NewReader(conn)
      req, err := reader.ReadString('\n') 
      check(err)

      file := "" 
      tokens := strings.Split(req, " ") 
      for i := range tokens {
         //fmt.Println(tokens[i])
         if i == 1{
            file = strings.Replace(tokens[i], "/", "", -1)
            break
         }
      }

      f, err1 := os.Open(file) 
      if err1 != nil {
         fmt.Printf("File not found\n")
         conn.Close()
         continue
      }

      stats, err2 := os.Stat(file)
      check(err2)

      fmt.Printf("File size = %d\n", stats.Size())
      f.Close()
      conn.Close()
   } 
}
