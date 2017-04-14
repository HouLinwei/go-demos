package main

import (
	"github.com/rcrowley/go-metrics"
	"fmt"
	"time"
	"log"
	"os"
)

func main()  {
	c := metrics.NewCounter()
	metrics.Register("counter", c)
	c.Inc(100)
	fmt.Println(c)
	fmt.Println(metrics.Get("counter"))

	g := metrics.NewGauge()
	metrics.Register("gauge", g)
	g.Update(101)
	fmt.Println(metrics.Get("gauge"))

	//r := metrics.NewRegistry()
	//g := metrics.NewRegisteredFunctionalGauge("xxx", r, func() int64{
	//	return cache.get
	//})

	go metrics.Log(metrics.DefaultRegistry, 5 * time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	time.Sleep(100*time.Second)
}
