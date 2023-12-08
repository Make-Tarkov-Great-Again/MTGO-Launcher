<template>
  <div v-on:="updatePopoutPosition" class="settings-popout-container" v-if="showSettings"
    :style="{ top: popoutTop + 'px', left: popoutLeft + 'px' }">
    <settings-popout @closeSettings="toggleSettings" />
  </div>


  <h1 v-if="!isLoading && selected && !selectedProfile"
    style="position: absolute; top: 25px; left: 50%; translate: -50%;">Pick a profile nerd</h1>
  <div :class="[profileLayoutClass]" v-if="!isLoading && selected && !selectedProfileVariable">
    <template v-if="profiles.length > 8">
      <profileSelectionTable :profiles="profiles" />
    </template>
    <template v-else-if="profiles.length >= 1">
      <profileSelection class="profilesig" v-for="(profile, index) in profiles" :key="profile.id" :profile="profile"
        @click="selectProfile(profile.id)" />
    </template>
    <template v-else>
      <p>You dont have any profiles! Open the settings and set your AKI server path!</p>
      <svg class="introduction-arrow" data-name="1-Arrow Up" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32">
        <path d="m26.71 10.29-10-10a1 1 0 0 0-1.41 0l-10 10 1.41 1.41L15 3.41V32h2V3.41l8.29 8.29z" />
      </svg>
    </template>
  </div>

  <FirtScreen v-if="selectedprofile" />

  <div class="game-choices" v-if="!selectingProfile">

    <h1 class="flavor">Pick your poison</h1>
    <div class="game-choice" v-if="!selectingProfile" @click="selectGame('MTGA')">
      <img src="/src/assets/images/MTGA-Logo.png" alt="MTGA" class="game-image">
      <p class="game-name">MTGA</p>
    </div>
    <div class="game-choice" v-if="!selectingProfile" @click="selectGame('AKI')">
      <img src="/src/assets/images/Fork-Logo.png" alt="" class="game-image">
      <p class="game-name">MTGA-Forks</p>
    </div>
  </div>
</template>

<script>
import AkiWorkshopComponent from '../components/Workshop/AkiWorkshopComponent.vue';
import MtgaContentComponent from '../components/Workshop/AkiWorkshopComponent.vue';
import ModPageComponent from '../components/ModPageComponent.vue';
import SettingsPopout from '../components/Settings.vue';
import profileSelection from '../components/profileSelection.vue';
import profileSelectionTable from '../components/profileselectiontable.vue';
import { GetProfiles } from "../../wailsjs/go/profile/ProfileRunT";
import { Menu as VMenu } from "floating-vue";
import mainpage from './MainPage/mainpage.vue';
import FirtScreen from './Startup/FirtScreen.vue';

export default {
  components: {
    AkiWorkshopComponent,
    MtgaContentComponent,
    ModPageComponent,
    SettingsPopout,
    profileSelection,
    profileSelectionTable,
    VMenu,
    mainpage,
    FirtScreen
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
    };
  },
  computed: {
    selectedComponent() {
      return this.selected === 'AKI' ? 'AkiWorkshopComponent' : this.selected === 'MTGA' ? 'MtgaContentComponent' : null;
    },
    profileLayoutClass() {
      return this.selectedProfiles.length > 10 ? 'profile-table' : 'profiles';
    },
    selectedProfiles() {
      return this.selected && this.selected in this.profiles ? this.profiles[this.selected] : [];
    },
    selectedProfile() {
      return this.selectedProfiles.length > 0 ? this.selectedProfiles[0] : null;
    },
  },
  methods: {
    getSelectedGame() {
      return sessionStorage.getItem('selectedGame');
    },
    toggleSettings() {
      this.showSettings = !this.showSettings;
      document.body.classList.toggle('settings-open', this.showSettings);
    },
    async selectGame(game) {
      this.selected = game;
      this.selectedProfileVariable = null;
      this.selectingProfile = true;
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
      settingsIcon.addEventListener('click', this.toggleSettings);
    }
    const profiles = await GetProfiles("AKI");
    this.profiles = profiles.profiles;
  },
};
</script>

<style scoped>
.game-choices {
  position: relative;
  display: inline-block;
  left: 50%;
  top: 50%;
  translate: -50% -50%;
}

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
  left: 50%;
  top: 50%;
  translate: -50% -50%;
}

.introduction-arrow {
  filter: invert(100%) sepia(100%) saturate(0%) hue-rotate(230deg) brightness(106%) contrast(103%);
    position: absolute;
    color: white;
    height: 50px;
    width: 50px;
    right: -10px;
    top: -265px;
    rotate: 15deg;
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
