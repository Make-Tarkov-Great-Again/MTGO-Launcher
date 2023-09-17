<template>
    <div v-on:="updatePopoutPosition"
      class="settings-popout-container"
      v-if="showSettings"
      :style="{ top: popoutTop + 'px', left: popoutLeft + 'px' }"
    >
      <settings-popout @closeSettings="toggleSettings" />
    </div>
      <component :is="selectedComponent" v-if="selected" />

    <!-- Settings popout container -->


      <div class="game-choices" v-else>
        <h1 class="flavor">Pick your poison</h1>
        <div class="game-choice" @click="selectGame('spt-aki')">
          <img
            src="https://external-preview.redd.it/M8IFSwhfPeYfLy4mI3v74lRDacK1aVpQ3wNZY1eKQgs.jpg?auto=webp&s=0c340f58b18e758a795835741dd457bae09a3c8f"
            alt="SPT-AKI" class="game-image">
          <p class="game-name">SPT-AKI</p>
        </div>
        <div class="game-choice" @click="selectGame('mtga')">
          <img
            src="https://media.discordapp.net/attachments/964671313975312474/1083167689352159252/Logo.png?width=700&height=700"
            alt="MTGA" class="game-image">
          <p class="game-name">MTGA</p>
        </div>
      </div>
      <div class="loading-background" v-if="isLoading"></div>
      <div class="loading-box" v-if="isLoading"></div>
</template>

<script>
import AkiWorkshopComponent from './components/AkiWorkshopComponent.vue';
import MtgaContentComponent from './components/AkiWorkshopComponent.vue'; //haha aki template go brrrr
import ModPageComponent from './components/ModPageComponent.vue'; //haha aki template go brrrr
import SettingsPopout from './components/Settings.vue';


export default {
  components: {
    AkiWorkshopComponent,
    MtgaContentComponent,
    ModPageComponent,
    SettingsPopout
    //TODO: Mod pages
    // TODO: Mod-pack pages
    // TODO: Search
    // TODO: Upload policy
    // You get the gist by now kestrel ffs.
  },
  data() {
    return {
      selected: '',
      isLoading: false,
      showSettings: false,
      popoutTop: 0, // Initialize with the desired initial top position
      popoutLeft: 0, // Initialize with the desired initial left position
    };
  },
  mounted() {
  // Access the Vue instance using the ref and add the click event listener
  const settingsIcon = document.querySelector('.settings');
  if (settingsIcon) {
    settingsIcon.addEventListener('click', () => {
      this.toggleSettings();
    });
  }

  if(this.showSettings)

  // Register the scroll event listener here
  window.addEventListener('scroll', this.updatePopoutPosition);
},
  computed: {
    selectedComponent() {
      if (this.selected === 'spt-aki') {
        return 'AkiWorkshopComponent';
      } else if (this.selected === 'mtga') {
        return 'MtgaContentComponent';
      } else {
        return null;
      }
    }
  },
  methods: {
    updatePopoutPosition() {
      // Calculate the new position based on scroll
      console.log("Fuck you. I hate you.")
      const scrollTop = window.scrollY || window.scrollY;
      const scrollLeft = window.scrollX || window.scrollX;

      // Update the popout's position
      this.popoutTop = 50 + scrollTop; // Adjust the top position as needed
      this.popoutLeft = 50 + scrollLeft; // Adjust the left position as needed
    },
    toggleSettings() {
    // Toggle the visibility of settings popout
    this.showSettings = !this.showSettings;

    // Add or remove the class on the body element to disable or enable scrolling
    if (this.showSettings) {
      document.body.classList.add('settings-open');
    } else {
      document.body.classList.remove('settings-open');
    }
  },
    selectGame(game) {
      this.selected = game;
    }
  }
};
</script>
<style scoped>
.settings-popout-container {
  position: absolute;
}
body.settings-open {
    overflow: hidden;
  }

  /* Modify the styles for your settings popout */

</style>





