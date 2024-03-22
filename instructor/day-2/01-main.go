package main

func main() {
	panic("dummy")
}

// n:m scheduler
// n: no of goroutines -> this can be more (this is managed by Go scheduler 2KB)
// m: no of OS threads -> this can be less (how less depends on the work)
