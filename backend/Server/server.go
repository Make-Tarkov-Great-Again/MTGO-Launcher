package server

//Made this before mtgolauncher/backend/Launcher, and dunno where im gonna go from here, so keeping this as reference.

//package main
//
//import (
//	"fmt"
//	"os"
//	"os/exec"
//	config "mtgolauncher/backend/Storage/config"
//)
//
//func main() {
//	exePath := config.localvar.serverPath
//	workingDir := "C:/path/to/working/directory"
//
//	cmd := exec.Command(exePath)
//	cmd.Dir = workingDir
//
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//
//	err := cmd.Start()
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	err = cmd.Wait()
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	fmt.Println("Child process has finished.")
//}
//
