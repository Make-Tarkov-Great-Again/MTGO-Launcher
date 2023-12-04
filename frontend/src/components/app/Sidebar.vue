<template>
  <div :class="$style.frameParent">
    <div :class="$style.rectangleParent">
      <div :class="$style.frameChild" />
      <div :class="$style.frameItem" />
      <div :class="$style.house">
        <img :class="$style.shapeIcon" alt="" src="@/assets/images/SVG/shape.svg" />
      </div>
      <div :class="$style.gear">
        <img :class="$style.shapeIcon1" alt="" src="@/assets/images/SVG/shape1.svg" />
      </div>
      <img :class="$style.gearIcon" alt="" src="@/assets/images/SVG/36gear.svg" />
      <img :class="$style.gearIcon" alt="" src="@/assets/images/SVG/36gear1.svg" />
    </div>
    <div :class="$style.frameInner" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "Frame",
  });
</script>
<style module>
  .frameChild {
    position: relative;
    border-radius: var(--br-5xs);
    background-color: #294a88;
    width: 48px;
    height: 48px;
  }
  .frameItem {
    position: relative;
    background-color: var(--color-gray-100);
    width: 48px;
    height: 1px;
  }
  .shapeIcon {
    position: relative;
    width: 26px;
    height: 25px;
  }
  .house {
    border-radius: var(--br-5xs);
    background-color: var(--color-gray-200);
    border: 1px solid rgba(189, 189, 192, 0.08);
    overflow: hidden;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    padding: var(--padding-base);
  }
  .shapeIcon1 {
    position: relative;
    width: 26px;
    height: 26.14px;
  }
  .gear {
    border-radius: var(--br-5xs);
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-base);
  }
  .gearIcon {
    position: relative;
    width: 36px;
    height: 36px;
    overflow: hidden;
    flex-shrink: 0;
  }
  .rectangleParent {
    position: absolute;
    top: calc(50% - 331px);
    left: calc(50% - 510px);
    background-color: var(--color-gray-100);
    height: 662px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: var(--padding-base) 32px;
    box-sizing: border-box;
    gap: 24px;
  }
  .frameInner {
    position: absolute;
    top: 454px;
    left: 353px;
    border-radius: 8888px;
    background-color: #375a9b;
    filter: blur(888px);
    width: 415px;
    height: 415px;
  }
  .frameParent {
    background-color: var(--color-gray-200);
    width: 60px;
    height: 662px;
    overflow: hidden;
  }
</style>
