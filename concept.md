# MTGO Launcher - Concept Explanation

MTGO Launcher is an emulator launcher that offers a range of features to enhance the modding experience. This document explains two key features of the launcher: Subscription-based mod updates and Mod Profiles management.

## Subscription-Based Mod Downloading/Updates

MTGO Launcher provides a seamless way for users to update their mods through a subscription-based model. Here's how it works:

1. **Checking for Updates**: When the user launches the launcher, it scans the mods folder of the emulator to identify installed mods.

2. **Mod Information**: Each mod in the mods folder contains a `modinfo.mtgo` file that stores metadata about the mod, including the mod's ID and version.

3. **Checking for Updates**: The launcher compares the version of each installed mod with the version listed in the `modinfo.mtgo` file.

4. **Automatic Updates**: If a newer version is available, the launcher automatically downloads and installs the updated mod, ensuring the user has the latest improvements and fixes.

5. **Manual Updates**: Users can also manually trigger updates by clicking the "Get" button on a mod page, which will download and install the latest version of the selected mod.

This process streamlines the update process, ensuring that users always have access to the latest versions of their favorite mods.

## Mod Profiles Management

MTGO Launcher offers the ability to create and manage Mod Profiles, allowing users to switch between different sets of mods and configurations. Here's how it's implemented:

1. **Profile Storage**: The launcher stores mod profiles within its own folder. Each profile includes information about the mods and configurations associated with it.

2. **Dynamic Switching**: When a user selects a different mod profile, the launcher dynamically swaps out the mods and configurations according to the selected profile.

3. **Mod Packs Compatibility**: Mod profiles created using mod packs are tied to those packs. If a user switches to a different mod pack or back to vanilla, the launcher adjusts the mods and configurations accordingly.

This feature enhances user flexibility by enabling them to customize their modding experience for different gameplay scenarios or preferences.
