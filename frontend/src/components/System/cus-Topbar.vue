<template>
  <div class="settings-popout-container" v-if="showSettings">
  </div>
  <svg class="home" onclick="">
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
    width="19px" height="24px">
    <image x="0px" y="0px" width="19px" height="24px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAABMAAAAYCAQAAACyng6EAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCBUOMRhPjd8YAAABOklEQVQoz63SQUtUURyG8d8591zHUSYZiZGZdZIVgjphLaRW7YL2QZtADCRs0bJv0WfoK/Q1wl27XEQqTQbChHWdbhuT69StWfisznl5zp/3cE7oG6Nn3tBHRTXMeuPSY48sGdivxqmynnHTQ88sKHW88c7R36bd88oTVwQsW3Zsd1xr2fbcuqaAIOpYdM0HX35rmdu2bLolr1SIOla1DR0qsl7Lupc2zfmTZMUNnw2y3lM7Nkyp46o1KdnQN6OehuvuR/NO/Zvv5qIffv5HKxXRREyopQu7wldDI0FTW7NOO/DaW0ca7nrhTv20T94boetbfbdoVhvTWhcHXO5NJ9ZS5dnz82Ox0m1KHqXKL8tlAgiVNEqh/0BXQ4ng2K49J5IFq7pKcGoQ+rOSeKaNnCiUgsy0/EwrjX4BTS4/5jIoxv4AAAAASUVORK5CYII=" />
  </svg>
  <svg class="minimize" onclick="window.runtime.WindowMinimise()">
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
    width="11px" height="2px">
    <image x="0px" y="0px" width="11px" height="2px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAsAAAACCAQAAAAkqu5lAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCBUOMQUsi7PBAAAAGUlEQVQI12M0fskgyoAO/rIwPGT4hCH8HwByTwUGWZVcMwAAAABJRU5ErkJggg==" />
  </svg>

  <svg class="downloads" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="23px"
    height="24px">
    <image x="0px" y="0px" width="23px" height="24px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAABcAAAAYCAQAAAC7da7+AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCRYNFw3MHYDQAAABoUlEQVQ4y73SvWtTYRiG8V9yzmlqP20VxEODHwUJRRSp6KLg6CIiSLsUOjuIoIOD7STi7CC6uDqIm0TwP3CvFato6SCCZtCmMWmT0+OQtiZtLZm8pxee633fB+4rM25HIhdMW/DM9+2j7E7asAnXXHdy52g3vEfBgBEHO8MTK1hT6QxPJUilneF75H/hGTlduzJ9ou146JLHZhwCEbKy6DLhuZt6NrFmul0xrWbEbctqSK1i0gOx496Yb329YVFVnwmzIp+w4oeLHsrLWlJqYkG8iS8pO6dfwajDjqo4YMqYda/c8rUd57d3qsYNKcjL6XHKEYmiWR83Kwvi/UbU1FHzXs0Z/XKyIpHEa/fMSxEZFQfxU5f1W7C2cWHVab2gruiuD1KEzptxNYhfKoi99Q3UzEmc1S1RdMfnjTV6TbrhWCiwLtS9VcqyJ1JT5ty3uKVZzqBAGiJVbpP1p0deKCu1OLlm+W9NdY222qu+bBOhYUUiGyJjyJiywZZxKtNyrhtwQiDNjKeoKqkI9lRx2FBzd/bJdyZwqOEXLV//OznVPxopd5X3cz0NAAAAAElFTkSuQmCC" />
  </svg>

  <div class="download-progress" @mouseenter="showProgress = true" @mouseleave="showProgress = false">
    <label for="progress" class="progress-label" style="position: fixed; bottom: 20px; left: 50%; translate: -50%;">No
      downloads</label>
    <progress class="progress-bar" :style="{ visibility: showProgress ? 'visible' : 'hidden' }"
      style="position: fixed; bottom: 5px; left: 50%; translate: -50%; width: 330px;" id="file" value="0"
      max="100"></progress>
  </div>

  <svg class="close" onclick="window.runtime.Quit()">
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
    width="14px" height="14px">
    <image x="0px" y="0px" width="14px" height="14px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAQAAAC1QeVaAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCBUOMBxR+ypAAAAA4klEQVQY01XQMUrDABTG8V9NMYMFC+JgDtALZNKi4qiiIlqbgngCFyf1HoKCSO0tTK+gS0cLghUcHcyoFBzS1OYtD96fj+/Pq8R1V2rOlWde10vVtkv8uiihnsROEA292dW0KJ2gUE8ikwTR2MDI3hSHHiUyLWkQwcDIvjU1z+51ZE70qcRFz5kHVUMN3xJPkCfz9KcDS34cFe1zM4Ybk71VnIpkqOfUlzuxzcI8h4Vh4tbHv3kQEerqyLT0Z8zr0iBquNGWaU+fkONVy0HUdG3sOJef4neHViox6xZKKJ+W1z8KUj3KY02Q2QAAAABJRU5ErkJggg==" />
  </svg>



  <svg class="settings" @click="toggleSettings" xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink" width="22px" height="22px">
    <image x="0px" y="0px" width="22px" height="22px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAABYAAAAWCAQAAABuvaSwAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCQ8RDwufSP+9AAAB5UlEQVQoz3XTX2jNcRgG8M9vzizUURTzZzTtwkS0sxukbBfEDY5cTGozSYTJ8rckLWK5UMrVcKRIzG6sJeXPLgwzF0qKHdxQyMUyZps5LvY7/yzP1fvteb5Pz/t+328QMwZHNIjoUG84nxg3819piTvueaHOd4/zqYJMdUKbItQibrukHeC8duNHJZFQelajAX2eW+42aJKQNKjcL73KDKWdj2q00ETt5lpjI7iiTFKhQhOk9Gadi/GKUJZG0qqw6lWdzbzXM28FIVXnni6nRcNzk2qzcht8owwEelxWjEP6VIE5+JQWL3HQFtuk0KVCiUWWCtxw3/RwPg/VEMR2uoA2caz0QKkPmdR9nliNrU4p1hnE3huwICQTNpic0+JJRzOd1LhWIOJHhgzk409OPUiBFpW+aAZXReW+/y4PQ9d3Wj0JYqy3Tp3NrqNHhem+hKFqzfYRKU+1SATh1l1UL2JEoZfm69ZvmSJrdeCSraMB03OOSgkwrNxuv03VYpoO8A1TYNS5wTlR3/0Pj1SalHYuCu8e121FRjJDm55wx/uzi9SsxAevlUvqdMsmxLX6ZMSAn4bMy93nPYZUKZN0IBzjMV/NQkKp1X5lM+dims9u+mifMw7nU2P/4A+D4ha7a7+hfOovMDmCqAPHbRoAAAAASUVORK5CYII=" />
  </svg>

  <VMenu :theme="'notifications'" :placement="'auto'" :hideTriggers="['click']" @click="show = true" :triggers="['click']"
    :popperTriggers="['click']" :delay="{ show: 0, hide: 0 }" :force-show="show">
    <div class="notification-container">
      <img class="notification"
        src="https://cdn.discordapp.com/attachments/1116718992808226837/1155182924329660416/untitled_5.png" alt="">
    </div>
    <template #popper>
      <KeepAlive>
        <notification></notification>
      </KeepAlive>
    </template>
  </VMenu>
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
import profileHover from '../profileHover.vue';
const containerElement = ".notification-container"
import notification from '../System/notification/notificationBox.vue'
import eventEmitter from ".././js/notifacation.js"


