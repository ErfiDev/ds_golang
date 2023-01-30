package main

import "fmt"

type Queue []Order

type Order struct {
	priority     int
	customerName string
	product      string
}

func (q *Queue) add(o Order) {
	if len(*q) == 0 {
		*q = append(*q, o)
	} else {
		appended := false
		for i, order := range *q {
			if o.priority > order.priority {
				*q = append((*q)[:i], append(Queue{o}, (*q)[i:]...)...)
				appended = true
				break
			}
		}

		if !appended {
			*q = append(*q, o)
		}
	}
}

func main() {
	q := Queue{}

	q.add(Order{150, "john", "iphone"})
	q.add(Order{48, "robert", "bmw"})
	q.add(Order{36, "ali", "headphone"})
	q.add(Order{220, "hany", "charger"})
	q.add(Order{77, "markus", "xbox"})

	fmt.Println(len(q))

	for _, o := range q {
		fmt.Printf("priority: %d, customer: %s\n", o.priority, o.customerName)
	}
}
