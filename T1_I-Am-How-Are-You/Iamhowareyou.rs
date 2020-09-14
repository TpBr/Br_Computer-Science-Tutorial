//reference https://www.tutorialspoint.com and https://www.quantamagazine.org/how-godels-incompleteness-theorems-work-20200714/

fn main() {

    //println!("hello world!");

    let mut temp = String::new();
    println!("I am Arustprogam.  How are you (good/any other typed input)? \n");
    let temp2 = std::io::stdin().read_line(&mut temp).unwrap();
    
    if temp2 == 6 {
      println!("Great!");
    } else {
      println!("Sounds like you are definitely not good right now.");
    }
    println!("Say check this out. Let '0' be 6, let the successor function be 7, let '(' be 8, let ')' be 9, let '=' be 5, and let '+' be 11.");

    println!("Then...");

    println!("Wait for it...");
    println!("The formula '1 + 2 = 3' has the Godel number 263871010580431552941101982536173915134895149480003604169911130273946997941207917298748114520726903950820819280916133192695469786920156944682314047557861444082586688365002070884826660983439423268888302848351517888024116980474108961030912541130206000000 and is true (if I did my basic math right). What do you think about that! \n");
	  let _temp3: String = std::io::stdin().read_line(&mut temp).unwrap().to_string();
    println!("Well it is now my favorite number and definition of truth.");

  }
  