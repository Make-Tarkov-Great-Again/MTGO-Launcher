<template>
  <div class="container">
    <div class="workshop_header">
      <div class="browseAppDetails loadable"></div>
      <div class="workshop_browse_menu_area">
        <div class="workshop_browse_tab active">
          <a href="#">Home</a>
        </div>
        <div class="workshop_browse_tab drop">
          <a>Installed mods</a>
        </div>
        <div class="workshop_browse_tab">
          <a href="#">Upload mod</a>
        </div>
        <div class="workshop_browse_tab">
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
        <div class="left-container">
          <div class="sorting-bar">
            <input type="text" class="search-bar" placeholder="Search AKI...">
          </div>
          <select class="mod-type-select">
            <option value="">All Types</option>
            <option value="client">Client Mod</option>
            <option value="server">Server Mod</option>
            <option value="tool">Tool</option>
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
      </div>
      <div class="slideshow-container">
        <h1>Mod highlight</h1>
        <div v-for="(slide, index) in slideshowItems" :key="index" class="slideshow-item" :data-modID="slide.modID"
          :style="{ display: index === currentSlideIndex ? 'block' : 'none' }">
          <img class="loadable" :src="slide.imageUrl" :alt="slide.title">
          <div class="slideshow-item-details">
            <div class="slideshow-item-title">{{ slide.title }}</div>
            <div class="slideshow-item-creator">By {{ slide.author }}</div>
          </div>
        </div>
        <button id="prevBtn" @click="prevSlide">❮</button>
        <button id="nextBtn" @click="nextSlide">❯</button>
      </div>
      <div class="mod-results">
        <br>
        Latest mods
        <div class="mod-entry-row" v-for="(mod, index) in mods" :key="index" :data-modID="mod.id"
          @click="modPage(mod.id)">
          <img class="preview_image loadable" :src="mod.imageUrl" :alt="mod.title">
          <div class="mod_entry_title ellipsis">{{ mod.title }}</div>
          <div class="user_actions">
          </div>
          <div class="grab-button" style="cursor: pointer;" @click="handleGrabButtonClick($event, mod.id)">+</div>
          <div class="mod_author">By {{ mod.author }}</div>
          <div class="mod_tags ellipsis tag version">{{ mod.version }}</div>
          <div class="mod_tags ellipsis tag" v-for="tag in mod.tags" :key="tag">{{ tag }}</div>
        </div>
      </div>
    </div>
    <div class="loading-box" style="display: none;"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      mods: [ // temp json-like mods for testing
        {
          id: 1,
          title: 'Escape from Hell',
          imageUrl: 'https://i.imgur.com/oyrIBpI.png',
          author: 'EFHDev',
          tags: ['3.5.0', 'Server', 'Client', 'Overhaul']
        },
        {
          id: 2,
          title: 'MP-43 12GA SAWED-OFF DOUBLE-BARREL SHOTGUN',
          imageUrl: 'https://hub.sp-tarkov.com/files/images/file/6e/1395.png',
          author: 'Mighty_Condor',
          tags: ['3.5.0', 'Server', 'Weapons']
        }
      ]
    };
  },
  methods: {
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
    modPage(modID) {
      // lol none
    }
  }
};
</script>

<link rel="stylesheet" href="@/assets/styles/style.css">
