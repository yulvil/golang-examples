package main

import "fmt"
import "reflect"

import (
   "encoding/json"
   "math/rand"
   "strconv"
   "time"
   "net/http"
   //"crypto"                     // Compile error if import not used
)

// Comment
/* Multi-line comment */


// ==========
// Types
// ==========

/*
string
bool
byte
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
rune
float32 float64
complex64 complex128
*/

func testTypes() {
   fmt.Println("=== TYPES ===")

   var i int = 2055
   var h byte = byte(i)           // Explicit cast required. Precision loss silently.
   fmt.Println(h)                 // 7

   // Strings are immutable
   a := "abc" + "123"             // String concat
   b := "def" +  strconv.FormatInt(456, 10)   // String concat
   c := string("xyz"[1])          // String concat
   d := append([]string{}, "abc") // Append preferable in a loop
   e := fmt.Sprintf("jhi%d", 789) // String concat
   f := `\r\n
         \u12e4`                  // Raw string. Multiline.
   fmt.Println(a,b,c,d,e)         // abc123 def456 y [abc] jhi789
   fmt.Println(f)                 // the raw string including all the spaces

   type MyInt int                 // Uppercase types are exported

   s := make([]int, 2)            // pointers to build-in structures are using make, not new
                                  // make's return type is the same as the first parameter
   t := new(MyInt)                // new returns a pointer to the type
   *t = 6
   _,_ = s,t
}

// ==========
// Numbers
// ==========


// ==========
// Variables
// ==========

func testVariables() {
   fmt.Println("=== VARIABLES ===")

   fmt.Println("Hello, 世界")     // UTF-8

   // Variables
   var x int = 5
   var y = 8                      // type optional
   var z bool;                    // Has default value.
   var (q = 5
        r = 6)                    // Multiple declarations
   var _ = x                      // Blank identifier will discard the value
   fmt.Println(x,y,z,q,r)         // 5 8 false 5 6

   var a,b,c int = 3,2,1          // Multiple assignments. Numbers must match.
   fmt.Println(a,b,c)             // 3 2 1

   var aa,bb,cc = 10, false, "ok" // type can be omitted
   fmt.Println(aa,bb,cc)          // 10 false ok

   xyz := 989                     // short syntax available inside functions
   fmt.Println(xyz)               // 989

   var zz =
             777                  // Statement on multiple line
   fmt.Println(zz)                // 777

   const PI = 3.1415              // Constant, cannot reassign

   fmt.Println(reflect.TypeOf(PI), reflect.TypeOf(zz)) // float64 int

   { var i=0; var j=1; i = i+j; } // i, j only accessible within block

   const (
      Monday = 1
      Tuesday = 2
      Wednesday = 3
   )

   const (
      Mon = iota                  // iota starts at zero
      Tue                         // incremented automatically
      Wed                         // within the same const ()
   )
   fmt.Println(Mon,Tue,Wed)       // 0 1 2

   type Letter rune               // Simulate enum
   const (
      A Letter = 97 + iota
      B
      C
   )
   fmt.Println(A,B,C)             // 97 98 99
}

// ==========
// Control statements
// ==========

func testControlStatements() {
   fmt.Println("=== CONTROL STATEMENTS ===")

   i := 0
   for i := 0; i < 10; i++ {}     // '{}' required, '()' not allowed
   for ; i > 10; {}               // init, increment optional
   for i > 10 {}                  // ';' optional
   for {break}                    // "infinite" loop

   if i < 0 {}                    // '{}' required, '()' not allowed
   if j:=0; j < 0 {}              // init statement before condition, scoped to the if/else block

   k := 1
   switch k=1; k {
      case 0: k = k + 3           // break implicit
      case 1: k = k + 5; fallthrough
      case 2: k = k * 2
      default:
   }
   fmt.Println(k)                 // 12

   switch {                       // switch true
      case k < 3: fmt.Println("small")
      case k > 9: fmt.Println("big")
      default: fmt.Println("medium")
   }
}


// ==========
// Struct
// ==========

type Point struct {
   x int   "my tag"
   y int
}

type Model struct {
   name string
}
func (m *Model) call() int {return 10}

