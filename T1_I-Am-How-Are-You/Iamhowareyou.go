package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Reference: https://tutorialedge.net/ and https://gobyexample.com/ and https://www.quantamagazine.org/how-godels-incompleteness-theorems-work-20200714/

func main() {

	r := bufio.NewReader(os.Stdin)

	fmt.Println("I am Agoprogam.  How are you (good/any other typed input)?   ")
	temp, errors := r.ReadString('\n')

	temp = strings.Replace(temp, "\n", "", -1)

	if errors != nil {
		fmt.Println("something's wrong, sorry")
	}

	if strings.Contains(temp, "good") {
		fmt.Println("Great!")
	} else {
		fmt.Println("Sounds like you are definitely not good right now.")
	}

	fmt.Println("Say check this out. Let '0' be 6, let the successor function be 7, let '(' be 8, let ')' be 9, let '=' be 5, and let '+' be 11.")
	fmt.Println("Then...")
	fmt.Println("Wait for it...")
	fmt.Println("The formula '1 + 2 = 3' has the Godel number 263871010580431552941101982536173915134895149480003604169911130273946997941207917298748114520726903950820819280916133192695469786920156944682314047557861444082586688365002070884826660983439423268888302848351517888024116980474108961030912541130206000000 and is true (if I did my basic math right). What do you think about that!")
	r.ReadString('\n')
	fmt.Println("Well it is now my favorite number and definition of truth")

}
