package lib

import (
	"errors"
	"fmt"
)

// GoTickets 表示Goroutine票池的接口。
//“票”-抽象概念，它相当于程序为了启用一个goroutine而必须持有的令牌
//goroutine票池只负责增减票的数量，并以此真实地体现出正在运行的专用goroutine的数量
type GoTickets interface {
	// 拿走一张票。
	Take()
	// 归还一张票。
	Return()
	// 票池是否已被激活。
	Active() bool
	// 票的总数。
	Total() uint32
	// 剩余的票数。
	Remainder() uint32
}

// myGoTickets 表示Goroutine票池的实现。
type myGoTickets struct {
	total    uint32        // 票的总数。
	ticketCh chan struct{} // 票的容器。
	active   bool          // 票池是否已被激活（正确的初始化）。
}

// NewGoTickets 会新建一个Goroutine票池。
func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg :=
			fmt.Sprintf("The goroutine ticket pool can NOT be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}

//对lib.myGoTickets类型值的初始化工作都由它的包级私有的指针方法init来进行
func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)  //初始化一个元素类型为struct{}的缓冲通道ch
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take() {
	<-gt.ticketCh
}

func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}

func (gt *myGoTickets) Active() bool {
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
