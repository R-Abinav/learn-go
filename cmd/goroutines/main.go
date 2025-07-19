package main

// --- CODE SNIPPET 1 ---
//No concurrency, synchronous

// import (
// 	"fmt"
// 	"time"
// 	"math/rand"
// )

// var dbData []string = []string{"id1", "id2", "id3", "id4", "id5"}
// var results []string = []string{}

// func main(){
// 	t0 := time.Now()
// 	for i := 0; i<len(dbData); i++{
// 		dbCall(i)
// 	}

// 	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
// 	fmt.Printf("\nThe result from db: %v", results)
// }

// func dbCall(i int){
// 	//Simulate db call delay
// 	var delay float32 = rand.Float32() * 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)

// 	fmt.Println("The result from the database is: ", dbData[i])
// 	results = append(results, dbData[i])
// }

// -- OUTPUT OF ABOVE -- 
/* 
The result from the database is:  id1
The result from the database is:  id2
The result from the database is:  id3
The result from the database is:  id4
The result from the database is:  id5

Total execution time: 5.868822167s
The result from db: [id1 id2 id3 id4 id5]
*/

// --- CODE SNIPPET 2 ---
//Concurrency, with a fixed delay but no results array

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var dbData []string = []string{"id1", "id2", "id3", "id4", "id5"}

// func main(){
// 	t0 := time.Now()
// 	for i := 0; i<len(dbData); i++{
// 		wg.Add(1)
// 		go dbCall(i)
// 	}
// 	wg.Wait()

// 	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
// }

// func dbCall(i int){
// 	//Simulate db call delay
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)

// 	fmt.Println("The result from the database is: ", dbData[i])

// 	wg.Done()
// }

// --- OUTPUT OF THE ABOVE ---
/*
The result from the database is:  id4
The result from the database is:  id2
The result from the database is:  id3
The result from the database is:  id1
The result from the database is:  id5

Total execution time: 2.001448166s
*/



// --- CODE SNIPPET 3 ---
//This code snippet appends to results and displays results.
//This (most of the time, as delay is randomised) works without mutexes as the dbCall isn't a fixed delay!

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// 	"math/rand"
// )

// var wg = sync.WaitGroup{}
// var dbData []string = []string{"id1", "id2", "id3", "id4", "id5"}
// var results []string = []string{}

// func main(){
// 	t0 := time.Now()
// 	for i := 0; i<len(dbData); i++{
// 		wg.Add(1)
// 		go dbCall(i)
// 	}
// 	wg.Wait()

// 	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
// 	fmt.Printf("\nThe result from db: %v", results)
// }

// func dbCall(i int){
// 	//Simulate db call delay
// 	var delay float32 = rand.Float32() * 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)

// 	fmt.Println("The result from the database is: ", dbData[i])
// 	results = append(results, dbData[i])

// 	wg.Done()
// }

// --- OUTPUT OF THE ABOVE ---
/* 
The result from the database is:  id4
The result from the database is:  id1
The result from the database is:  id5
The result from the database is:  id3
The result from the database is:  id2

Total execution time: 1.956501666s
The result from db: [id4 id1 id5 id3 id2
*/

// --- CODE SNIPPET 4 ---
//Fixed delay and results array -> Notice whats happens when a slice is modified at the same time by multiple go routines
//answer -> Now you have multiple threads modifying the same memory location at the same time, you can get unexpected results -> Like missing data
//For example -> Two processes writing to the same memory location at the same time, could lead to corrupt memory

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var dbData []string = []string{"id1", "id2", "id3", "id4", "id5"}
// var results []string = []string{}

// func main(){
// 	t0 := time.Now()
// 	for i := 0; i<len(dbData); i++{
// 		wg.Add(1)
// 		go dbCall(i)
// 	}
// 	wg.Wait()

// 	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
// 	fmt.Printf("\nThe result from db: %v", results)
// }

// func dbCall(i int){
// 	//Simulate db call delay
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)

// 	fmt.Println("The result from the database is: ", dbData[i])
// 	results = append(results, dbData[i])

// 	wg.Done()
// }

// --- OUTPUT FROM THE ABOVE ---
/* 
The result from the database is:  id1
The result from the database is:  id2
The result from the database is:  id5
The result from the database is:  id4
The result from the database is:  id3

Total execution time: 2.001725958s
The result from db: [id1 id2 id4 id3]    <- Missing Data
*/



// --- CODE SNIPPET 5 ---
//Solution -> Mutexes, to control the writing to our slice in a way that makes it safe in a concurrent program

//A go routine reaches the place where lock code is return, it checks if lock is present (Set by another go routine) or not
//If not present, it continues to perform the task and set the Lock.
//If present, it will wait there till the task gets unlocked (By another go routine) and then proceeds to perform the task and set the Lock

//NOTE -> IT is VERY VERY IMP, on where you place your locks and unlocks

//One drawback of this sort of mutex, is that it completely locks out the other go routines to accessing the result slice.

import (
	"fmt"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData []string = []string{"id1", "id2", "id3", "id4", "id5"}
var results []string = []string{}

func main(){
	t0 := time.Now()
	for i := 0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe result from db: %v", results)
}

func dbCall(i int){
	//Simulate db call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()

	wg.Done()
}

// --- OUTPUT OF THE ABOVE ---
/* 
The result from the database is:  id4
The result from the database is:  id5
The result from the database is:  id3
The result from the database is:  id1
The result from the database is:  id2

Total execution time: 2.001743791s
The result from db: [id4 id5 id3 id1 id2]
*/