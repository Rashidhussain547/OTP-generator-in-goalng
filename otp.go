package main

import ( 
    "fmt"
    "math/rand"
    "time"
)        

var randowCodes = [...]byte{
    '1', '2', '3', '4', '5', '6', '7', '8', '9', '0','Q','W','E','R','T','Y','U','I','O','P','A','S','D','F','G','H','J','K','L','Z','X','C','V','B','N','M',     
}        

func main() {
    var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < 3; i++ {
        var pswrd []byte = make([]byte, 6)

        for j := 0; j < 6; j++ {
            index := r.Int() % len(randowCodes)
            pswrd[j] = randowCodes[index]
        }

        fmt.Printf("%s\n", string(pswrd))                                                                  
    }    
} 