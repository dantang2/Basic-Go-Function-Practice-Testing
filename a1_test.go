/*  a1_test.go
**  Dan Tang - 301256468
**  CMPT 383 - Assignment 1
**  Last modified: January 24, 2018
*/


package a1

import (

  "fmt"
  "testing"
  "reflect"
  "errors"

)


func TestCountPrimes(t *testing.T) {
  fmt.Println("\n---- STARTING TEST FOR countPrimes() ----\n")

  fmt.Println("Testing for negative integers: ")
  number:= countPrimes(-1)
  anotherNumber := countPrimes (-6)
  if number !=0 || anotherNumber != 0{
    t.Errorf("Expected value was 0")
  }

  fmt.Println("\nTesting for correct answer: ")
  number = countPrimes(10000)
  if number != 1229 {
    t.Errorf("Expected value was 1229")
  }

  number = countPrimes(5)
  if number != 3 {
    t.Errorf("Expected value was 3")
  }

    fmt.Println("\n---- ENDING TEST FOR countPrimes() ----\n")
}


func TestCountStrings(t *testing.T) {
  fmt.Println("---- STARTING TEST FOR countStrings() ----\n")

  fmt.Println("Testing for empty file:")
  words := countStrings("empty.txt")
  if len(words) != 0 {
    t.Errorf("Map should be empty")
  }


  fmt.Println("\nTesting for correct answer: ")
  words = countStrings("sample.txt")
  var m = map[string]int {
    "The": 1,
    "big": 3,
    "dog": 1,
    "ate": 1,
    "the": 1,
    "apple": 1,
   }

   fmt.Println(m)

  equal := reflect.DeepEqual(m, words)
  if !equal {
    t.Errorf("Expecting correct answer")
  }


  fmt.Println("\n---- ENDING TEST FOR countStrings() ----\n")

}


func TestTime24(t *testing.T) {

  fmt.Println("---- STARTING TEST FOR Time24 functions ----\n")

  time1 := Time24{hour: 23, minute: 39, second: 8}
  time2 := Time24{hour: 5, minute: 39, second: 8}

  fmt.Println("Testing for non-equal times: ")

  equal := equalsTime24(time1, time2)
  if equal {
    fmt.Println("error?")
    t.Errorf("Expected non equal times")
  }

  time1 = Time24{hour: 23, minute: 39, second: 8}
  time2 = Time24{hour: 23, minute: 39, second: 8}

  fmt.Println("\nTesting for equal times: ")

  equal = equalsTime24(time1, time2)
  if !equal {
    fmt.Println("error?")
    t.Errorf("Expected equal times")
  }

  fmt.Println("\nTesting lessThanTime24(): ")
  time1 = Time24{hour: 15, minute: 20, second: 7}
  time2 = Time24{hour: 15, minute: 39, second: 8}

  lessThan := lessThanTime24(time1, time2)
  if !lessThan {
    t.Errorf("Expected time 1 to be less than time 2")
  } else {
    fmt.Printf("Time1: %v is less than Time2: %v\n",time1, time2)
  }

  time1 = Time24{hour: 18, minute: 20, second: 7}
  lessThan = lessThanTime24(time1, time2)
  if lessThan {
    t.Errorf("Expected time 1 to be greater than or equal to time 2")
  }

  fmt.Println("\nTesting Time24 to string conversion: ")
  time1 = Time24{hour: 18, minute: 20, second: 7}

  timeString := time1.String()
  fmt.Printf("Time1: {hour: 18, minute: 20, second: 7} = %v\n",time1)
  if timeString != "18:20:07"{
    t.Errorf("Expected correct conversion")
  }

  fmt.Println("\nTesting ValidTime(): ")

  time1 = Time24{hour: 18, minute: 20, second: 7}
  if !time1.ValidTime() {
    t.Errorf("Expected a valid time.")
  } else {
    fmt.Printf("Time1: %v is valid.\n",time1)
  }

  fmt.Println("\nTesting for non-valid time:")
  time2 = Time24{hour: 25, minute: 20, second: 7}
  if time2.ValidTime(){
    t.Errorf("Expected a non-valid time.")
  }

  fmt.Println("\nTesting minTime24(): ")
  times := make([]Time24, 4)


  times[0] = Time24{23,59,59}
  times[1] = Time24{hour:15, minute:32, second:05 }
  times[2] = Time24{hour:20, minute:45, second:24 }
  times[3] = Time24{hour:16, minute:39, second:55 }

  minTime, err := minTime24(times)

  if err != nil {
    t.Errorf("Expecting no errors")
  }

  if minTime != times[1]{
    fmt.Printf("Test failed: times[1]: %v supposed to be minTime\n",times[1])
    t.Errorf("Expecting minimum time")
  } else {
    fmt.Printf("Minimum time is %v\n",minTime)
  }

  fmt.Println("\nTesting for empty []Time24 slice: ")
  var times2 []Time24

  minTime, err = minTime24(times2)

  emptyStringError := errors.New("ERROR: Your []Time24 slice is empty.\n")

  if err.Error() == emptyStringError.Error() {
    fmt.Println(err)
  } else {
    t.Errorf("Expecting empty string error.")
  }


    fmt.Println("\n---- ENDING TEST FOR Time24() ----\n")
}


func TestLinearSearch(t *testing.T) {

    fmt.Println("--- STARTING TEST FOR LinearSearch() ----\n")

    fmt.Println("Testing for string search: ")
    value := "cat"
    pos,err := linearSearch(value, []string{"4","2","cat","5","0"})

    if err != nil {
      t.Errorf("Expected to find 'cat'")
    } else {
      fmt.Printf("%v is at position: %v in the list\n", value, pos)
    }

    fmt.Println("\nTesting for integer search:")
    number := 59
    pos, err = linearSearch(number, []int{10,9,8,7,59,23,4})

    if err != nil {
      t.Errorf("Expected to find 59")
    } else {
      fmt.Printf("%v is at position: %v in the list\n", number, pos)
    }

    fmt.Println("\nTesting for non-matching data types:")

    pos, err = linearSearch(value, []int{10,9,8,7,59,23,4})

    if pos == 404 && err != nil {
      fmt.Println(err)
    } else {
     t.Errorf("Expected non-matching data types")
    }

    fmt.Println("\n---- ENDING TEST FOR LinearSearch() ----\n")
}


func TestAllGeneSeqs(t *testing.T){

  fmt.Println("--- STARTING TEST FOR allGeneSeqs() ----\n")
  fmt.Println("Testing for empty string: ")

  geneSequence := allGeneSeqs(0)
  var emptyString []string

  if !reflect.DeepEqual(geneSequence,emptyString){
    t.Errorf("Expecting empty string.")
  } else {
    fmt.Println(geneSequence)
  }

  fmt.Println("\nTesting for correct gene sequences: ")
  geneSequence = allGeneSeqs(4)
  for i := 0; i < len(geneSequence); i++ {
    for j := 0; j < len(geneSequence); j++ {
      if geneSequence[i] == geneSequence[j] && i != j {
        t.Errorf("Not the correct sequence")
      }
    }
  }



  fmt.Println("\n--- ENDING TEST FOR allGeneSeqs() ----\n")
}
