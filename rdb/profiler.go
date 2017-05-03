package rdb


import (
	"fmt"
	"os"
	"github.com/cupcake/rdb"
)


type MemoryUsage struct {
	StringUsage    int64
	HashUsage      int64
	SetUsage       int64
	ListUsage      int64
	SortedSetUsage int64
}

func newMemoryUsage() *MemoryUsage {
	return &MemoryUsage{}
}

func (memoryUsage *MemoryUsage) GetTotal() int64 {
	return memoryUsage.StringUsage + memoryUsage.HashUsage + memoryUsage.SetUsage + memoryUsage.ListUsage + memoryUsage.SortedSetUsage
}

type Profiler struct {
	memoryUsages map[int]*MemoryUsage
	currentDatabase int
}

func NewProfiler() *Profiler {
	return &Profiler{
		make(map[int]*MemoryUsage),
		1,
	}
}

func (profiler *Profiler) StartProfile(filepath string) (map[int]*MemoryUsage, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	rdb.Decode(f, profiler)

	return profiler.memoryUsages, nil
}

func (profiler *Profiler) getCurrentMemoryUsage() *MemoryUsage {
	return profiler.memoryUsages[profiler.currentDatabase]
}

func (profiler *Profiler) setCurrentMemoryUsage(memoryUsage *MemoryUsage) {
	profiler.memoryUsages[profiler.currentDatabase] = memoryUsage
}

func (profiler *Profiler) StartRDB() {}

func (profiler *Profiler) StartDatabase(n int) {
	profiler.currentDatabase = n
	if profiler.getCurrentMemoryUsage() == nil {
		profiler.setCurrentMemoryUsage(newMemoryUsage())
	}
}

func (profiler *Profiler) Aux(key, value []byte) {
	// TODO: what is it?
}

func (profiler *Profiler) ResizeDatabase(dbSize, expiresSize uint32) {
	// TODO
}

func (profiler *Profiler) Set(key, value []byte, expiry int64) {
	profiler.getCurrentMemoryUsage().SetUsage += int64(len(key) + len(value))
}

func (profiler *Profiler) StartHash(key []byte, length, expiry int64) {
	profiler.getCurrentMemoryUsage().HashUsage += int64(len(key)) + length
}

func (profiler *Profiler) Hset(key, field, value []byte) {
	// TODO IMME
}
func (profiler *Profiler) EndHash(key []byte)                              {}
func (profiler *Profiler) StartSet(key []byte, cardinality, expiry int64)  {}
func (profiler *Profiler) Sadd(key, member []byte)                         {}
func (profiler *Profiler) EndSet(key []byte)                               {}
func (profiler *Profiler) StartList(key []byte, length, expiry int64)      {}
func (profiler *Profiler) Rpush(key, value []byte)                         {}
func (profiler *Profiler) EndList(key []byte)                              {}
func (profiler *Profiler) StartZSet(key []byte, cardinality, expiry int64) {}
func (profiler *Profiler) Zadd(key []byte, score float64, member []byte)   {}
func (profiler *Profiler) EndZSet(key []byte)                              {}
func (profiler *Profiler) EndDatabase(n int)                               {}
func (profiler *Profiler) EndRDB()                                         {}


func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO remove
func test() {
	fmt.Println("start")
	defer fmt.Println("end")

	f, err := os.Open("./dump-master.rdb")
	checkErr(err)
	decoder := &Profiler{}

	rdb.Decode(f, decoder)
}
