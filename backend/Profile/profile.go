package profile

import (
	"encoding/json"
	"fmt"
	storage "mtgolauncher/backend/Storage" // Import your storage package
	"os"
	"path"
	"path/filepath"
)

type ProfileList struct {
	Profiles []ProfileInfo `json:"profiles"`
}

type InventoryItem struct {
	ID       string `json:"_id"`
	Template string `json:"_tpl"`
	// ... other fields ...
	StackCount int `json:"upd.StackObjectsCount"`
}

type ProfileInfo struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Edition        string `json:"edition"`
	Side           string `json:"side"`
	Experience     int    `json:"experience"`
	Level          int    `json:"level"`
	Head           string `json:"head"`
	ProfileMods    []string
	ProfileConfigs []string
	SavageLockTime float64 `json:"SavageLockTime"`
	USDAmount      float64 `json:"usdAmount"`
	RubAmount      float64 `json:"rubAmount"`
	EURamount      float64 `json:"eurAmount"`
}

type ProfileRunT struct {
}

func NewProfileRunT() *ProfileRunT {
	return &ProfileRunT{}
}

func ParseAndSaveProfile(profilePath string) error {
	profileData, err := os.ReadFile(profilePath)
	if err != nil {
		return err
	}

	// Update the struct to include Inventory within Customization
	var profileStruct struct {
		Info struct {
			ID        string  `json:"id"`
			Username  string  `json:"username"`
			Edition   string  `json:"edition"`
			USDAmount float64 `json:"usdAmount"`
			RubAmount float64 `json:"rubAmount"`
			EURamount float64 `json:"eurAmount"`
		} `json:"info"`
		Characters struct {
			PMC struct {
				Inventory struct {
					Items []struct {
						ID       string `json:"_id"`
						Template string `json:"_tpl"`
						Upd      struct {
							StackObjectsCount int `json:"StackObjectsCount"`
						} `json:"upd`
					} `json:"items"`
				} `json:"Inventory"`
				Customization struct {
					Head string `json:"Head"`
				} `json:"Customization"`
				Info struct {
					Side       string `json:"Side"`
					Experience int    `json:"Experience"`
					Level      int    `json:"Level"`
				} `json:"Info"`
			} `json:"pmc"`
			scav struct {
				SavageLockTime float64 `json:"SavageLockTime"`
			}
		} `json:"characters"`
		ProfileMods    struct{}
		ProfileConfigs struct{}
	}

	if err := json.Unmarshal(profileData, &profileStruct); err != nil {
		return err
	}

	// Create maps to store the counts for each currency
	currencyCounts := map[string]int{
		"5449016a4bdc2d6f028b456f": 0, // Roubles
		"569668774bdc2da2298b4568": 0, // Euros
		"5696686a4bdc2da3298b456a": 0,
	}

	// Iterate through inventory and get our currency values ehehehehehe
	for _, item := range profileStruct.Characters.PMC.Inventory.Items {
		if item.Template == "5449016a4bdc2d6f028b456f" ||
			item.Template == "569668774bdc2da2298b4568" ||
			item.Template == "5696686a4bdc2da3298b456a" {
			currencyCounts[item.Template] += item.Upd.StackObjectsCount
		}
	}

	// Add the accumulated currency counts to the profileInfo struct
	profileInfo := ProfileInfo{
		ID:             profileStruct.Info.ID,
		Username:       profileStruct.Info.Username,
		Edition:        profileStruct.Info.Edition,
		Side:           profileStruct.Characters.PMC.Info.Side,
		Experience:     profileStruct.Characters.PMC.Info.Experience,
		Level:          profileStruct.Characters.PMC.Info.Level,
		Head:           profileStruct.Characters.PMC.Customization.Head,
		ProfileMods:    []string{},
		ProfileConfigs: []string{},
		SavageLockTime: profileStruct.Characters.scav.SavageLockTime,
		RubAmount:      float64(currencyCounts["5449016a4bdc2d6f028b456f"]),
		EURamount:      float64(currencyCounts["569668774bdc2da2298b4568"]),
		USDAmount:      float64(currencyCounts["5696686a4bdc2da3298b456a"]),
	}

	profileDir := filepath.Dir(profilePath)

	appDir, err := storage.GetAppDataDir()
	if err != nil {
		return err
	}

	profilesDir := path.Join(appDir, "profiles")
	if _, err := os.Stat(profilesDir); os.IsNotExist(err) {
		if err := os.MkdirAll(profilesDir, 0755); err != nil {
			return err
		}
	}

	profileDetailsPath := path.Join(profilesDir, "AKI", fmt.Sprintf("%s (%s)", profileInfo.Username, profileInfo.ID), "profile-details.json")

	if err := os.MkdirAll(filepath.Dir(profileDetailsPath), 0755); err != nil {
		return err
	}

	profileDetailsData, err := json.MarshalIndent(profileInfo, "", "    ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(profileDetailsPath, profileDetailsData, 0644); err != nil {
		return err
	}

	files, err := os.ReadDir(profileDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(profileDir, file.Name())
		destPath := filepath.Join(filepath.Dir(profileDetailsPath), file.Name())

		if err := storage.CopyFile(srcPath, destPath); err != nil {
			return err
		}
	}

	return nil
}

func (p ProfileRunT) GetProfiles(gameType string) (ProfileList, error) {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		return ProfileList{}, err
	}

	profilesDir := path.Join(appDir, "profiles", gameType)
	if _, err := os.Stat(profilesDir); os.IsNotExist(err) {
		if err := os.MkdirAll(profilesDir, 0755); err != nil {
			return ProfileList{}, err
		}
	}

	profileFolders, err := os.ReadDir(profilesDir)
	if err != nil {
		return ProfileList{}, err
	}

	var profileList ProfileList

	for _, profileFolder := range profileFolders {
		if profileFolder.IsDir() {
			profileDetailsPath := path.Join(profilesDir, profileFolder.Name(), "profile-details.json")

			profileData, err := os.ReadFile(profileDetailsPath)
			if err != nil {
				return ProfileList{}, err
			}

			var profileInfo ProfileInfo
			if err := json.Unmarshal(profileData, &profileInfo); err != nil {
				return ProfileList{}, err
			}

			// Add the parsed profile to the profile list
			profileList.Profiles = append(profileList.Profiles, profileInfo)
		}
	}
	fmt.Println(profileList.Profiles)

	return profileList, nil
}
