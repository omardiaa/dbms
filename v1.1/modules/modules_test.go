package modules_test

import (
	"testing"
	"modules/modules"
	"sync"
	"fmt"
)

func TestSet(t *testing.T){
	k := "1"
	expected_v := "{name:'custom_name'}"
	modules.DbSet(k, expected_v)
	v := modules.DbGet(k)
	if v != expected_v {
		t.Errorf("TestSet expected %s and got %s", expected_v, v)
	}
	fmt.Println("Set and get test passed.")
}

func TestUpdate(t *testing.T){
	k := "2"
	expected_v := "{name:'custom_name_2'}"
	modules.DbSet(k, "{name:'custom_name'}")
	modules.DbSet(k, expected_v)
	v := modules.DbGet(k)
	if v != expected_v {
		t.Errorf("TestSet expected %s and got %s", expected_v, v)
	}
	fmt.Println("Update and get test passed.")
}

func TestConcurrency(t *testing.T){
	const TEXT_SIZE = 10
	const THREADS = 100

	var wg sync.WaitGroup
	wg.Add(THREADS)
	var k,v, expected_v string

	for i:=0;i<THREADS;i++{
		k = ""
		v = ""
		for j:=0;j<TEXT_SIZE;j++{
			k += fmt.Sprintf("KEY_%d",i)
			v += fmt.Sprintf("VALUE_%d",i)
		}

		go func(key, value string){
			modules.DbSet(key, value)
			wg.Done()
		}(k, v)
	}
	fmt.Println("Waiting for all threads to write...")
	wg.Wait()
	fmt.Println("All threads have written successfully!")

	for i:=0;i<THREADS;i++{
		k = ""
		expected_v = ""
		for j:=0;j<TEXT_SIZE;j++{
			k += fmt.Sprintf("KEY_%d",i)
			expected_v += fmt.Sprintf("VALUE_%d",i)
		}
		v = modules.DbGet(k)
		if v != expected_v {
			t.Errorf("TestSet expected %s and got %s", expected_v, v)
		}
	}
	fmt.Println("Concurrency test passed!")
}

func BenchmarkDBGetWithHashMap(b *testing.B) {
	const COUNT = 10000

	for i:=0;i<COUNT;i++ {
		k := fmt.Sprintf("KEY_%d",i)
		v := fmt.Sprintf("VALUE_%d",i)
		modules.DbSet(k, v)
	}

	for n := 0; n < b.N; n++ {
		//Read first value to make it harder
		_ = modules.DbGet("KEY_0")
	}
}

func BenchmarkDBGetNoHashMap(b *testing.B) {
	const COUNT = 10000

	for i:=0;i<COUNT;i++ {
		k := fmt.Sprintf("KEY_%d",i)
		v := fmt.Sprintf("VALUE_%d",i)
		modules.DbSet(k, v)
	}
	modules.DbGet("KEY_0") //To load data in the hashmap

	for n := 0; n < b.N; n++ {
		//Read first value to make it harder
		_ = modules.DbGetOld("KEY_0")
	}
}
