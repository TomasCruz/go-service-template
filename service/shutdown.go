package service

// Shutdown shuts down the service operations
func Shutdown() {
	wg.Wait()
	//fmt.Println("Graceful shutdown")
}
