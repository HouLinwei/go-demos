package go_context

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("cancel ctx")
		cancel()
	}()
	fmt.Println(A(ctx))

	//result is as follows, why?
	//cancel ctx
	//C
	//B
	//A
	//A Done
	//C Done
	//B Done
}

//func D(ctx context.Context) string {
//	select {
//	case <-ctx.Done():
//		fmt.Println("D")
//		return "D Done"
//	}
//	return ""
//}

func C(ctx context.Context) string {
	//go fmt.Println(D(ctx))
	select {
	case <-ctx.Done():
		fmt.Println("C")
		return "C Done"
	}
	return ""
}

func B(ctx context.Context) string {
	//ctx, _ = context.WithCancel(ctx)
	go fmt.Println(C(ctx))
	select {
	case <-ctx.Done():
		fmt.Println("B")
		return "B Done"
	}
	return ""
}

func A(ctx context.Context) string {
	go fmt.Println(B(ctx))
	select {
	case <-ctx.Done():
		fmt.Println("A")
		//time.Sleep(time.Second * 1)
		return "A Done"
	}
	return ""
}
