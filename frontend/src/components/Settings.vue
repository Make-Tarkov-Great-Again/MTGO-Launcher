<template>
    <div class="settings-popout">
      <h2>Settings</h2>
      <div class="setting">
        <label for="aki-server-path">AKI Server Path:</label>
        <input
          type="text"
          id="aki-server-path"
          v-model="akiServerPath"
          readonly
          @click="openFileDialog('akiServerPath')"
        />
        <button @click="openFileDialog('akiServerPath')">Browse</button>
      </div>

      <div class="setting">
        <label for="mtga-server-path">MTGA Server Path:</label>
        <input
          type="text"
          id="mtga-server-path"
          v-model="mtgaServerPath"
          readonly
          @click="openFolderDialog('mtgaServerPath')"
        />
        <button @click="openFolderDialog('mtgaServerPath')">Browse</button>
      </div>

      <div class="setting">
        <label for="client-path">Client Path:</label>
        <input
          type="text"
          id="client-path"
          v-model="clientPath"
          readonly
          @click="openFileDialog('clientPath')"
        />
        <button @click="openFileDialog('clientPath')">Browse</button>
      </div>

      <!-- Placeholder settings -->
      <div class="setting">
        <label for="placeholder-setting-1">Placeholder Setting 1:</label>
        <input type="text" id="placeholder-setting-1" v-model="placeholderSetting1" />
      </div>

      <div class="setting">
        <label for="placeholder-setting-2">Placeholder Setting 2:</label>
        <input type="text" id="placeholder-setting-2" v-model="placeholderSetting2" />
      </div>
    </div>
  </template>

  <script>
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
    methods: {
        updatePopoutPosition() {
      // Calculate the new position based on scroll
      const scrollTop = window.scrollY || window.scrollY;
      const scrollLeft = window.scrollX || window.scrollX;

      // Update the popout's position
      this.popoutTop = 50 + scrollTop; // Adjust the top position as needed
      this.popoutLeft = 50 + scrollLeft; // Adjust the left position as needed
    },
      openFileDialog(property) {
        // Implement file dialog logic for the specified property
        // For example, you can use an HTML input element with type="file" here.
        // When a file is selected, set the corresponding data property.
      },
      openFolderDialog(property) {
        // Implement folder dialog logic for the specified property
        // You may need to use a custom solution or a third-party library for folder selection.
        // When a folder is selected, set the corresponding data property.
      },
    },
    mounted() {
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
  position: fixed; /* Change "sticky" to "fixed" */
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
