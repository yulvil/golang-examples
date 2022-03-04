package main                                          
                                                      
/*                                                    
#cgo LDFLAGS: -L${SRCDIR} -lm mylib.a                 
#include <math.h>                                     
#include "mylib.h"                                    
int mysub(int n, int m) { return n-m; }               
typedef enum {RANDOM=100, IMMEDIATE, SEARCH} strategy;
const double PI = 3.1415;                             
typedef struct { int x, y, type; } mystruct;          
*/                                                    
import "C"                                            
                                                      
import "fmt"                                          
                                                      
func zz() (int, float64) {                            
        return 1, 2.2                                 
}                                                     
                                                      
func main() {                                         
        fmt.Println(C.myfunc())                       
        fmt.Println(C.myadd(100, 200))                
        fmt.Println(C.mysub(100, 200))                
        fmt.Println(C.exp(1.23))                      
        fmt.Println(C.IMMEDIATE)                      
        fmt.Println(C.PI)                             
        var s = C.mystruct{10, 20, 30}                
        fmt.Println(s, s.x, s.y, s._type)             
        fmt.Println(C.struct_mystruct{})              
        fmt.Println(C.longlong(98765))                
        fmt.Println(C.uchar('a'))                     
                                                      
        //var res, err = C.sqrt(1)                    
        res, err := zz()                              
        fmt.Println(res, err)                         
}                                                     
