package main

import (
	"fmt"
	"hscan/hscan"
	"log"
	"os"
)

func main() {

	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0"
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032"

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <filepath>")
	}
	var file = os.Args[1]

	hscan.GenHashMaps(file)
	s, err := hscan.GetMD5(drmike1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	hscan.GuessSingle(md5hash, file)
	hscan.GuessSingle(sha256hash, file)
	hscan.GenHashMaps(file)
	hscan.GetSHA(sha256hash)
	hscan.GetMD5(md5hash)

	hscan.GuessSingle(drmike1, file)
	hscan.GuessSingle(drmike2, file)
}
