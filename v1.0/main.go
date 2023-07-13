package main

import(
	"fmt"
	"modules/modules"
	"os"
)

func main(){
	args 				:= os.Args[1:]

	command 		:= args[0]
	key 				:= args[1]

	switch command{
	case "db_get":
		v := modules.DbGet(key)
		fmt.Println(v)
		break;
	case "db_set":
		value := args[2]
		modules.DbSet(key, value)
		break;
	case "db_del":
		modules.DbDel(key)
		break;
	}
}
