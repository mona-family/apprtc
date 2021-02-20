package collider

import "log"

type firebase struct {
}

func newFirebase() *firebase {
	return &firebase{}
}

func (fb *firebase) logCallStart()  {

	log.Printf("Starting call!")

}
