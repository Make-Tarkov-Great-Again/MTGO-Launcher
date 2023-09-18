package launch

import (
	"fmt"
	launcher "mtgolauncher/backend/Launcher"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"github.com/shirou/gopsutil/process"
)

func isProcessRunning(processName string) (bool, error) {
	processList, err := process.Processes()
	if err != nil {
		return false, err
	}

	for _, p := range processList {
		name, _ := p.Name()
		if name == processName {
			return true, nil
		}
	}

	return false, nil
}

func Launch( /*folderPath string, sessionID string, address string*/ ) {
	var wg sync.WaitGroup
	UI := launcher.NewUI()

	//Special people safety, wait for both to finish before contuining.
	wg.Add(1)
	go func() {
		defer wg.Done()
		bsgLauncherRunning, err := isProcessRunning("BsgLauncher.exe")
		if err != nil {
			fmt.Printf("Error checking for BsgLauncher process: %v\n", err)
			return
		}

		if bsgLauncherRunning {
			fmt.Println("BsgLauncher is running. Panic time")
			defer UI.PanicStatement("Battlestate Games Launcher is running.", "The Battlestate Games Launcher is running. Continuing could lead to a live ban. To protect you, I'm shutting down.")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		escapeFromTarkovRunning, err := isProcessRunning("EscapeFromTarkov.exe")
		if err != nil {
			fmt.Printf("Error checking for EscapeFromTarkov.exe process: %v\n", err)
			return
		}

		if escapeFromTarkovRunning {
			fmt.Println("EscapeFromTarkov.exe is already running.")
			UI.ErrorStatement("Tarkov 2 Electric Boogaloo", "There is already another instance of Tarkov running. Please close the other instance before launching.")
		}
	}()

	gameExecutable := "EscapeFromTarkov.exe"
	gameFolderPath := "E:/Tarkov"
	userToken := "BallsInYoJaw"
	backendURL := "http://127.0.0.1:6969"
	//arugments are nesscary or people will get banned lol
	arguments := fmt.Sprintf("-token=%s -config={'BackendUrl':'%s','Version':'live'}", userToken, backendURL)

	wg.Wait()
	cmd := exec.Command(gameExecutable)
	cmd.Dir = gameFolderPath //working dir

	//Output, but im pretty sure there wont be any
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Set the arguments
	cmd.Args = append(cmd.Args, arguments)

	//start
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting the game: %v\n", err)
		UI.ErrorStatement("Failed to start game", fmt.Sprintf("Failed to start game with error: \n%s", err))
		return
	}
	// Wait for the game process to finish
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode := status.ExitStatus()
				fmt.Printf("Game exited with code %d\n", exitCode)
				//launcher.Game.postGameFuncs
			}
		} else {
			fmt.Printf("Error waiting for the game to finish: %v\n", err)
		}
	}
}
