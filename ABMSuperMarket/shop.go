package main
import "math/rand"
import "time"
import "fmt"


var timeofday int

var handsanitizerremaining int
func timeloop(){
    
    timeofday=timeofday+1
}
func setCovid() int  {
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := 5
     covidlv:= rand.Intn((max - min + 1) + min)
   

var maxcapacity int
if(covidlv==0){

  maxcapacity = 100  
    
}
if covidlv == 1 {
maxcapacity = 100 
}

if covidlv == 2 {

maxcapacity = 75

}
if covidlv == 3 {

maxcapacity = 50

}

if covidlv == 4 {

maxcapacity = 25

}

if covidlv == 5 {

maxcapacity = 10

}


return maxcapacity

}

func openshop() {
    
    fmt.Println("Tills opening") 

}

func customer(){
    
     fmt.Println("customers incoming") 
    handsanitizerremaining= handsanitizerremaining-1
}

func closeshop(){
    
    fmt.Println("no more customers allowed") 
    fmt.Println("processremaining customers") 
    fmt.Println("close all tills")
    fmt.Println("close shop") 
    
}



func handsanitizer(){
    
    if handsanitizerremaining==0{
    
    
        fmt.Println("refiling hand sanitizer") 
        handsanitizerremaining=100
    }
    
}

func main() {
    timeofday=540
    
    handsanitizerremaining=100
    
   maxcapacity := setCovid()

   fmt.Println(maxcapacity)
   
   if timeofday== 540{
       
       openshop()
   }
     if timeofday==1320{
         
         closeshop()
     }
     
     for timeofday < 1320 && timeofday >= 540{
		
		customer()
		handsanitizer()
		timeloop()
		
	}
	for timeofday > 1320 && timeofday < 540{
		
		fmt.Println("shop is closed")
		
	}
}
