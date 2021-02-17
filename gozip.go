package main 
  
// Declare all the libraries which are needed. 
import ( 
    "compress/gzip"
    "fmt"
    "os"
    "bufio"
    "io/ioutil"
) 
  
func main() { 
  fi, _ := os.Open(os.Args[1]) 
  
  read := bufio.NewReader(fi)
  
  data, _ := ioutil.ReadAll(read) 
  
  fo, _ := os.Create(os.Args[1]+".gz") 
  
  p := gzip.NewWriter(fo) 
  
  p.Write(data)
  
  p.Close() 
  
  fmt.Println("Done") 
}
