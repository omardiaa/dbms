package modules

import(
	"fmt"
	"os"
	"regexp"
	"strings"
)
const DB_FILEPATH = "./database"
var		HASH_TABLE	= make(map[string]string)

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

	for _, line := range(lines){
		parsed_line := r2.FindAllString(line, -1)
		if len(parsed_line) == 0 {
			continue
		}
		k 	:= parsed_line[0]
		v 	:= strings.Join(parsed_line[1:], ",")
		HASH_TABLE[k] = v
	}
}

func DbGet(key string) string{
	loadHashTableIfNotLoaded()

	return HASH_TABLE[key]
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

	file, err := os.OpenFile(DB_FILEPATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)

	defer file.Close()

	HASH_TABLE[key] = value
	s := fmt.Sprintf("%s,%s\n",key,value)
	_, err = file.WriteString(s)
	check(err)
}

func DbDel(key string){

}
