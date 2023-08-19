<template>
    <div class="single-mod-page">
      <header>
        <!-- Header content goes here -->
      </header>

      <div class="container">
        <div class="main-container">
          <div class="mod-name">{{ mod.title }}</div>

          <div class="mod-picture-gallery">
            <img v-for="(image, index) in mod.images" :key="index" :src="image" :alt="mod.title">
          </div>

          <div class="mod-name-subscribe">
            <div class="mod-name">{{ mod.title }}</div>
            <div class="grab-button" @click="handleGrabButtonClick(mod.id)">Grab</div>
          </div>

          <div class="mod-description">{{ mod.description }}</div>
        </div>
      </div>
    </div>
  </template>

<script>
export default {
  props: ['id'], // Receive the mod ID as a prop
  data() {
    return {
      mod: {}, // Mod data will be fetched from API in the future
    };
  },
  methods: {
    handleGrabButtonClick(modID) {
      // Handle grab button click
    },
  },
  created() {
    // Fetch mod data using the mod ID
    // For now, use switch statement for mock data
    switch (this.id) {
      case 1:
        this.mod = {
          id: 1,
          title: 'Escape from Hell',
          images: [
            'https://i.imgur.com/oyrIBpI.png',
            'https://i.imgur.com/oyrIBpI.png'
          ],
          author: 'EFHDev',
          tags: ['Server', 'Client', 'Overhaul'],
          version: '3.5.0',
          description: 'This is a description of the mod "Escape from Hell". Lorem ipsum dolor sit amet...',
        };
        break;
      case 2:
        this.mod = {
          id: 2,
          title: 'MP-43 12GA SAWED-OFF DOUBLE-BARREL SHOTGUN',
          images: [
            'https://hub.sp-tarkov.com/files/images/file/6e/1395.png',
            'https://hub.sp-tarkov.com/files/images/file/6e/1395.png'
          ],
          author: 'Mighty_Condor',
          tags: ['Server', 'Weapons'],
          version: '3.5.0',
          description: 'This is a description of the mod "MP-43 12GA SAWED-OFF DOUBLE-BARREL SHOTGUN". Lorem ipsum dolor sit amet...',
        };
        break;
      // Add more cases for other mod IDs here
      default:
        // Handle unknown mod IDs
        break;
    }
  },
};
</script>


  <style scoped>
:root {
  --GithubM-dark: #01040900; /* Main color */
  --GithubA-dark: #0d111780; /* Accent */
  --Github-orange: #F78166;
  --text-color: #EDEDED;
  --aki-orange: #ffc107;
  --mtga-blue: #262f83;
}

.single-mod-page {
  background-color: var(--GithubM-dark);
  color: var(--text-color);
  padding: 20px;
}

.container {
  background-color: var(--GithubA-dark);
  padding: 20px;
  border-radius: 8px;
  margin-top: 20px;
}

.main-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.mod-name {
  font-size: 24px;
  margin-bottom: 10px;
}

.mod-picture-gallery {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.mod-picture-gallery img {
  max-width: 100%;
  height: auto;
  border-radius: 8px;
}

.mod-name-subscribe {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
  margin-bottom: 20px;
}

.grab-button {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: var(--mtga-blue);
  color: white;
  font-size: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}

.mod-description {
  font-size: 16px;
  line-height: 1.5;
}  </style>
