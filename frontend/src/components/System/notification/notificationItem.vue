<template>
    <div class="notification-item">
        <p class="title">Ligma</p>
        <p class="from">From System</p>
        <p class="description">The system caught ligma.</p>
    </div>
</template>
<script>
export default {

}
</script>
<style scoped>
.notification-item {
  position: relative;
  left: 0px;
  top: 0px;
  width: 277px;
  height: 44px;
  margin: none;
}
.from {
    font-size: 12px;
    color: rgb(105, 105, 105);
    line-height: 1.2;
    text-align: center;
    position: relative;
    margin: inherit;
    left: 14%;
    top: -39%;
    translate: -50% -50%;
}

.description {
    font-size: 12px;
    color: rgb(105, 105, 105);
    height: 85px;
    width: max-content;
    /* line-height: 1.2; */
    text-align: left;
    position: relative;
    margin: inherit;
    top: -40px;
    left: 90px;
}

.title {
    font-size: 18px;
    color: rgb(105, 105, 105);
    line-height: 1.2;
    text-align: center;
    position: relative;
    margin: inherit;
    top: 50%;
    left: 10%;
    translate: -50% -50%;
}


</style>