type Road struct {
   Model                          // Embedded type. No name.
   loc string
}

func testStructs() {
   fmt.Println("=== STRUCTS ===")

   p := Point{1, 2}
   q := &Point{3, 4}                      // Point is initialized. Not just a pointer
   ptr := &p
   var i, j = p.x, q.x                    // Both Point and *Point accessed using '.' notation
   fmt.Println(p, p.y, ptr.x, q.x, i, j)  // {1, 2} 2 1 3 1 3

   s := Point{y: 5}                       // x:O implicit
   t := Point{}                           // x:O, y:0 implicit
   fmt.Println(s, t)                      // {0, 5} {0, 0}

   var u *Point = new(Point)              // Pointer to newly allocated Point
   v := new(Point)                        // same
   fmt.Println(u, v)                      // &{0, 0} &{0, 0}

   r := new(Road)
   fmt.Println(r.Model.call())            // Call embedded type using type name
   fmt.Println(r.call())                  // Omit embedded type

   field1, ok1 := reflect.TypeOf(p).FieldByName("x")
   fmt.Println(field1.Tag,ok1)            // my tag true
   field2, ok2 := reflect.TypeOf(&p).Elem().FieldByName("z")
   fmt.Println(field2.Tag,ok2)            //  false
}


// ==========
// Arrays
// ==========

// Arrays max size is uint64

func testArrays() {
   fmt.Println("=== ARRAYS ===")

   var arr [3]string              // Initialized with zero value for this type
   arr[0] = "A"
   arr[1] = string('X')
   arr[2] = "C"
   fmt.Println(arr, len(arr))     // [A X C] 3

   var arr2 = [...]int{9,8,7,6,5} // ... will calculate the array length
   fmt.Println(reflect.TypeOf(arr), reflect.TypeOf(arr2)) // [3]string [5]int
   arr3 := arr2                   // Array copy (full content)
   arr3[0] = 10                   // Only affects arr11
   fmt.Println(arr2, arr3)        // [9 8 7 6 5] [10 8 7 6 5]
}


// ==========
// Slices
// ==========

// Slices do not have element count
// A slice is a struct describing a section of an array

func testSlices() {
   fmt.Println("=== SLICES ===")

   var arr = [6]int{9,8,7,6,5,4}
   var s1 = arr[2:5]              // Slice an array
   var s10 = s1[1:2]              // Slice a slice
   var s11 = arr[:]               // Slice whole array
   fmt.Println(s1, s10, s11)      // [7 6 5] [6] [9 8 7 6 5 4]

   s2 := []int{10,20,30,40,50}    // No need to specify the size
   fmt.Println(s2)                // [10 20 30 40 50]
   s3 := s2[1:4]                  // Same underlying array as s2. Not a copy.
   s4 := s2[2:]                   // Slice from 2 to len()
   s5 := s2[:3]                   // Slice from 0 to 3
   fmt.Println(s3, s4, s5)        // [20 30 40] [30 40 50] [10 20 30]

   s6 := make([]int, 2)           // zero initialized
   s6 = append(s6, 1, 2, 3)
   s6 = append(s6, []int{4,5}...) // Expand slice and append
   fmt.Println(s6)                // [0 0 1 2 3 4 5]

   s7 := s2                       // Not slice copy
   s7[0] = 11                     // Modifies the arrays backing both slices
   fmt.Println(s2, s7)            // [11 20 30 40 50] [11 20 30 40 50]
}


// ==========
// Ranges
// ==========

// Iterate over string, array, slice, map, channel

func testRanges() {
   fmt.Println("=== RANGES ===")

   var arr = [5]int{9,8,7,6,5}
   for i, v := range arr {        // i:index, arr[i]
      fmt.Println(i, v)
   }
   for _, v := range arr {        // Ignore index
      fmt.Println(v)
   }
   for i := range arr {           // Ignore value
      fmt.Println(i)
   }
   for _, c := range "abc123" {
      fmt.Printf("%c\n", c)
   }
}

// ==========
// Maps
// ==========

