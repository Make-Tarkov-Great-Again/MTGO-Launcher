<template>
  <div v-on:="updatePopoutPosition" class="settings-popout-container" v-if="showSettings"
    :style="{ top: popoutTop + 'px', left: popoutLeft + 'px' }">
    <settings-popout @closeSettings="toggleSettings" />
  </div>


  <!-- Render profiles here -->
  <h1 v-if="!isLoading && selected && !selectedProfile" style="position: absolute; top: 25px;">Pick a profile nerd</h1>
  <div :class="[profileLayoutClass]" v-if="!isLoading && selected && !selectedProfileVariable">
    <template v-if="profiles.length > 8">
      <profileSelectionTable :profiles="profiles" />
    </template>
    <template v-else>
      <profileSelection class="profilesig" v-for="(profile, index) in profiles" :key="profile.id" :profile="profile"
        @click="selectProfile(profile.id)" />
    </template>
  </div>

    <component :is="selectedComponent" v-if="selectedprofile" />

  <div class="game-choices" v-if="!selectingProfile">

    <h1 class="flavor">Pick your poison</h1>

    <div class="game-choice" v-if="!selectingProfile" @click="selectGame('AKI')">
      <img src="/src/assets/images/AKI-Logo.jpg" alt="" class="game-image">
      <p class="game-name">SPT-AKI</p>
    </div>
    <div class="game-choice" v-if="!selectingProfile" @click="selectGame('MTGA')">
      <img
        src="/src/assets/images/MTGA-Logo.png"
        alt="MTGA" class="game-image">
      <p class="game-name">MTGA</p>
    </div>
  </div>
</template>

<script>
import AkiWorkshopComponent from '../components/AkiWorkshopComponent.vue';
import MtgaContentComponent from '../components/AkiWorkshopComponent.vue';
import ModPageComponent from '../components/ModPageComponent.vue';
import SettingsPopout from '../components/Settings.vue';
import profileSelection from '../components/profileSelection.vue';
import profileSelectionTable from '../components/profileselectiontable.vue';
import { GetProfiles } from "../../wailsjs/go/profile/ProfileRunT";
import { Menu as VMenu } from "floating-vue";


export default {
  components: {
    AkiWorkshopComponent,
    MtgaContentComponent,
    ModPageComponent,
    SettingsPopout,
    profileSelection,
    profileSelectionTable,
    VMenu,
  },
  data() {
    return {
      selected: '',
      isLoading: false,
      showSettings: false,
      selectingProfile: false,
      selectedProfileVariable: null,
      selectedprofile: false,
      profiles: {
        'AKI': [],
        'MTGA': [],
      },
      selectedProfileVariable: null, // Initialize selected profile variable
    };
  },
  computed: {
    selectedComponent() {
      if (this.selected === 'AKI') {
        return 'AkiWorkshopComponent';
      } else if (this.selected === 'MTGA') {
        return 'MtgaContentComponent';
      } else {
        return null;
      }
    },
    profileLayoutClass() {
      const selectedGameProfiles = this.selectedProfiles;
      return selectedGameProfiles.length > 10 ? 'profile-table' : 'profiles';
    },
    selectedProfiles() {
      if (this.selected && this.selected in this.profiles) {
        return this.profiles[this.selected];
      }
      return [];
    },
    selectedProfile() {
      if (this.selectedProfiles && this.selectedProfiles.length > 0) {
        return this.selectedProfiles[0];
      }
      return null;
    },
    profileLayoutClass() {
      return this.profiles.length > 10 ? 'profile-table' : 'profiles';
    },
  },
  methods: {
    getSelectedGame() {
      return sessionStorage.getItem('selectedGame');
    },
    toggleSettings() {
      this.showSettings = !this.showSettings;

      if (this.showSettings) {
        document.body.classList.add('settings-open');
      } else {
        document.body.classList.remove('settings-open');
      }
    },
    async selectGame(game) {
      this.selected = game;
      this.selectedProfileVariable = null;
      this.selectingProfile = true;

      // Store the selected game in sessionStorage
      sessionStorage.setItem('selectedGame', game);

    },
    selectProfile(profileID) {
      sessionStorage.setItem('selectedProfile', profileID);
        this.isLoading = true;
        this.selectedprofile = true;

    },
    async prefetchProfiles(game) {
      try {
        this.profiles[game] = await GetProfiles(game);
      } catch (error) {
        console.error('Error prefetching profiles', error);
      }
    },
  },
  async mounted() {
    const settingsIcon = document.querySelector('.settings');
    if (settingsIcon) {
      settingsIcon.addEventListener('click', () => {
        this.toggleSettings();
      });
    }
    const profiles = await GetProfiles("AKI");
    this.profiles = profiles.profiles
  },
};
</script>

<style scoped>
.settings-popout-container {
  position: absolute;
}

body.settings-open {
  overflow: hidden;
}

.profiles {
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  position: relative;
  display: flex;
}

.profile-table {}

.profilesig {
  display: inline-block;
  margin: 20px;
}

.profilesig:hover {
  border: 0.5px solid white;
}
</style>
