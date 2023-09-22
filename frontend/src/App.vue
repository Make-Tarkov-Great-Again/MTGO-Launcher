<template>
  <div class="settings-popout-container" v-if="showSettings">
    <settings-popout></settings-popout>
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



  <div class="download-progress">
    <label for="progress" class="progress-label" style="position: fixed; bottom: 25px; left: 50%; translate: -50%;">No
      downloads</label>
    <progress class="progress-bar" style="position: fixed; bottom: 5px; left: 50%; translate: -50%; width: 330px;"
      id="file" value="0" max="100"></progress>
  </div>
  <svg class="close" onclick="window.runtime.Quit()">
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
    width="14px" height="14px">
    <image x="0px" y="0px" width="14px" height="14px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAQAAAC1QeVaAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCBUOMBxR+ypAAAAA4klEQVQY01XQMUrDABTG8V9NMYMFC+JgDtALZNKi4qiiIlqbgngCFyf1HoKCSO0tTK+gS0cLghUcHcyoFBzS1OYtD96fj+/Pq8R1V2rOlWde10vVtkv8uiihnsROEA292dW0KJ2gUE8ikwTR2MDI3hSHHiUyLWkQwcDIvjU1z+51ZE70qcRFz5kHVUMN3xJPkCfz9KcDS34cFe1zM4Ybk71VnIpkqOfUlzuxzcI8h4Vh4tbHv3kQEerqyLT0Z8zr0iBquNGWaU+fkONVy0HUdG3sOJef4neHViox6xZKKJ+W1z8KUj3KY02Q2QAAAABJRU5ErkJggg==" />
  </svg>

  <svg class="branding" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="147px"
    height="16px">
    <image x="0px" y="0px" width="147px" height="16px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJMAAAAQCAQAAAC4yIXoAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCQ8RECdgyp3AAAAFBklEQVRYw73YbYxUVxkH8N+9M8vuskC7dGltC0ql21aqYNJA6lv7BV9SU9OXNEFiTY0aNY2QaqxNaLVpQmtNU0ATo5HYJhKibRpDibGFpBVMlKBgEbCUStkFumXWHVjY3e7s7Nx7/bCXYWZ2Bhgi/OfD3POc/zn3PP/znOc8M8Eti2T9TT180b+9Z6npYsQiLaDFNrt0WKZVklr+7I3yuE9apAQCiU0Oa4ROc/3TpcFVuuy70MGhO6zTWadngdVulzghLy9vhrsNyMsb8L7pntCVtvOOW+H2dNy97k+teXlFj1vQ8O1zLb1EItHhAdMudHBW3iu+am2Nvd1SL4iN2ZhaPmSK35f7H7fXcxX8rR512CFLLPGIkxU9O6yw0rG6b48ULplM7whcafhCZWr1R4sstqPKfqeD8mZWWKZpLT9P1W5LFb/XPtc75EavVonEbgfMaSBTI1ymQ4LQydSxUIuxCkZGRlGgzahWXWIEhp0qM65IU0JoOF1RrKTT1JQ7oFjmXp1+n7ZmUdKpXSJxjKzAuN/5rl1pPoEut/mBLwkaOhKZMmm3chiXmcSNRU1I1OY+C9JZEpdZbR8+7D5PVbAWWuyXslbY4CGBGIlOP0tz5N3ucEqAWKdf2YGiu9xsSCjRqmSV/yL0kG4jAmQknpLTbbE9HjSINr3WZNOHHvdWHKmv+5NR2aYiYFPDnqCJWYi86TUDQkQ+arkn9WpzbRWrwwdQMtUvrLUtZX/ECqsc9Gmf82MDMojM8z2DDrjGzZbLpVvwBSt9X+TbpnhYUYDEZ/3QSj3W2u9JRwWm+IYHwvS1691mVvr8cTNtbsq1iZ1Pmh5TD+P+4V1jRo0q2mW/uUhqInKinXjXK7aU2bvtdh3m2aJPMbW+ab8bMOwxR1LrqD9o144PetkpBaNGFWxScoVRPdY5aEzBKc/aeDpiBm12v2cRWua3xs/brS6zxQIjehuOikyzpMrRUH+DMmTy2PgsvSNGqtqxGKWaCH7PcRRquBOrLVadmsCYQGjAiQpr7gxlo2cs8C/36LOnid2/3l2KQt2e8fcGnMRUN1SJmJXlvOLv7JygTi6cjA0pN6zTl1Sta64rDSGs5lYq+byv6PcpTzchEtttB9/S0ZDTot9P6zpZD5/xYEXkzU7nv1josMyBtHAO3WKDwUnXU5VMe/X5tXVNXt5npGi8781krYW+5hG95fZybRdVpki/Y2lMhv5ifz1S9W22Xq+XL+qizo0b/VVPRfv/czE0RsGr3joXqVqm4+Wa+1xomZSuJ05zvSIiUzcnNEJUVUaeFiqskevslVhYUQOeC5nzyW7NOHAGo/LurLLc5GMOYa/PlwuLCdzqOkcazBPUuRuL5lW05vuydrxvVpo/oM13TD/L+sZ0n7efGXNq1lQnY2aFDcu/6mxfeU8k1njUT/xHBolpFlqtB9tkPe0NhbRcm+Emq+QavKHkVt+smLfFHlt9whOOCsRmmO0l3bZ4x1Zr7BVLZMy3y+W6DNTeSKk3mz3mR46lVXiXjJ87KVPj6USUP+9h89MqPDFHj9+Ia7mZa4b0GqrrxIgj8uXWuJy+Cgd3mCUQicRiL5X/SDkkZ6aSSCRR8FxFOp686xM/ZU5/Ajlv2+5qgZJYyYteNyyHnVq1GRdjp/VOGjBmxFGDFTMOO2JIwU5XCZRE6LdVn8QJvVUH+oTDInlvuda4SIS3bTeY9p35zed/3JS00OIZEmAAAAAASUVORK5CYII=" />
  </svg>

  <svg class="settings" @click="toggleSettings" xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink" width="22px" height="22px">
    <image x="0px" y="0px" width="22px" height="22px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAABYAAAAWCAQAAABuvaSwAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCQ8RDwufSP+9AAAB5UlEQVQoz3XTX2jNcRgG8M9vzizUURTzZzTtwkS0sxukbBfEDY5cTGozSYTJ8rckLWK5UMrVcKRIzG6sJeXPLgwzF0qKHdxQyMUyZps5LvY7/yzP1fvteb5Pz/t+328QMwZHNIjoUG84nxg3819piTvueaHOd4/zqYJMdUKbItQibrukHeC8duNHJZFQelajAX2eW+42aJKQNKjcL73KDKWdj2q00ETt5lpjI7iiTFKhQhOk9Gadi/GKUJZG0qqw6lWdzbzXM28FIVXnni6nRcNzk2qzcht8owwEelxWjEP6VIE5+JQWL3HQFtuk0KVCiUWWCtxw3/RwPg/VEMR2uoA2caz0QKkPmdR9nliNrU4p1hnE3huwICQTNpic0+JJRzOd1LhWIOJHhgzk409OPUiBFpW+aAZXReW+/y4PQ9d3Wj0JYqy3Tp3NrqNHhem+hKFqzfYRKU+1SATh1l1UL2JEoZfm69ZvmSJrdeCSraMB03OOSgkwrNxuv03VYpoO8A1TYNS5wTlR3/0Pj1SalHYuCu8e121FRjJDm55wx/uzi9SsxAevlUvqdMsmxLX6ZMSAn4bMy93nPYZUKZN0IBzjMV/NQkKp1X5lM+dims9u+mifMw7nU2P/4A+D4ha7a7+hfOovMDmCqAPHbRoAAAAASUVORK5CYII=" />
  </svg>

  <svg class="notification" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="21px"
    height="24px">
    <image x="0px" y="0px" width="21px" height="24px"
      xlink:href="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAABUAAAAYCAQAAAC/gH7DAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfnCRYJJABBzjGBAAABLElEQVQ4y42TPUsDQRCGn70kih8kNiF4xgieIgpWCloE0omdYCEiVloKYmPhf/AviKBga2Fhd6QQtPMXBMRCUFMYNaRKziLn3c1usvpus/POs8Msu6OWMDSGzxObup02SXIsMm/ajhYXOeGGNAP47JNNppRoYJsLMom4xi4PZlXFEZcCBI97ymbVSZ7ppRrLfMiqp/SWx5rewBb9tC7RAv01J9FBCzoi0cCCdiTatqBtiWYtaE6iCxZ0Nok6bGDTAfy+VoWqFW0xw4sDuFxj1xDnZFJuHp8Sf8lj3KHwDxCgknLr5FmNjAZNWtEi+pQd9rrXKqFCY4Lv6FjAMK9hpkm9O1vxTz1mhzcUEDDFIWdxB0qb2DK3jIb7T4p8xSl9DO9Y4ZEG71wxnQThB74BOFIQk88GAAAAAElFTkSuQmCC" />
  </svg>


  <div class="outerApp">
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
  Menu
} from 'floating-vue'
import MainApp from './components/main.vue';
export default {
  components: {
    MainApp,
},
  data() {
    return {
      showSettings: false,
      downloading: false,
      count: 0
    };
  },
  methods: {
    toggleSettings() {
      this.showSettings = !this.showSettings;
    }
  },
  mounted() {
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

<style scoped>
.outerApp {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  position: relative;
  color: dimgray;
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
}</style>
