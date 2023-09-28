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

import mitt from 'mitt';
import { ref, onMounted, onBeforeUnmount } from 'vue';

import MainApp from './components/main.vue';
import eventEmitter from "./components/js/notifacation.js"

export default {
  components: {
    Topbar,
    MainApp,
  },
  setup() {
    const notifications = ref([]);

    const emitter = eventEmitter

    emitter.on('new-notification', (notification) => {
      console.log("New notification");
    });

    onMounted(() => {
      const storedNotifications = JSON.parse(localStorage.getItem('notifications'));
      if (storedNotifications) {
        notifications.value = storedNotifications;
      } else {
        notifications.value = [];
      }
    });

    const addEpNotification = () => {
      let notification = {};

      notification = {
        title: "RE: Hey",
        description: `Hey, feds are watching my texts, gotta talk through here. Wanna come to my island and party?`,
        type: "sus",
        from: "From Epstein"
      };

      emitter.emit("new-notification", notification);

      const existingNotifications = JSON.parse(localStorage.getItem('notifications')) || [];
      existingNotifications.push(notification);
      localStorage.setItem('notifications', JSON.stringify(existingNotifications));


    }


    const addNotification = () => {
      const types = ["scavTimers", "normal", "system", "modUpdate"];
      const randomType = types[Math.floor(Math.random() * types.length)];

      let notification = {};

      switch (randomType) {
        case "scavTimers":
          const randomUsername = ["Kestrel", "Nehax", "King", "Pixel", "SlejmUr", "SSH__"];
          const randomScavUsername = randomUsername[Math.floor(Math.random() * randomUsername.length)];
          notification = {
            title: "Scav Timer",
            description: `${randomScavUsername}'s scav is ready!`,
            type: "scavTimers",
            from: "From System"
          };
          break;

        case "normal":
          notification = {
            title: "Hideout production complete",
            description: "Your production for \n Bitcoin \n is complete.",
            type: "normal",
            from: "From Hideout"
          };
          break;

        case "system":
          notification = {
            title: "MTGA Updater Ready",
            description: "MTGA has a new release \n version 1.2.3, over 9000 bug fixes.",
            type: "system",
            from: "From System"
          };
          break;

        case "modUpdate":
          const randomModNames = [
            { name: "MAS", author: "SSH__" },
            { name: "KMC", author: "The_Katto" }
          ];
          const randomMod = randomModNames[Math.floor(Math.random() * randomModNames.length)];
          notification = {
            title: `Mod update for ${randomMod.name}`,
            description: `A new update was released for ${randomMod.name} by ${randomMod.author}. Click to download now.`,
            type: "mods",
            from: "From Mods"
          };
          break;

        default:
          notification = {
            title: "Unknown Notification Type",
            description: "This is an unknown notification type.",
            type: "unknown",
            from: "From System"
          };
          break;
      }

      emitter.emit("new-notification", notification);

      const existingNotifications = JSON.parse(localStorage.getItem('notifications')) || [];
      existingNotifications.push(notification);
      localStorage.setItem('notifications', JSON.stringify(existingNotifications));
    };

    const handleKeyPress = (event) => {
      if (event.key === 'r') {
        addNotification();
      } else if (event.key === 'e') {
        addEpNotification()
      }
    };

    onBeforeUnmount(() => {
      document.removeEventListener('keydown', handleKeyPress);
    });

    return {
      notifications,
      addNotification,
      handleKeyPress,
      emitter,
    };
  },
  mounted() {
    document.addEventListener('keydown', this.handleKeyPress);
  },
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
  background: transparent;

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

.notification:active {
  border: white 1px solid;
  animation: ring 1.5s infinite alternate;
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
  background-color: transparent;
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
