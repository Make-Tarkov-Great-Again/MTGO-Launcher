<template>
   <div class="settings-popout">
    <h2>Settings</h2>
    <a>* Changed</a>
    <div class="setting">
      <label for="AkiPath">AKI Server Path:</label>
      <input type="text" id="AkiPath" v-model="akiServerPath" readonly
        @click="openFolderDialog('AkiPath', 'Select your AKI Server folder', filters3)" />
      <button @click="openFolderDialog('AkiPath', 'Select your AKI Server folder', filters3)">Browse</button>
    </div>

    <div class="setting">
      <label for="MtgaPath">MTGA Server Path:</label>
      <input type="text" id="MtgaPath" v-model="mtgaServerPath" readonly
        @click="openFolderDialog('MtgaPath', 'Select your MTGA Server folder', filters2)" />
      <button @click="openFolderDialog('MtgaPath', 'Select your MTGA Server folder', filters2)">Browse</button>
    </div>

    <div class="setting">
      <label for="ClientPath">Client Path:</label>
      <input type="text" id="ClientPath" v-model="clientPath" readonly @click="openFolderDialog('ClientPath','Select your NON LIVE Tarkov folder (Where eft.exe is)', filters1)" />
      <button @click="openFolderDialog('ClientPath','Select your NON LIVE Tarkov folder (Where eft.exe is)', filters1)">Browse</button>
    </div>

    <!-- Placeholder settings -->
    <div class="setting">
      <label for="Language">Language:</label>
      <input type="text" id="Language" v-model="placeholderSetting1" />
    </div>

    <div class="setting">
      <label for="Theme">Theme:</label>
      <input type="text" id="Theme" v-model="placeholderSetting2" />
      <br>
      <a>Wanna spice it up? Load custom css!</a>
    </div>
  </div>
</template>

<script>
import { GetRuntimeConfig } from "../../wailsjs/go/config/ConfigRunT"
import { OpenFileSelector } from "../../wailsjs/go/launcher/UI"
import { SetConfigVariable } from "../../wailsjs/go/config/ConfigRunT"
export default {
  data() {
    return {
      akiServerPath: "",
      mtgaServerPath: "",
      clientPath: "",
      placeholderSetting1: "",
      placeholderSetting2: "",
    };
  },
  data() {
    return {
      akiServerPath: "",
      mtgaServerPath: "",
      clientPath: "",
      placeholderSetting1: "",
      placeholderSetting2: "",
      filters1: [
        { displayName: 'Select your Client folder', pattern: '*' },
      ],
      filters2: [
        { displayName: 'Select MTGA server path', pattern: '*' },
      ],
      filters3: [
        { displayName: 'Select AKI server path', pattern: '*' },
      ],
    };
  },
  methods: {
    updatePopoutPosition() {
      const scrollTop = window.scrollY || window.scrollY;
      const scrollLeft = window.scrollX || window.scrollX;

      this.popoutTop = 50 + scrollTop;
      this.popoutLeft = 50 + scrollLeft;
    },
    openFileDialog(property) {
    },
    async openFolderDialog(inputField, title, filters) {
      const selectedPath = await OpenFileSelector(title, filters);
      document.getElementById(inputField).value = selectedPath;
      await SetConfigVariable(inputField, selectedPath)
      const labelElement = document.querySelector(`label[for="${inputField}"]`);
      labelElement.innerHTML = labelElement.innerHTML + "\*"
    },
    async fetchData() {
      try {
        // Call GetRuntimeConfig to retrieve the configuration data
        const runtimeConfig = await GetRuntimeConfig();

        // Set the component's data properties based on the received data
        this.akiServerPath = runtimeConfig.UserSettings.server.akiServerPath;
        this.mtgaServerPath = runtimeConfig.UserSettings.server.mtgaServerPath;
        this.clientPath = runtimeConfig.UserSettings.clientPath;
        this.placeholderSetting1 = runtimeConfig.UserSettings.language;
        this.placeholderSetting2 = runtimeConfig.UserSettings.theme;
      } catch (error) {
        // Handle any errors here, e.g., log or display an error message
        console.error("Error fetching configuration:", error);
      }
    },
  },
  mounted() {
    this.fetchData()
  }
};
</script>

<style scoped>
.settings-popout {
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #010409;
  width: 300px;
  position: fixed;
  z-index: 99999;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.setting {
  margin-bottom: 10px;
}

label {
  font-weight: bold;
  display: inline-block;
  width: 150px;
}

input[type="text"] {
  width: 60%;
  padding: 5px;
}

button {
  padding: 5px 10px;
  cursor: pointer;
}
</style>
