<template>
    <div class="box">
    <h1 class="boxTitle">Notifications</h1>
    <hr color="#333333">
    <notificationItem class="noti"></notificationItem>
    <hr color="#333333">
    <p class="clearbtn" @click="clicked">Clear</p>
</div>
</template>
<script>
import notificationItem from './notificationItem.vue';
export default {
    components: {
        notificationItem,
    },
    methods: {
        clicked() {
            clearButton = document.querySelector(".clearbtn").textContent = "SEX"
            console.log("ASWA ADJKBAFJKABJKFBAJK")
        }
    }
}
</script>
<style scoped>
.noti {
    overflow: hidden;
}
.box {
    position: relative;
    left: 0px;
    top: 0px;
    width: 400px;
    height: fit-content;
    border-radius: 15px;
    z-index: 9999;
    background-color: #0d1117;
}

.clearbtn {
    position: relative;
    left: 50%;
    translate: -50%;
    width: 100%;
    height: 15px;
    z-index: 9999;
    background-color: #0d1117;
    border: none;
    color: white;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 0px !important;
}

.clearbtn:hover {
    background-color: #111418;
}

.boxTitle {
    position: relative;
    left: 47%;
    translate: -50%;
    width: 106px;
    text-align: center;
    height: 14px;
    z-index: 999;
}

hr {
    margin: none;
}
</style>