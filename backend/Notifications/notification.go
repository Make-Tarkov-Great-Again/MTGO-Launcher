package notifications

import (
	"fmt"
	"log"

	"gopkg.in/toast.v1"
)

func Download(downloaded string) {
	notification := toast.Notification{
		AppID:   "MTGO-Launcher",
		Title:   "Finished downloading",
		Message: fmt.Sprintf("Download for %s is complete", downloaded),
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func ModsUpdated() {
	notification := toast.Notification{
		AppID:   "MTGO-Launcher",
		Title:   "Finished updating your mods",
		Message: "All your mods are up to date!",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateAvaiable() {
	notification := toast.Notification{
		AppID:   "MTGO-Launcher",
		Title:   "A new version of MT-GO Launcher is available",
		Message: "Click below to download",
		Actions: []toast.Action{
			{Type: "protocol", Label: "Download Now", Arguments: "https://github.com/Make-Tarkov-Great-Again/MTGO-Launcher/releases"},
		}}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
