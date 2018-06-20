package main 

import (
	"fmt"
)

type Person struct{
	name string
}

//Worker is a Person
type Worker struct{
	Person
	job string 
}

func (p *Person) Talk(){
	fmt.Println("Hi, name is:",p.name)
}

func (w *Worker) TalkAsWorker(){
	w.Talk()
	fmt.Println("I have a job:", w.job)
}

func (p *Person) ChangeName(){
	p.name = "John"
}

func changeNameWithParam(p *Person){
	p.name = "Eric"
}

func GetNames(people ...Person) string{
	var result = "Names["	

	for i:= 0; i<len(people); i++{
		result += people[i].name
		if(i != len(people)-1){
			result += ","
		}
	}
	result += "]"
	return result
}

func main() {

	personObj := Person{"Steve"}
	personObj.Talk()
	
	personObj.ChangeName()
	personObj.Talk()
	
	changeNameWithParam(&personObj)
	personObj.Talk()
	
	workerObj := &Worker{
					Person:Person{"Meggie"},
					job:"Software"	,	
				}
	workerObj.TalkAsWorker()
	
	fmt.Println(GetNames(personObj, workerObj.Person))
}

