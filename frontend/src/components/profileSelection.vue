<template>
  <VDropdown :theme="'hover'" :placement="'auto'" :triggers="['hover']" @show="showHoverInfo" @hide="hideHoverInfo">

    <div class="profile-container" @mouseover="showHoverInfo" @mouseleave="hideHoverInfo" alt="Profile Background">

      <img class="side-logo" :src="getSideImage(profile.side)" alt="Side Logo" />
      <p class="username">{{ profile.username }}</p>
      <img class="rank-image" :src="getRankImage(profile.level)" alt="Rank Image" />
      <img class="profile-icon" src="/src/assets/images/Tarkov/faces/usec/Hudson.png" alt="Profile Icon" />
      <p class="hover level" :style="{ display: hoverInfoVisible ? 'block' : 'none' }">{{ profile.Level }}</p>




    </div>
    <template #popper>
      <hover :profile="profile" />
    </template>
  </VDropdown>
</template>



<script>
import {
  // Directives
  VTooltip,
  VClosePopper,
  // Components
  Tooltip,
  Menu
} from 'floating-vue'
import levels from './js/levels.js'
import hover from './profileHover.vue'

export default {
  props: {
    profile: {
      type: Object,
      required: true,
    },
  },
  components: {
    hover,
  },
  data() {
    return {
      computedFillStyle: '',
      progressBarText: '',
      hoverInfoVisible: false,
    };
  },
  methods: {
    showHoverInfo() {
      this.hoverInfoVisible = true;
    },
    hideHoverInfo() {
      this.hoverInfoVisible = false;
    },
    getSideImage(side) {
      if (side === 'Bear') {
        return '/src/assets/images/Tarkov/BEAR_icon.png';
      } else if (side === 'Usec') {
        return '/src/assets/images/Tarkov/USEC_icon.png';
      } else {
        return 'https://upload.wikimedia.org/wikipedia/en/d/d0/Dogecoin_Logo.png';
      }
    },
    getRankImage(level) {
      if(level === 1) {
        level = 5
            }
      const rankNumber = Math.ceil(level / 5) * 5;
      return `/src/assets/images/Tarkov/ranks/Rank${rankNumber}.png`;
    },
    getProfileIcon(icon) {
      return `/src/assets/images/Tarkov/icons/${icon}.png`;
    },
  },
  computed: {
    fillStyle() {
      const userLevel = this.profile.Level;
      const levelData = levels[0];
      const maxExperience = levelData[userLevel + 1];

      const experience = this.profile.Experience;
      const progress = (experience / maxExperience) * 100;

      this.progressBarText = `${progress.toFixed(1)}%`;

      return progress;
    },
  },
  watch: {
    profile: {
      immediate: true,
      handler(newProfile) {
        this.$nextTick(() => {
          this.computedFillStyle = this.fillStyle;
        });
      },
    },
  },
  mounted() {

    const maxExperience = levels[0]['80'];
    const nextLevel = Math.ceil(this.profile.Level / 5) * 5 + 5;
    const experience = this.profile.Experience;
    const progress = (experience / maxExperience) * 100;
    const remainingExperience = Math.max(0, levels[0][nextLevel] - experience);

    this.computedFillStyle = (1 - remainingExperience / maxExperience) * 100 + '%';
  }
};
</script>

<style scoped>



.profile-container {
  position: relative;
  background-color: #010205;
  left: 0px;
  top: 0px;
  width: 153px;
  height: 192px;
  z-index: 7;
  border-radius: 5px;
}

.side-logo {
  position: absolute;
  left: 117px;
  top: 10px;
  width: 25px;
  height: 30px;
  z-index: 6;
}

.username {
  font-size: 12px;
  color: rgb(255, 255, 255);
  line-height: 1.2;
  text-align: center;
  -moz-transform: matrix(1.45162454511416, 0, 0, 1.45162454511416, 0, 0);
  -webkit-transform: matrix(1.45162454511416, 0, 0, 1.45162454511416, 0, 0);
  -ms-transform: matrix(1.45162454511416, 0, 0, 1.45162454511416, 0, 0);
  position: relative;
  top: 175px;
  z-index: 5;
  user-select: none;
}

.rank-image {
  position: absolute;
  left: 11px;
  top: 9px;
  width: 34px;
  height: 31px;
  z-index: 4;
}

.profile-icon {
  position: absolute;
  left: 11px;
  top: 8px;
  width: 133px;
  height: 166px;
  z-index: 3;
}

.hover {
  position: absolute;
}

.hover.level {
  color: white;
  top: 160px;
  left: 15px;
}

.hover.progress-barr {
  color: white;
  background-color: #010205;
  width: 10px;
  top: 189px;
  left: 13px;
  height: 10px;
}

.progress-fill {
  background-color: #4CAF50;
  height: 100%;
  transition: width 0.3s ease-in-out;
  width: 130px;
}


.progress-text {
  position: absolute;
  top: 6px;
  left: 55px;
  color: #fff;
  font-size: x-small;
}
</style>
