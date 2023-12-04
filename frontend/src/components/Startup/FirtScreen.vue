<template>
  <div :class="$style.rectangleParent">
    <div :class="$style.frameChild" />
    <div :class="$style.frameItem" />
    <div :class="$style.frameInner" />
    <div :class="$style.rectangleDiv" />
    <b :class="$style.mtga">MTGA</b>
    <div :class="$style.itJustText">It just text lalalal yes</div>
    <div :class="$style.buttonWrapper">
      <b :class="$style.button">button</b>
    </div>
    <b :class="$style.pickYourPoison">Pick your poison </b>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "Frame3",
  });
</script>
<style module>
  .frameChild {
    position: absolute;
    top: 454px;
    left: calc(50% - 207px);
    border-radius: var(--br-8869xl);
    background-color: var(--color-darkslateblue-100);
    filter: blur(888px);
    width: 415px;
    height: 415px;
  }
  .frameItem {
    position: absolute;
    top: calc(50% - 298px);
    left: calc(50% - 459px);
    border: 2px solid var(--color-crimson);
    box-sizing: border-box;
    width: 918px;
    height: 596px;
  }
  .frameInner {
    position: absolute;
    top: 208px;
    left: 841px;
    border-radius: var(--br-5xs);
    background: linear-gradient(
        rgba(189, 189, 192, 0.08),
        rgba(189, 189, 192, 0.08)
      ),
      #151618;
    width: 671px;
    height: 312px;
  }
  .rectangleDiv {
    position: absolute;
    top: 205px;
    left: 146px;
    border-radius: var(--br-5xs);
    background: linear-gradient(
        rgba(189, 189, 192, 0.08),
        rgba(189, 189, 192, 0.08)
      ),
      #151618;
    width: 671px;
    height: 312px;
  }
  .mtga {
    position: absolute;
    top: 430px;
    left: 170px;
  }
  .itJustText {
    position: absolute;
    top: 455px;
    left: 170px;
    display: inline-block;
    width: 128px;
  }
  .button {
    position: relative;
  }
  .buttonWrapper {
    position: absolute;
    top: 446px;
    left: 665px;
    border-radius: var(--br-5xs);
    background-color: var(--color-gray-100);
    border: 1px solid var(--color-silver-200);
    box-sizing: border-box;
    width: 128px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-base) var(--padding-13xl);
    font-size: var(--font-size-xs);
  }
  .pickYourPoison {
    position: absolute;
    top: 142px;
    left: 146px;
    font-size: var(--font-size-13xl);
  }
  .rectangleParent {
    position: relative;
    background-color: var(--color-gray-200);
    width: 100%;
    height: 100%;
    overflow: hidden;
    text-align: left;
    font-size: var(--font-size-base);
    color: var(--color-white);
    font-family: var(--font-inter);
  }
</style>
