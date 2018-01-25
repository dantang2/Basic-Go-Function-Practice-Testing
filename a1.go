/*  a1.go
**  Dan Tang - 301256468
**  CMPT 383 - Assignment 1
**  Last modified: January 24, 2018
*/


package a1

import (
  "fmt"
  "math"
  "io/ioutil"
  "strings"
  "reflect"
  "errors"


)

func countPrimes(value int) int{

  numberOfPrimes := 0;

  if value < 0{
    fmt.Println("ERROR: Negative integer was entered.")
    return 0
  }

  f :=make([]bool, value+1)

  // Use Sieve Of Eratosthenes' algorithm
  for i:=2; i<=int(math.Sqrt(float64(value))); i++ {
    if f[i] == false {
      for j := i * i; j <= value; j+=i {
        f[j] = true
      }
    }
  }

  for i:=2; i<=value; i++ {
    if f[i] == false {
      numberOfPrimes++
    }
  }
  fmt.Printf("The number of prime numbers less than or equal to %d is %d\n", value, numberOfPrimes)

  return numberOfPrimes


}

func countStrings(file_name string) map[string]int {

  file, err := ioutil.ReadFile(file_name)
  word_count :=make(map[string]int)

  if err != nil{
    fmt.Println(err)
    return word_count
  }


  s := string(file)
  words := strings.Fields(s)

  fmt.Println("Text file: ",words)



  for i := 0; i < len(words); i++ {
    aWord := words[i]
    word_count[aWord]++
  }

  return word_count


}

type Time24 struct {

  hour, minute, second uint8

}

func equalsTime24(a, b Time24) bool{

  if !a.ValidTime() {
    fmt.Printf("ERROR: Time A: %v is not a valid time\n",a)
    return false
  } else if !b.ValidTime() {
    fmt.Printf("ERROR: Time B: %v is not a valid time\n",b)
    return false
  }

  equal := false

  if a==b {
    equal = true
    fmt.Printf("Time A: %v and Time B: %v are equal\n",a, b)
  } else {
    fmt.Printf("Time A: %v and Time B: %v are not equal\n",a, b)
  }

  return equal

}

func lessThanTime24(a, b Time24) bool {

    if !a.ValidTime() {
      fmt.Printf("ERROR: Time A: %v is not a valid time\n",a)
      return false
    } else if !b.ValidTime() {
      fmt.Printf("ERROR: Time B: %v is not a valid time\n",b)
      return false
    }

  lessThan := false

  if a.hour < b.hour {
    lessThan = true
  }

  if a.hour == b.hour {
    if a.minute < b.minute {
      lessThan = true
    }

    if a.minute == b.minute {
      if a.second < b.second {
        lessThan = true
      }
    }
  }

  /*if lessThan {
    fmt.Printf("Time A: %v is less than Time B: %v\n",a,b)
  } else {
    fmt.Printf("Time A: %v is not less than Time B: %v\n",a,b)
  }*/

  return lessThan

}

func (t Time24) String() string {

  if !t.ValidTime(){
    fmt.Printf("This time: %v:%v:%v is not a valid time\n",t.hour, t.minute, t.second)
  }

  var thisHour string = fmt.Sprint(t.hour)
  var thisMinute string = fmt.Sprint(t.minute)
  var thisSecond string = fmt.Sprint(t.second)





  if t.hour < 10 {
  thisHour = "0" + thisHour
  //thisHour = "0" + thisHour
  }


  if t.minute < 10 {
    thisMinute = "0" + thisMinute
    //thisMinute = "0" + thisMinute
  }

  if t.second < 10 {
    thisSecond = "0" + thisSecond
    //thisSecond =
  }

  theTime := thisHour + ":" + thisMinute + ":" + thisSecond

  return theTime

}

func (t Time24) ValidTime() bool {

  timeIsValid := false

  if ((0 <= t.hour && t.hour < 24) && (0 <= t.minute && t.minute < 60) && (0 <= t.second && t.second < 60)) {
    timeIsValid = true
  } else {
    fmt.Printf("ERROR: This time: %v:%v:%v is not valid.\nHours   must be between 0 and 23\nMinutes must be between 0 and 59\nSeconds must be between 0 and 59\n",t.hour,t.second,t.minute)
  }


  return timeIsValid
}

func minTime24(times []Time24) (Time24, error) {

  if len(times) == 0 {
    err := errors.New("ERROR: Your []Time24 slice is empty.\n")
    minTime := Time24{hour:0, minute:0, second:0}
    return minTime, err
  }

  validMin := false
  var minTime Time24;
  for i := 0; validMin == false; i++{
    if times[i].ValidTime(){
      minTime = times[i]
      validMin = true
    }
  }

  for i := 0; i < len(times); i++ {
    if times[i].ValidTime() {
      if lessThanTime24(times[i], minTime) {
        minTime = times[i];
      }
    }
  }

  return minTime, nil
}

func linearSearch(x interface{}, lst interface{}) (int, error) {

  searchKey := reflect.ValueOf(x)
  list := reflect.ValueOf(lst)


  if reflect.TypeOf(x) != reflect.TypeOf(lst).Elem() {
    err := errors.New("ERROR: The data types do not match\n")
    return 404, err
  }

  for i := 0; i < list.Len(); i++ {
    if list.Index(i).Interface() == searchKey.Interface() {
      return i, nil
    }
  }

  return 0, nil

}


var counter int

func allGeneSeqs(n int) []string {
  counter = 0
  if n <= 0 {
    var emptyString []string
    return emptyString
  }

  numberOfGenes := int(math.Pow(4,float64(n)))

  geneSequence := make([]string, numberOfGenes)

  genes := []string{"A","C","G","T"}

  lengthOfSequence := n
  base := ""

  RecurseGenes(genes, geneSequence, lengthOfSequence, base)

  fmt.Println(geneSequence)

  return geneSequence
}


func RecurseGenes(genes []string, geneSequence []string, length int, thisGene string) string{
  if (length == 0) {
    geneSequence[counter] = thisGene
    counter++
    return thisGene
  }

  for i := 0; i < 4; i++ {
    newGene := thisGene + genes[i] //AA

    RecurseGenes(genes, geneSequence, length - 1, newGene)

  }

  return thisGene
}
