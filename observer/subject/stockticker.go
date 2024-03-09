package subject

import (
	"github.com/RicardoKim/Study/DesignPattern/Observer/observer"
)

type StockTicker struct {
	//Observer들은 동적으로 변화하기 때문에 slice객체로 선언한다.
	observers   []observer.IObserver
	stockSymbol string
	//price또한
	price float64
}

func (s *StockTicker) RegisterObserver(o observer.IObserver) {
	//slice의 append는 append(붙일 배열, 추가할 값들)로 선언하게 된다.
	s.observers = append(s.observers, o)
}

func (s *StockTicker) RemoveObserver(o observer.IObserver) {
	for i, observer := range s.observers {
		if observer == o {
			//i-1번째 뒤에 i + 1 배열을 붙임으로써 삭제를 진행한다.
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *StockTicker) NotifyObservers() {
	//순회하면서 stocke의 이름과 가격을 전달한다.
	for _, observer := range s.observers {
		observer.Update(s.stockSymbol, s.price)
	}
}

//stockeSymbol은 소문자로 시작하기 때문에 같은 패키지에서만 접근이 가능하다.
//따라서 setter를 만들어 수정이 되게끔 만든다.
func (s *StockTicker) SetStockSymbol(stockSymbol string) {
	s.stockSymbol = stockSymbol
}

func (s *StockTicker) SetPrice(price float64) {
	//push방식으로 동작하는 것을 알 수 있다.
	s.price = price
	s.NotifyObservers()
}
