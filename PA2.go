package main

import "fmt"
import "os"
import "bufio"

func check(e error) {
	if e != nil {
		panic(e) 
	}
}

func main () {

   	fmt.Printf("Input filename: ")
   	in_file := ""
   	fmt.Scanf("%s", &in_file)

   	f1, err := os.Open(in_file) 
   	check(err)
	
   	fmt.Printf("Output filename: ")
   	out_file := ""
   	fmt.Scanf("%s", &out_file)

	f2, err := os.Create(out_file) 
   	check(err)
   	defer f2.Close()

	writer := bufio.NewWriter(f2)
	scanner := bufio.NewScanner(f1)
	if scanner.Scan() {
		writer.WriteString(fmt.Sprintln(scanner.Text())) 
  	}
   		
	f1.Close()
	writer.Flush()
	  	
} 

