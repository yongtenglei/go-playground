package main

import "fmt"

func SendNotification(user string) chan string {
	notification := make(chan string, 500)

	go func() {
		notification <- fmt.Sprintf("hello, %s, welcome to the world of GO", user)
	}()
	return notification
}

func main() {
	rey := SendNotification("rey")
	charlotte := SendNotification("charlotte")

	fmt.Println(<-rey)
	fmt.Println(<-charlotte)

}
