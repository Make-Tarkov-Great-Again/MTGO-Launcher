<template>
  <div class="single-mod-page">
    <!-- Loop through each mod -->
    <!-- Grab Button -->

    <svg class="close-button" @click="closeFloatingPage" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="14px" height="14px">
      <image x="0px" y="0px" width="14px" height="14px"
        xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAQAAAC1QeVaAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCQ8SNjQl7OFjAAAA4klEQVQY01XQMUrDABTG8V9NMYMFC+JgDtALZNKi4qiiIlqbgngCFyf1HoKCSO0tTK+gS0cLghUcHcyoFBzS1OYtD96fj+/Pq8R1V2rOlWde10vVtkv8uiihnsROEA292dW0KJ2gUE8ikwTR2MDI3hSHHiUyLWkQwcDIvjU1z+51ZE70qcRFz5kHVUMN3xJPkCfz9KcDS34cFe1zM4Ybk71VnIpkqOfUlzuxzcI8h4Vh4tbHv3kQEerqyLT0Z8zr0iBquNGWaU+fkONVy0HUdG3sOJef4neHViox6xZKKJ+W1z8KUj3KY02Q2QAAAABJRU5ErkJggg==" />
    </svg>

    <div class="grab_button"></div>

    <!-- Pill 1 -->
    <div class="pill rounded-rectangle">Pill 1</div>

    <!-- Pill 2 -->
    <div class="pill rounded-rectangle">Pill 2</div>

    <!-- Pill 3 -->
    <div class="pill rounded-rectangle">Pill 3</div>

    <!-- Mod Title -->
    <div class="mod_title">{{ mod.title }}</div>

    <!-- Mod Image -->
    <img class="mod_image" :src="mod.images[0]" alt="Mod Image">

    <!-- Mod Author -->
    <div class="mod_author">{{ mod.author }}</div>

    <!-- Mod Description -->
    <div class="mod_description">{{ mod.description }}</div>

    <!-- Mod Big Image -->
    <img class="mod_big_image" :src="mod.images[1]" alt="Big Mod Image">
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
  background-color: var(--github-black);
  /* Add height to ensure the background is visible */
  width: 903px;
  height: 592px;
  position: relative; /* Use relative positioning */
  z-index: 11;
  color: white; /* Add text color for better visibility */
}


.close-button {
  background-image: url("float-close.png");
  position: absolute;
  left: 874px;
  top: 99px;
  width: 14px;
  height: 14px;
  z-index: 15;
}
.grab_button {
  position: absolute;
  left: 780px;
  top: 415px;
  width: 58px;
  height: 58px;
  z-index: 14;
}
.pill {
  position: absolute;
  left: 226px;
  top: 327px;
  width: 64px;
  height: 30px;
  z-index: 13;
}
.pill {
  position: absolute;
  left: 149px;
  top: 327px;
  width: 64px;
  height: 30px;
  z-index: 12;
}
.pill {
  position: absolute;
  left: 201px;
  top: 204px;
  width: 52px;
  height: 25px;
  z-index: 11;
}
.pill {
  position: absolute;
  left: 72px;
  top: 327px;
  width: 64px;
  height: 30px;
  z-index: 10;
}
.mod_title {
  position: absolute;
  left: 201px;
  top: 127px;
  width: 248px;
  height: 39px;
  z-index: 9;
}
.mod_image {
  position: absolute;
  left: 71px;
  top: 119px;
  width: 112px;
  height: 111px;
  z-index: 8;
}
.mod_author {
  position: absolute;
  left: 202px;
  top: 169px;
  width: 126px;
  height: 27px;
  z-index: 7;
}
.mod_description {
  position: absolute;
  left: 57px;
  top: 409px;
  width: 672px;
  height: 163px;
  z-index: 6;
}
.mod_big_image {
  position: absolute;
  left: 57px;
  top: 91px;
  width: 845px;
  height: 227px;
  z-index: 5;
}
.floating_page_container {
  border-radius: 10px;
  background-color: rgb(13, 17, 23);
  position: absolute;
  left: 27px;
  top: 35px;
  width: 903px;
  height: 592px;
  z-index: 4;
}

</style>
