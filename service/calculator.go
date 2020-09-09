package service

import (
	"context"
	"log"
	"sync"

	"github.com/michaelsidharta/cubic-weight/entity"
	"github.com/michaelsidharta/cubic-weight/external"
)

type ICalculator interface {
	GetAverage(category string) (float64, error)
}

type sumResult struct {
	Sum   float64
	Count int64
}

type Calculator struct {
	RestAPI external.IAPI
}

func InitCalculator(a external.IAPI) ICalculator {
	return &Calculator{RestAPI: a}
}

func (c Calculator) GetAverage(category string) (float64, error) {
	var res float64
	agg := make(chan sumResult)
	done := make(chan bool)
	aggResult := aggregate(agg, done)
	next := "/api/products/1"
	var wg sync.WaitGroup
	for next != "" {
		resp, err := c.RestAPI.Get(context.Background(), next)
		if err != nil {
			log.Println("Error in getting REST API", err)
			return 0.0, err
		}
		acObjects := resp.FilterObjectByCategory(category)
		if len(acObjects) > 0 {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				agg <- sum(acObjects, category)
			}(&wg)
		}
		next = resp.Next
	}
	wg.Wait()
	done <- true
	res = <-aggResult
	return res, nil
}

func sum(objects []entity.Object, category string) sumResult {
	res := sumResult{}
	for _, obj := range objects {
		res.Sum += obj.Size.CubicWeight()
		res.Count++
	}
	return res
}

func aggregate(agg chan sumResult, done chan bool) <-chan float64 {
	c := make(chan float64)
	go func() {
		sum := float64(0)
		count := int64(0)
		for {
			select {
			case val := <-agg:
				sum += val.Sum
				count += val.Count
			case <-done:
				if count == 0 {
					c <- 0.0
					break
				}
				c <- sum / float64(count)
				break
			}
		}
	}()
	return c
}