FloatingVue.options.container = containerElement


export default {
  components: {
    VMenu,
    Dropdown,
    profileHover,
    notification,
  },
  data() {
    return {
      show: false,
      showSettings: false,
      downloading: false,
      count: 0,
      showProgress: false
    };
  },
  methods: {
    toggleSettings() {
      this.showSettings = !this.showSettings;
    }
  },
  mounted() {
    eventEmitter.on('new-notification', (notification) => {
  const notificationIcon = document.querySelector(".notification");
  const notificationIconActive = document.querySelector(".notification.active");



  if (!notificationIcon.classList.contains("active")) {
    notificationIcon.classList.add("active");
  }

  notificationIcon.addEventListener("click", function clickHandler() {
    notificationIcon.classList.remove("active");
      notificationIcon.removeEventListener("click", clickHandler);
    });
});



    var downloading = false
    var count = 0

    const downloadProgress = document.querySelector('.download-progress');
    const progressBar = document.querySelector('.progress-bar');
    const label = document.querySelector('.progress-label')

    downloadProgress.addEventListener('mouseenter', () => {
      downloadProgress.style.bottom = '5px';
      label.style.opacity = "100%"
    });

    downloadProgress.addEventListener('mouseleave', () => {
      if (!downloading) {
        downloadProgress.style.bottom = '-24px';
      } else { downloadProgress.style.bottom = '-21px' }
      label.style.opacity = "15%"
    });

    downloadProgress.addEventListener('mousemove', (e) => {
      const mouseY = e.clientY;
      const progressRect = progressBar.getBoundingClientRect();
      const progressTop = progressRect.top;
      const progressBottom = progressRect.bottom;
    });

    const socket = new WebSocket("ws://localhost:42145/ws");

    socket.onopen = () => {
      // WebSocket connection is open
      console.log("WebSocket is running.");
    };
    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data && data.percentage !== undefined) {
        if (data.percentage > 1) {
          downloadProgress.style.bottom = "-21px"
        }
        const progressBar = document.querySelector('.progress-bar');
        progressBar.value = data.percentage;

        const label = document.querySelector('.progress-label');
        const convertedsize = data.filesize / 1000000
        if (data.percentage < 100) {
          label.textContent = `Downloading ${data.filename} (${convertedsize} MB)`;
        }
        else {
          count++
          if (count = 1) {
            label.textContent = `${count} mod downloaded`;
          } else {
            label.textContent = `${count} mods downloaded`;
          }
        }
      }
    };


    socket.onclose = () => {
      console.log("WebSocket is not running.");
    };


  },
};
</script>

<style>
.v-popper__popper.v-popper--theme-notifications.v-popper__popper--hidden.v-popper__popper--hide-to,
.v-popper__popper.v-popper--theme-notifications.v-popper__popper--shown.v-popper__popper--show-to {
  left: 300px;
  top: 32px;
  z-index: 999999999;
}

.v-popper__popper.v-popper--theme-notifications.v-popper__popper--shown.v-popper__popper--show-to .v-popper__arrow-container {
  left: 307px !important;
  top: 15px;
  left: 195px;
  overflow: visible;
  top: 0;
}


.notification {
  position: absolute;
  left: 909px;
  top: 5px;
  width: 21px;
  height: 24px;
  z-index: 9;
  border-radius: 120px;
}

.notification.active {
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

.main_container {
  border-radius: 10px;
  background-color: rgb(13, 17, 23);
  position: absolute;
  left: 27px;
  top: 70px;
  width: 903px;
  height: 592px;
  z-index: 4;
}
</style>