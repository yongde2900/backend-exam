package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct {
	ID    int
	Count int
}

func (e *Employee) work(item Item) {
	now := time.Now()
	fmt.Printf("員工編號%d開始處理%s\n", e.ID, item.Name())
	item.Process()
	fmt.Printf("員工編號%d處理%s完畢，耗時 %f 秒\n", e.ID, item.Name(), time.Since(now).Seconds())
	e.Count++
}

type Item1 struct{}

func (Item1) Process() {
	time.Sleep(300 * time.Millisecond)
}

func (Item1) Name() string {
	return "Item1"
}

type Item2 struct{}

func (Item2) Process() {
	time.Sleep(500 * time.Millisecond)
}

func (Item2) Name() string {
	return "Item2"
}

type Item3 struct{}

func (Item3) Process() {
	time.Sleep(800 * time.Millisecond)
}

func (Item3) Name() string {
	return "Item3"
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
	Name() string
}

func worker(itemsCh chan Item, wg *sync.WaitGroup, employee *Employee) {
	for item := range itemsCh {
		employee.work(item)
		wg.Done()
	}
}

const (
	quantityPerItem = 10
	employeeCount   = 5
)

func main() {
	now := time.Now()
	items := make([]Item, 0, quantityPerItem*3)
	for range quantityPerItem {
		items = append(items, Item1{}, Item2{}, Item3{})
	}

	rand.Shuffle(len(items), func(i int, j int) { items[i], items[j] = items[j], items[i] })

	var wg sync.WaitGroup
	itemCh := make(chan Item, 3)

	// 啟用員工
	employees := make([]*Employee, 0, employeeCount)
	for i := range employeeCount {
		employees = append(employees, &Employee{ID: i + 1})
		go worker(itemCh, &wg, employees[i])
	}

	// 放入工作項目
	wg.Add(len(items))
	for _, item := range items {
		itemCh <- item
	}

	close(itemCh)

	wg.Wait()

	fmt.Printf("總共耗時 %f 秒\n", time.Since(now).Seconds())

	for _, employee := range employees {
		fmt.Printf("員工編號%d 處理 %d項物品\n", employee.ID, employee.Count)
	}
}
