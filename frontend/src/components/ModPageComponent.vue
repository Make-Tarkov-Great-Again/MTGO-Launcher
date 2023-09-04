<template>
    <div class="single-mod-page">
      <div class="mod-page-container">
        <div class="mod-page-main-container">
          <div class="mod-name">{{ mod.title }}</div>



          <div class="slideshow-containerMp">
          <div v-for="(image, index) in mod.images" :key="index" class="slideshow-itemMp"
            :style="{ display: index === currentSlideIndexMp ? 'block' : 'none' }">
            <img class="loadable" :src="image" :alt="mod.title">
            </div>
          </div>

          <button id="prevBtn" @click="prevSlideMp">❮</button>
          <button id="nextBtn" @click="nextSlideMp">❯</button>
        </div>

          <div class="mod-name-subscribe">
            <div class="mod-title">{{ mod.title }}</div>
            <div class="grab-buttonmp" @click="handleGrabButtonClick(mod.id)">Grab</div>
          </div>

          <div class="mod-description">{{ mod.description }}</div>
        </div>
      </div>
</template>

<script>
import modexamples from './mods.json'; // TODO: Replace with API fetch

export default {
  computed: {
    modID() {
      return parseInt(this.$route.params.id);
    },
    mod() {
      return modexamples.find(mod => mod.id === this.modID);
    },
  },
  data() {
    return {
      currentSlideIndexMp: 0,
    };
  },
  methods: {
    handleGrabButtonClick() {
      // Handle subscription using this.modID
    },
    prevSlideMp() {
      if (this.currentSlideIndexMp > 0) {
        this.currentSlideIndexMp--;
      }
    },
    nextSlideMp() {
      if (this.currentSlideIndexMp < this.mod.images.length - 1) {
        this.currentSlideIndexMp++;
      }
    },
  },
  created() {
    // Fetch mod data using the mod ID
    // For now, use switch statement for mock data
  },
};
</script>

<style scoped>
.single-mod-page {
  background-color: var(--GithubM-dark);
  color: var(--text-color);
}

.mod-page-container {
    background-color: var(--GithubM-dark);
}

.mod-page-main-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    align-content: flex-start;
    flex-wrap: wrap;
}

.mod-name {
  font-size: 24px;
  margin-bottom: 10px;
}

.mod-picture-gallery {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.gallery-image {
  flex-basis: calc(33.33% - 20px);
  overflow: hidden;
  border-radius: 8px;
}

.gallery-img {
  width: 100%;
  height: auto;
}

.mod-name-subscribe {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
  margin-bottom: 20px;
}

.grab-buttonmp {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: var(--mtga-blue);
  color: white;
  font-size: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}

.mod-description {
  font-size: 16px;
  line-height: 1.5;
}

.slideshow-containerMp {
  position: relative;
  max-width: 630px;
  height: 400px;
  overflow: hidden;
  background-color: var(--GithubA-dark);
  border: #21262d 1px;
  padding: 5px;
  left: 35px
}

.slideshow-itemMp {
  display: none;
  width: 100%;
  height: 100%
}

.slideshow-itemMp img {
  object-fit: contain;
  width: 100%;
  height: 35%
}

.slideshow-itemMp-details {
  background-color: rgb(0 0 0 / 40%);
  color: #fff;
  padding: 10px;
  position: absolute;
  bottom: 0;
  width: 100%
}

#nextBtn,
#prevBtn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  font-size: 24px;
  background: 0 0;
  border: none;
  color: #fff;
  cursor: pointer;
  z-index: 1
}

#prevBtn {
  left: 10px
}

#nextBtn {
  right: 10px
}
</style>
