package modules

import(
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)
const DB_FILEPATH = "./database"
var		HASH_TABLE	= make(map[string]int)
var		MUTEX sync.Mutex

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func reverse(arr []string) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func loadHashTableIfNotLoaded(){
	if len(HASH_TABLE) != 0 {
		return
	}

	db, err := os.ReadFile(DB_FILEPATH)
	check(err)

	r1, _ := regexp.Compile(`\n`)
	lines := r1.Split(string(db), -1)
	
	r2, _ := regexp.Compile(`[^,]+`)

	var offset int
	offset = 0

	for _, line := range(lines){
		parsed_line := r2.FindAllString(line, -1)
		if len(parsed_line) == 0 {
			continue
		}
		k 	:= parsed_line[0]
		v 	:= offset + len(k) + 1
		HASH_TABLE[k] = v
		offset += int(len(line) + 1)
	}
}

func DbGet(key string) string{
	loadHashTableIfNotLoaded()

	db, err := os.ReadFile(DB_FILEPATH)
	check(err)

	offset := HASH_TABLE[key]

	r1, _ := regexp.Compile(`\n`)
	value := r1.Split(string(db)[offset:], -1)[0]
	
	return value
}

func DbGetOld(key string) string{
	db, err := os.ReadFile(DB_FILEPATH)
	check(err)

	r1, _ := regexp.Compile(`\n`)
	lines := r1.Split(string(db), -1)

	reverse(lines)
	
	r2, _ := regexp.Compile(`[^,]+`)

	for _, line := range(lines){
		parsed_line := r2.FindAllString(line, -1)
		if len(parsed_line) == 0 {
			continue
		}
		k 	:= parsed_line[0]
		v 	:= strings.Join(parsed_line[1:], ",")
		if k==key{
			return v
		}
	}
	return ""
}

func DbSet(key, value string){
	loadHashTableIfNotLoaded()
	
	MUTEX.Lock() 
	defer MUTEX.Unlock()
	file, err := os.OpenFile(DB_FILEPATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	fileInfo, err := file.Stat()
	check(err)
	
	defer file.Close()

	HASH_TABLE[key] = int(fileInfo.Size()) + len(key) + 1
	//Mutex is needed before opening the file because the fileInfo may change by another thread

	s := fmt.Sprintf("%s,%s\n",key,value)
	_, err = file.WriteString(s)
	//Mutex is needed here to avoid writing to the file in a location different than the offset stored in the hashmap
	check(err)
}

func DbDel(key string){

}
