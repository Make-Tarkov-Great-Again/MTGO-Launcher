<template>
      <component :is="selectedComponent" v-if="selected" />


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


export default {
  components: {
    AkiWorkshopComponent,
    MtgaContentComponent,
    ModPageComponent
    //TODO: Mod pages
    // TODO: Mod-pack pages
    // TODO: Search
    // TODO: Upload policy
    // You get the gist by now kestrel ffs.
  },
  data() {
    return {
      selected: '',
      isLoading: false
    };
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
    selectGame(game) {
      this.selected = game;
    }
  }
};
</script>
<link rel="stylesheet" href="@/assets/styles/style.css">





