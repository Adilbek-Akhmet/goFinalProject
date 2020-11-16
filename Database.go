package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)
type Database interface {
	search(name string) [3]string
	showOrderList()
}

type singletonDatabase struct {
	orderList map[string] [3]string
}
func (db *singletonDatabase) search(name string) [3]string  {
	return db.orderList[name]
}
func (db *singletonDatabase) showOrderList() {
	fmt.Println(db)
	fmt.Println()
}
func readData(path string) (map[string][3]string, error) {
	file, err := os.Open( path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := map[string][3]string{}
	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v1 := scanner.Text()
		scanner.Scan()
		v2 := scanner.Text()
		scanner.Scan()
		v3 := scanner.Text()
		result[k] = [3]string{v1, v2, v3}
	}
	return result, nil
}
var once sync.Once
var instance Database
func GetSingletonDatabase() Database  {
	once.Do(func() {
		caps,_ := readData("./order.txt")
		db := singletonDatabase{caps}
		instance = &db
	})
	return instance
}

