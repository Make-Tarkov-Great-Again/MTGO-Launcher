<template>
  <div :class="$style.rectangleParent">
    <div :class="$style.groupChild" />
    <div :class="$style.houseParent">
      <img :class="$style.houseIcon" alt="" src="@/assets/images/SVG/shape.svg" />
      <div :class="$style.mainHub">Main hub</div>
    </div>
    <div :class="$style.trayDownParent">
      <img :class="$style.houseIcon" alt="" src="@/assets/images/SVG/shape1.svg" />
      <div :class="$style.mainHub">Workshop</div>
    </div>
    <div :class="$style.listParent">
      <img :class="$style.houseIcon" alt="" style="fill: white;" src="@/assets/images/SVG/server.svg" />
      <div :class="$style.mainHub">Servers</div>
    </div>
    <div :class="$style.mtga">MTGA</div>
    <div :class="$style.bellParent">
      <img :class="$style.houseIcon" alt="" src="@/assets/images/SVG/36gear.svg" />
      <div :class="$style.mainHub">Notify</div>
    </div>
    <div :class="$style.gearParent">
      <img :class="$style.houseIcon" alt="" src="@/assets/images/SVG/36gear1.svg" />
      <div :class="$style.mainHub">Settings</div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "Group",
});
</script>
<style module>
.mtga {
  content: "MTGA";
  font-weight: bold;
  background-image: url("@/assets/images/colors.png");
  -webkit-background-clip: text;
  color: transparent;
  transition: color 0.3s ease-in-out;
  background-size: 200% 200%;
  cursor: pointer;
}

.mtga::after{
  content: ">";
  position: absolute;
  top: -3px;
  font: arial;
  right: -25px;
  color: silver;
}

.mtga:hover {
  color: rgba(0, 0, 0, 0);
  /* Make text transparent on hover */
  background-position: 100% 100%;
  animation: gradientAnimation 2s linear;
  /* Animated gradient on hover */
}


.trayDownParent,
.houseParent,
.bellParent,
.listParent,
.gearParent {
  transition: background 300ms ease-in;
  user-select: none
}

.trayDownParent:hover,
.houseParent:hover,
.bellParent:hover,
.listParent:hover,
.gearParent:hover {
  background-color: var(--select);
  border: 1px solid var(--stroce);
  cursor: pointer;
}

.groupChild {
  position: absolute;
  top: 0px;
  left: 0px;
  border: 1px solid var(--stroce);
  box-sizing: border-box;
  width: 160px;
  height: 662px;
}

.houseIcon {
  position: relative;
  width: 24px;
  height: 24px;
  overflow: hidden;
  flex-shrink: 0;
}

.mainHub {
  position: relative;
  line-height: 140%;
}

.houseParent {
  position: absolute;
  top: 73px;
  left: 11px;
  border-radius: var(--br-5xs);
  width: 138px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  padding: var(--padding-5xs) var(--padding-base);
  box-sizing: border-box;
  gap: var(--gap-7xs);
}

.trayDownParent {
  position: absolute;
  top: 125px;
  left: 11px;
  border-radius: var(--br-5xs);
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  padding: var(--padding-5xs) var(--padding-base);
  gap: var(--gap-7xs);
}


.listParent {
  position: absolute;
  top: 177px;
  left: 11px;
  border-radius: var(--br-5xs);
  box-sizing: border-box;
  width: 138px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  padding: var(--padding-5xs) var(--padding-base);
  gap: var(--gap-7xs);
}

.mtga {
  position: absolute;
  top: 16px;
  left: 40px;
  font-size: var(--h1-size);
  line-height: 140%;
  font-family: var(--font-bender);
  color: var(--green);
}

.bellParent {
  position: absolute;
  top: 229px;
  left: 11px;
  border-radius: var(--br-5xs);
  width: 138px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  padding: var(--padding-5xs) var(--padding-base);
  box-sizing: border-box;
  gap: var(--gap-7xs);
}

.gearParent {
  position: absolute;
  top: 281px;
  left: 11px;
  border-radius: var(--br-5xs);
  width: 138px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  padding: var(--padding-5xs) var(--padding-base);
  box-sizing: border-box;
  gap: var(--gap-7xs);
}

.rectangleParent {
  position: relative;
  width: 160px;
  height: 662px;
  text-align: left;
  font-size: var(--text-size);
  background-image: url("@/assets/images/R3XfcOd.png");
  color: var(--color-white);
  background-color: var(--select);
  font-family: var(--text);
}


@keyframes gradientAnimation {
  0% {
    background-position: 0% 50%;
  }
  33.333% {
    background-position: 33.333% 50%;
  }
  66.666% {
    background-position: 66.666% 50%;
  }
  100% {
    background-position: 100% 50%;
  }
}
</style>