func testMaps() {
   fmt.Println("=== MAPS ===")

   m := make(map[string]string)
   m["key1"] = "value1"
   m["key2"] = "value2"
   fmt.Println(m, m["key2"])      // map[key2:value2 key1:value1] value2
   for k,v := range m {           // iterate over keys
      fmt.Println(k,v)
   }

   m2 := map[string]int { "key1": 10, "key2": 20, }
   fmt.Println(m2, m2["key2"])    // map[key1:10 key2:20] 20

   m3 := map[string]Point {
      "key1": {1,2},              // No need to specify Point{1,2}
      "key2": {3,4},              // Extra comma is ok
   }
   fmt.Println(m3, m3["key2"].y)  // map[key1:{1 2} key2:{3 4}] 4
   m3["key3"] = Point{5, 6}       // Point required here
   p := m3["key2"]
   fmt.Println(p)                 // {3,4}
   value, contains := m3["key9"]  // zero value for type Point
   fmt.Println(value, contains)   // {0,0} false

   m4 := map[string]int{}         // same as using make
   delete(m4, "key8")             // silent, no return values
}


// ==========
// Functions
// ==========

// Can be called before being declared
// Nested functions not allowed
// Everything is pass by value

func add(x int, y int) int {
   return x + y
}

func sub(x int, y int) int {
   return x - y
}

func swap(x, y string) (string, string) {  // return multiple results
   return y, x
}

func adder(i int) func(int) int {   // Closure
   return func(x int) int {
      return i + x
   }
}

type binFunc func(int, int) int

func leak() *int {
   var i int = 88
   return &i                      // local variable will be accessible to caller
}                                 // as long as it keeps a reference to it

func callMe(arg int) (result string, err int) {
   result, err = "", 0            // default values for named result parameters
   if arg > 10 {
      result, err = "too big", -1
   }
   return                         // implicit result parameters
}

func testFunctions() {
   fmt.Println("=== FUNCTIONS ===")

   fmt.Println(add(6,7))
   fmt.Println(swap("abc", "def"))

   f1 := add
   fmt.Println(f1(7,8))           // 15

   f2 := func(x, y int) Point {   // Assign anonymous function
      return Point{x, y}
   }
   fmt.Println(f2(3,2))           // {3,2}

   add2 := adder(2)
   add3 := adder(3)
   fmt.Println(add2(5), add3(6))  // 7 9
   //add4 := mkAdd(4)

   var f3 binFunc
   f3 = binFunc(add)              // "cast" add as a binFunc
   fmt.Println(f3(5,6))           // 11
   f3 = binFunc(sub)              // "cast" sub as a binFunc
   fmt.Println(f3(5,6))           // -1

   fmt.Println(*leak())           // 88
   fmt.Println(callMe(11))        // 88
}


// ==========
// Methods
// ==========

type State struct { id, code string }

func (s *State) toString() string { // State is the receiver
   return s.code + " " + s.id
}

func (s *State) setCode(code string) *State {
   s.code = code                  // Receiver needs to be a pointer otherwise
   return s                       // we will modify a copy (pass by value)
}

type MyByte byte
func (b MyByte) reverse() MyByte {  // Method on value type
   return ^b
}

func testMethods() {
   fmt.Println("=== METHODS ===")

   fmt.Println((&State{"25", "CA"}).toString()) // CA 25
   fmt.Println((&State{"25", "CA"}).setCode("CB")) // &{25 CB}
   fmt.Println(MyByte(7).reverse()) // 248
}


// ==========
// Interfaces
// ==========

// 'implements' implicit. No way/need to specify.

type Processor struct {}
func (p *Processor) eval() bool { return true }

type Calculator struct {}
func (c *Calculator) eval() bool { return false }

type Evaluable interface {
   eval() bool
}
func process(e Evaluable) bool {return e.eval()}

func testInterfaces() {
   fmt.Println("=== INTERFACES ===")

   fmt.Println(process(&Processor{}))   // true
   fmt.Println(process(&Calculator{}))  // false

   // Check if value implements Evaluable
   var isEvaluable = func (any interface{}) bool {
      if _, ok := any.(Evaluable); ok {
         return true
      }
      return false
   }
   fmt.Println(isEvaluable(2),            // false
               isEvaluable(&Processor{}), // true
               isEvaluable(&Processor{}), // true
               isEvaluable(&Point{}))     // false
}


