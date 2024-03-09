package subject

import (
	"github.com/RicardoKim/Study/DesignPattern/Observer/observer"
)

type ISubject interface {
	RegisterObserver(o observer.IObserver)
	RemoveObserver(o observer.IObserver)
	NotifyObservers()
}
