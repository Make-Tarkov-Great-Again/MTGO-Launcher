<template>
    <div class="box">
      <h1 class="boxTitle">Notifications</h1>
      <hr color="#333333">

      <div v-if="Object.keys(notifications).length > 0">
        <div v-for="(notification, key) in notifications" :key="key">
          <notificationItem class="noti" :notification="notification"></notificationItem>
          <hr color="#333333">
        </div>
        <p class="clearbtn" @click="clearNotifications">Clear</p>
      </div>

      <p v-else class="no-notifications" style=" position: relative; width: fit-content; left: 50%; translate: -50%;">No notifications</p>
    </div>
  </template>

<script>
import notificationItem from './notificationItem.vue';
import mitt from 'mitt';
import eventEmitter from "../../js/notifacation.js"


export default {
  components: {
    notificationItem,
  },
  data() {
    return {
      notifications: {},
    };
  },
  methods: {
    pollLocalStorage() {
      const storedNotifications = localStorage.getItem('notifications');
      if (storedNotifications !== JSON.stringify(this.notifications)) {
        this.notifications = JSON.parse(storedNotifications);
        this.emitter.emit('notifications-changed', this.notifications);
      }
    },
    clearNotifications() {
      localStorage.removeItem('notifications');
      this.notifications = [];
      localStorage.setItem("notifications", JSON.stringify(this.notifications));
    },
  },
  created() {
    const storedNotifications = localStorage.getItem('notifications');
    if (storedNotifications) {
      this.notifications = JSON.parse(storedNotifications);
    }
    this.localStoragePolling = setInterval(this.pollLocalStorage, 1000);
  },
  beforeUnmount() {
    clearInterval(this.localStoragePolling);
  },
  setup() {
    const emitter = eventEmitter //fuck you emitter
    return {
      emitter,
    };
  },
};
</script>
<style scoped>
.noti {

}
.box {
  position: relative;
    left: 0px;
    top: 0px;
    width: 400px;
    height: min-content;
    max-height: 600px;
    border-radius: 15px;
    z-index: 9999;
    background-color: #0d1117;
    overflow-y: auto;
    overflow-x: hidden;
}

.clearbtn {
  cursor: pointer;
    position: relative;
    left: 50%;
    translate: -50%;
    width: 100%;
    height: 25px;
    z-index: 9999;
    background-color: #0d1117;
    border: none;
    color: white;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 5px !important;
}

.clearbtn:hover {
    background-color: var(--GithubHov-dark);
}

.boxTitle {
    position: relative;
    left: 50%;
    translate: -50%;
    width: fit-content;
    text-align: center;
    height: 14px;
    z-index: 999;
}

hr {
    margin: none;
}
</style>