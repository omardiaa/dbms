package main

import(
	"fmt"
	"modules/modules"
)

func main(){
	for {
		var command, key, value string
    fmt.Scan(&command)
    fmt.Scan(&key)
	
		switch command{
		case "db_get":
			v := modules.DbGet(key)
			fmt.Println(v)
			break;
		case "db_set":
			fmt.Scan(&value)
			modules.DbSet(key, value)
			break;
		case "db_del":
			modules.DbDel(key)
			break;
		}
	}
}
