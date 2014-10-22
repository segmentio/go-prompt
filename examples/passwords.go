package main

import ".."

func main() {
	{
		pass := prompt.Password("something")
		println(pass)
	}

	{
		pass := prompt.PasswordMasked("something")
		println(pass)
	}
}
