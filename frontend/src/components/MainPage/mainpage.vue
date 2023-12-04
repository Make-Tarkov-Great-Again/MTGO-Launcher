<template >
    <div class="mainpage">
        <h1 class="greeting">Good Morning, What are we doing today?</h1>
        <ActionButton @option-selected="handleOptionSelected"></ActionButton>
    </div>
    <FirtScreen v-if="pickedWorkshop"/>

</template>
<script>
import AkiWorkshopComponent from '../Workshop/AkiWorkshopComponent.vue';
import ActionButton from '../buttons/action-button.vue';
import FirtScreen from '../Startup/FirtScreen.vue'; //sex

export default {
components: {
    ActionButton,
    AkiWorkshopComponent,
    FirtScreen,
},
data() {
    return {
        pickedWorkshop: false
    };
  },
methods: {
    handleOptionSelected(selectedOption) {
    },
}
}
</script>
<style scoped>



.mainpage {
    display: inline-block;
    width: 100%;
    height: 100%;
    max-width: 100%;
    max-height: 100%;
    min-width: 100%;
    min-height: 100%;
}
.greeting {
    position: relative;
    color: white;
    left: 50%;
    top: 10px;
    translate: -50% -50%;
    max-width: fit-content;
    max-height: fit-content;
}
</style>