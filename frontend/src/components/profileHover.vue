<template>
    <div class="container">
        <p class="craft1">Nothings Crafting!</p>
        <h3 class="craftingTime">Crafting Time</h3>
        <h3 class="scavTimer">Scav Timer</h3>
        <p class="scavtimeText" :style="{ color: scavTimeColor }">{{ formattedScavTime }}</p>
        <p class="euroAmount">{{ formatNumber(profile.eurAmount) }}</p>
        <p class="usdAmount">{{ formatNumber(profile.usdAmount) }}</p>
        <p class="roublesAmount">{{ formatNumber(profile.rubAmount) }}</p>
        <h1 class="roublesSymbol">₽</h1>
        <h1 class="euroSymbol">€</h1>
        <h1 class="usdSymbol">$</h1>
        <hr class="1">
    </div>
</template>

<script>
export default {
    props: {
        profile: {
            type: Object,
            required: true,
        },
    },
    data() {
        return {
            scavTimerInterval: null,
            scavTimeColor: 'white',
            remainingTime: 0,
            playSound: false,
        };
    },
    computed: {
        scavExpirationTime() {
            return this.profile.SavageLockTime * 1000;
        },
        formattedScavTime() {
            if (this.remainingTime <= 0) {
                this.scavTimeColor = 'green';
                return 'Done';
            }

            const hours = Math.floor(this.remainingTime / 3600000);
            const minutes = Math.floor((this.remainingTime % 3600000) / 60000);
            const seconds = Math.floor((this.remainingTime % 60000) / 1000);

            return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
        },
    },
    methods: {
        startScavTimer() {
            const now = Date.now();
            this.remainingTime = this.scavExpirationTime - now;

            this.scavTimerInterval = setInterval(() => {
                if (this.remainingTime > 0) {
                    this.remainingTime -= 1000;
                    this.playSound = true
                } else {
                    clearInterval(this.scavTimerInterval);
                }
            }, 1000);
        },
        formatNumber(number) {
            const abbreviations = ["", "k", "M", "B", "T"];
            let index = 0;

            while (number >= 1000 && index < abbreviations.length - 1) {
                number /= 1000;
                index++;
            }

            return number.toFixed(1) + abbreviations[index];
        },
    },
    mounted() {
        this.startScavTimer();

        const scavTimers = JSON.parse(localStorage.getItem('ScavTimers')) || {};

        scavTimers[this.profile.username] = this.profile.SavageLockTime;

        localStorage.setItem('ScavTimers', JSON.stringify(scavTimers));
    },
    beforeDestroy() {
        clearInterval(this.scavTimerInterval);
    },
};
</script>



<style scoped>
.container {
    background-color: #010205;
    color: white;
    height: 200px;
    width: 220px;
    border: 1px solid white;
}

.scavTimer {
    position: absolute;
    top: 55px;
    left: 50%;
    translate: -50%;
}

.scavtimeText {
    position: absolute;
    top: 75px;
    left: 50%;
    translate: -50%;
}

.craftingTime {
    white-space: nowrap;
    position: absolute;
    top: 100px;
    left: 50%;
    translate: -50%;
}

.craft1 {
    white-space: nowrap;
    position: absolute;
    top: 135px;
    left: 50%;
    translate: -50%;
}

.craft2 {
    white-space: nowrap;
    position: absolute;
    top: 154px;
    left: 50%;
    translate: -50%;
}

body {
    color: white;
}

.roublesSymbol {
    color: white;
    position: absolute;
    top: 1px;
    left: 25%;
    translate: -25%;
}

.usdSymbol {
    color: white;
    position: absolute;
    top: 1px;
    left: 50%;
    translate: -50%;
}

.euroSymbol {
    color: white;
    position: absolute;
    top: 1px;
    left: 75%;
    translate: -75%;
}

.roublesAmount {
    color: white;
    position: absolute;
    top: 35px;
    left: 23%;
    translate: -23%;
}

.usdAmount {
    color: white;
    position: absolute;
    top: 35px;
    left: 50%;
    translate: -50%;
}

.euroAmount {
    color: white;
    position: absolute;
    top: 35px;
    left: 76%;
    translate: -75%;
}
</style>