package main

func main() {
	println(comma("12345678"))
}

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}