// ==========
// Goroutines
// ==========

func testGoroutines() {
   fmt.Println("=== GOROUTINES ===")

   f := func(i int) {
      fmt.Println("go ", i)
   }
   for i:=0; i<5; i++ {
      go f(i)                     // Not necessarily in order
   }
}


// ==========
// Channels
// ==========

func testChannels() {
   fmt.Println("=== CHANNELS ===")

   testChannels1()
   testChannels2()
}

func testChannels1() {
   // Generate n random ints
   produce := func(c chan int, n int) {
      defer close(c)
      for i:=0; i<n; i++ {
         c <- rand.Intn(10)
      }
   }

   c := make(chan int)
   go produce(c, 5)
   for i := range c {
      fmt.Println("Recv: ", i)
   }

   // use channel to simulate python generators
   fib := func (n int) chan int {
      c := make(chan int)
      go func() {
         x, y := 0, 1
         for i := 0; i < n; i++ {
            c <- x
            x, y = y, x+y
         }
         close(c)
      }()
      return c
   }
   for i := range fib(10) {
      fmt.Println(i)
   }

   ch := make(chan int)
   stop_ch := make(chan bool)
   go func(src chan int) {
      for {
         select {
            case src <- rand.Intn(100):
            case <- stop_ch: {src = nil}
         }
      }
   }(ch)
   fmt.Println(<-ch, <-ch, <-ch)  // Read from channel
   stop_ch <- true
   //fmt.Println(<-ch)            // Channel is closed now

   time.Sleep(1000 * time.Millisecond)    // Let the goroutines finish
}

func testChannels2() {
   type work struct {
      url string                  // Request
      resp chan *http.Response    // Write response here
   }

   worker := func (q chan work) {
      for {
         item := <- q
         resp, _ := http.Get(item.url)
         item.resp <- resp        // Write to response channel
      }
   }

   q := make(chan work)
   go worker(q)

   resp_ch := make(chan *http.Response)
   q <- work{"http://www.google.com", resp_ch}
   fmt.Println(<- resp_ch)

   time.Sleep(1000 * time.Millisecond)    // Let the goroutines finish
}

// ==========
// Defer
// ==========

func testDefer() {
   fmt.Println("=== DEFER ===")

   testDefer1()
   testDefer2()
}

func testDefer1() {
   defer fmt.Println("clean up")  // Called immediately before this func returns
   fmt.Println("make a mess")
}

func testDefer2() {
   i := 5
   defer func () {
      fmt.Println(i)              // i will be evaluated right before the 'parent' func returns
   }()                            // defer will print 100
   i = 100
}


// ==========
// Json
// ==========

type Person struct {   // Annotate fields with tags. Available through reflection.
   FirstName  string   `json:"first_name"`
   LastName   string   `json:"last_name"`
   MiddleName string   `json:"middle_name,omitempty"`
}
func (p *Person) String() string {     // Change default "toString"
   return fmt.Sprintf("[%s] [%s] [%s]", p.FirstName, p.MiddleName, p.LastName)
}

func testJson() {
   fmt.Println("=== JSON ===")

   json_string := `{"first_name": "John", "last_name": "Smith"}`
   person := new(Person)
   json.Unmarshal([]byte(json_string), person)
   fmt.Println(person)            // [John] [] [Smith]

   p := Person{FirstName:"Jane", LastName:"Doe"}
   s, err := json.Marshal(p)
   fmt.Println(err, string(s))    // <nil> {"first_name":"Jane","last_name":"Doe"}
}


// ==========
// Misc
// ==========

func testMisc() {
   fmt.Println("=== MISC ===")

   time.Sleep(5 * time.Millisecond)

   // panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}


func main() {
   testTypes()
   testVariables()
   testControlStatements()
   testStructs()
   testArrays()
   testSlices()
   testRanges()
   testMaps()
   testFunctions()
   testMethods()
   testInterfaces()
   testChannels()
   testDefer()
   testJson()
   testMisc()
   testGoroutines()
   time.Sleep(200 * time.Millisecond)    // Let the goroutines finish
   // return                             // Explicit return statement for main
}
