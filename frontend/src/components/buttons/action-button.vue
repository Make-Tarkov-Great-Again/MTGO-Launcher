<template>
    <div class="btn-main">
      <div class="button" :style="{ height: isDropdownOpen ? '100px' : '50px', overflow: isDropdownOpen ? 'hidden' : 'visible' }">
        <p class="text" v-if="!isDropdownOpen">Workshop</p>
        <div class="drop_down" @click="toggleDropdown" :style="{ height: isDropdownOpen ? '100px' : '50px' }">
              <svg class="chevrons_down" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
              width="13px" height="15px" :style="{ transform: isDropdownOpen ? 'rotate(180deg)' : 'rotate(0deg)' }">
                  <image x="0px" y="0px" width="13px" height="15px"
                      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA0AAAAPCAQAAACVKo38AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCR0UBjVzMASgAAAAVklEQVQY09XOsRGAMAxD0e8mjOJtsgxrkGWyDXOkwk0ocjk7wAKo0z0VEuBgY8fwJAoXCI1OJQWodBpAxgIOMPLYOT4g4gscP2BiAFnOK3D+5/waRb3cntNFK1LKDxgAAAAASUVORK5CYII=" />
              </svg>
        </div>
        <div v-if="isDropdownOpen" class="options">
          <p class="text" v-for="(option, index) in options" :key="index" @click="selectOption(option)">{{ option }}</p>
        </div>
      </div>
    </div>
  </template>
<script>
export default {
  data() {
    return {
      isDropdownOpen: false,
      options: ["Workshop", "Option 2", "Option 3", "Sex"], // Example options
    };
  },
  methods: {
    toggleDropdown() {
      this.isDropdownOpen = !this.isDropdownOpen;
    },
    selectOption(option) {
      // Emit an event with the selected option's value
      this.$emit("option-selected", option);
    },
  },
};
</script>
<style scoped>
.btn-main {
    position: relative;
    width: 119px;
    height: 50px;
}

.options {
  display: flow-root;
}

.options .text:hover {
background-color: rgba(10, 13, 16, 0.855);
}


.chevrons_down {
    position: relative;
    top: 50%;
    left: 50%;
    translate: -50% -50%;
    width: 13px;
    height: 15px;
    transition: transform .25s ease-in-out;
}

.text {
  cursor: pointer;
    position: relative;
    left: 50%;
    top: 50%;
    translate: -50% -50%;
    font-size: 12px;
    font-family: "Motiva Sans";
    color: rgb(255, 255, 255);
    line-height: 1.2;
    text-align: center;
    -moz-transform: matrix(1.45113815004965, 0, 0, 1.45113815004965, 0, 0);
    -webkit-transform: matrix(1.45113815004965, 0, 0, 1.45113815004965, 0, 0);
    -ms-transform: matrix(1.45113815004965, 0, 0, 1.45113815004965, 0, 0);
    z-index: 3;
}

.drop_down {
  cursor: pointer;
  position: absolute;
    top: -1px;
    border-radius: 11.5px;
    border-width: 1px;
    border-color: rgb(16, 17, 23);
    border-style: solid;
    background-color: rgb(21, 27, 34);
    width: 23px;
    left: 96px;
    height: 50px;
    z-index: 50;
    transition:height 0.25s ease-in-out;
}

.button {
    border-radius: 10px;
    background-color: rgb(21, 27, 34);
    width: 107px;
    height: 50px;
    transition:height 0.25s ease-in-out;
}
</style>