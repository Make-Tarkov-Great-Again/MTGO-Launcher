<template>
  <header>
  </header>

  <div class="container">
    <div class="floating-page" v-if="showFloatingPage">
      <button class="close-button" @click="closeFloatingPage">Close</button>
      <router-view></router-view>
    </div>
    <div class="workshop_header">
      <div class="browseAppDetails loadable">
      </div>
      <div class="workshop_browse_menu_area">
        <div class="workshop_browse_tab active">
          <a href="#">Home</a>
        </div>
        <div class="workshop_browse_tab drop ">
          <a>Installed mods</a>
        </div>
        <div class="workshop_browse_tab ">
          <a href="#">Upload mod</a>
        </div>
        <div class="workshop_browse_tab ">
          <a href="#">Upload policy</a>
        </div>
        <div class="workshop_browse_tab dropdown">
          <a href="#">About AKI</a>
          <div class="dropdown-content">
            <a href="https://discord.gg/Xn9msqQZan" target="_blank">Official AKI Discord</a>
            <a href="https://sp-tarkov.com" target="_blank">Official Website</a>
          </div>
        </div>
      </div>
    </div>
    <div class="main-container">
      <div class="mod-discovery">

        <div class="left-container"> <!--Why does everything in here move around so much when zooming in or out?-->
          <div class="sorting-bar">
            <input type="text" class="search-bar" placeholder="Search AKI...">

          </div>
          <select class="mod-type-select">
            <option value="">All Types</option>
            <option value="client">Client Mod</option>
            <option value="server">Server Mod</option>
            <option value="server">Tool</option>
          </select>
          <div class="mod-tags-container" style="margin-bottom: 30px">
            <p>Sort by tags</p>
            <div class="mod-tags" style="margin-bottom: 5px">
              <div class="versions">
                <p>Versions</p>
                <a href="#" class="pill version">3.5.0</a>
                <a href="#" class="pill version">Outdated</a>
              </div>
              <div class="type">
                <p>Type</p>
                <a href="#" class="pill">Weapons</a>
                <a href="#" class="pill">Traders</a>
                <a href="#" class="pill">QoL</a>
                <a href="#" class="pill">Music</a>
                <a href="#" class="pill">Quests</a>
                <a href="#" class="pill">Clothing</a>
                <a href="#" class="pill">Profiles</a>
                <a href="#" class="pill">Attachments</a>
                <a href="#" class="pill">Items</a>
                <a href="#" class="pill">Overhaul</a>
                <a href="#" class="pill">Port</a>
                <a href="#" class="pill" style="background-color: var(--mtga-blue);">MTGA Staff Picks</a>
                <a href="#" class="pill">Other</a>

              </div>
            </div>
          </div>
        </div>

        <div class="slideshow-container">
          <h1 style="position: relative; top:1px;">Mod highlight</h1>

          <div v-for="(item, index) in slideshowItems" :key="index" class="slideshow-item" :data-modID="item.modID"
            :style="{ display: index === currentSlideIndex ? 'block' : 'none' }">
            <img class="loadable" :src="item.imageUrl" :alt="item.title">
            <div class="slideshow-item-details">
              <div class="slideshow-item-title">{{ item.title }}</div>
              <div class="slideshow-item-creator">By {{ item.author }}</div>
            </div>
          </div>

          <button id="prevBtn" @click="prevSlide">❮</button>
          <button id="nextBtn" @click="nextSlide">❯</button>
        </div>

        <div class="mod-results">
          <br>
          Latest mods
          <div v-for="mod in mods" class="mod_entry_row" :key="mod.id" @click="ModPage(mod.id)" :data-modID="mod.id"
            style="cursor:pointer;">
            <div class="mod-entry-content">
                <img class="preview_image loadable" :src="mod.imageUrl" :alt="mod.title">
                <div class="mod_entry_title ellipsis">{{ mod.title }}</div>
                <div class="mod_author">By {{ mod.author }}</div>
                <div class="mod_tags ellipsis tag version">{{ mod.version }}</div>
                <div class="mod_tags ellipsis tag" v-for="tag in mod.tags">{{ tag }}</div>
              <div class="grab-button" style="cursor: pointer;" @click="handleGrabButtonClick(mod.id, $event)">+</div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
  <div class="loading-box" style="display: none;"></div>
</template>

<script>
import ModPageComponent from './ModPageComponent.vue';

export default {
  data() {
    return {
      components: {
      },
      showFloatingPage: false, // Initially set to false
      selectedModID: null,

      mods: [ // temp json-like mods for testing
        {
          id: 1,
          title: 'Escape from Hell',
          imageUrl: 'https://i.imgur.com/oyrIBpI.png',
          author: 'EFHDev',
          tags: ['Server', 'Client', 'Overhaul'],
          version: '3.5.0'
        },
        {
          id: 2,
          title: 'MP-43 12GA SAWED-OFF DOUBLE-BARREL SHOTGUN',
          imageUrl: 'https://hub.sp-tarkov.com/files/images/file/6e/1395.png',
          author: 'Mighty_Condor',
          tags: ['Server', 'Weapons'],
          version: '3.5.0'
        }
      ],
      slideshowItems: [ //TODO: Make this use fetch
        {
          modID: 1,
          imageUrl: 'https://i.imgur.com/oyrIBpI.png',
          title: 'Escape from Hell',
          author: 'Kestrel'
        },
        {
          modID: 2,
          imageUrl: 'https://hub.sp-tarkov.com/files/images/file/6e/1395.png',
          title: 'MP-43 12GA SAWED-OFF DOUBLE-BARREL SHOTGUN',
          author: 'Mighty_Condor'
        }
        // Add more slideshow items here as needed
      ],

      currentSlideIndex: 0
    };
  },
  methods: {
    ModPage(modID) {
      console.log(modID);
      this.showFloatingPage = true;
      this.selectedModID = modID;
      this.$router.push({ name: 'modpage', params: { id: modID } });
    },
    openFloatingPage() {
      console.log('Door open sound')
      this.showFloatingPage = true;
    },
    closeFloatingPage() {
      console.log('Door close sound')

      this.showFloatingPage = false;
    },
    prevSlide() {
      if (this.currentSlideIndex > 0) {
        this.currentSlideIndex--;
      }
    },
    nextSlide() {
      if (this.currentSlideIndex < this.slideshowItems.length - 1) {
        this.currentSlideIndex++;
      }
    },
    handleGrabButtonClick(modID, event) {
      event.stopPropagation(); // Stop event propagation here
      // Rest of your logic
    },
    handleDocumentClick(event) {
      const floatingPageElement = this.$refs.floatingPage;

      if (floatingPageElement && !floatingPageElement.contains(event.target)) {
        this.closeFloatingPage();
      }
    }
  },
  mounted() {
    document.addEventListener('click', this.handleDocumentClick);
  },
  beforeUnmount() {
    document.removeEventListener('click', this.handleDocumentClick);
  },
};
</script>


<style scoped>
.floating-page {
  color: black;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  max-width: 85vw;
  max-height: 65vh;
  background-color: var(--GithubM-dark);
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  z-index: 1000;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  font-size: 18px;
  cursor: pointer;
}

.close-button:hover {
  background-color: #d32f2f;
}
</style>