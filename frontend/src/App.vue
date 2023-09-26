<template>
  <div class="systemApp">
    <Topbar></Topbar>
</div>
  <div class="outerApp2">
    <div class="main-container" id="app">
      <MainApp></MainApp>
    </div>
  </div>
</template>


<script>
import {
  // Directives
  VTooltip,
  VClosePopper,
  // Components
  Dropdown,
  Tooltip,
} from 'floating-vue'
import { Menu as VMenu } from "floating-vue";
import FloatingVue from 'floating-vue'
import Topbar from './components/System/cus-Topbar.vue';




import MainApp from './components/main.vue';
export default {
  components: {
    Topbar,
    MainApp,
  }
};
</script>


<style scoped>


body {
  background-color: rgb(255, 1, 1) !important;
}
.outerApp2 {
  width: 918px;
  height: 596px;
  overflow: scroll;
  color: dimgray;
  border-radius: 15px;
  border: red solid;
}

.systemApp {
  width: 100vw;
  height: 100vh;
  position: absolute;
  overflow: hidden;
  background: transparent;
}

.notification {
  position: absolute;
  left: 909px;
  top: 5px;
  width: 21px;
  height: 24px;
  z-index: 9;
  border-radius: 120px;
  animation: ring 1s infinite alternate;

}

.notification:hover {
  border: white 1px solid;
  background-color: white;
  animation: ring 1s infinite alternate;

}

@keyframes ring {
  0% {
    transform: rotate(0deg);
  }

  50% {
    transform: rotate(15deg);
  }

  100% {
    transform: rotate(-15deg);
  }
}


#app {
  background-color: #0d1117;
  position: absolute;
  width: 90%;
  height: 90%;
  overflow: auto;
}

.mainapp {
  background-color: black;
  z-index: 9;
}

.home {
  position: fixed;
  left: 32px;
  top: 12px;
  width: 19px;
  height: 24px;
  z-index: 8;
}

.downloads {
  position: absolute;
  left: 878px;
  top: 5px;
  width: 23px;
  height: 24px;
  z-index: 4;
}


.minimize {
  cursor: pointer;
  position: fixed;
  left: 975px;
  top: 16px;
  width: 11px;
  height: 2px;
  z-index: 7;
}

.close {
  position: fixed;
  left: 1000px;
  top: 10px;
  width: 14px;
  height: 14px;
  z-index: 6;
}

.branding {
  position: absolute;
  left: 50%;
  translate: -50%;
  top: 9px;
  width: 147px;
  height: 16px;
  z-index: 9;
  --wails-draggable: drag
}


.settings {
  position: absolute;
  left: 940px;
  top: 6px;
  width: 22px;
  height: 22px;
  z-index: 8;
}
</style>
