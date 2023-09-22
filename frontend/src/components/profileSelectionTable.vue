<template>
  <table class="profile-table">
    <thead>
      <tr>
        <th>Side</th>
        <th>Username</th>
        <th>Level</th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(profile, index) in profiles" :key="profile.id" class="profile-row">
        <td class="side">{{ profile.Side }}</td>
        <td class="username">{{ profile.username }}</td>
        <td class="level">{{ profile.Level }}</td>
        <td class="action">
          <button @click="selectProfile(profile)">Select</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  props: {
    profiles: {
      type: Array,
      required: true,
    },
  },
  methods: {
    getSideImage(side) {
      if (side === 'Bear') {
        return '/src/assets/images/Tarkov/BEAR_icon.png';
      } else if (side === 'Usec') {
        return '/src/assets/images/Tarkov/USEC_icon.png';
      } else {
        return 'https://upload.wikimedia.org/wikipedia/en/d/d0/Dogecoin_Logo.png';
      }
    },
    getRankImage(level) {
      const rankNumber = Math.ceil(level / 5) * 5;
      return `/src/assets/images/Tarkov/ranks/Rank${rankNumber}.png`;
    },
    getProfileIcon(icon) {
      return `/src/assets/images/Tarkov/icons/${icon}.png`;
    },
  },
    mounted() {
        console.log(this.profile);
    }
};
</script>

<style scoped>
table.profile-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
  font-family: Arial, sans-serif;
  border-radius: 15px;
}

table.profile-table th {
  background-color: var(--GithubA-dark);
  color: #ffffff;
  text-align: left;
  padding: 8px;
}

table.profile-table th, table.profile-table td {
  border: 1px solid #d1d5da;
  padding: 8px;
}

table.profile-table tr:nth-child(even) {
  background-color: var(--GithubM-dark);
}

table.profile-table tr:hover {
  background-color: #0D1117;
}

.profile-table .action button {
  background-color: var(--GithubA-dark);
  color: #ffffff;
  border: none;
  cursor: pointer;
  transition: background-color 0.3s;
}

.profile-table .action button:hover {
  background-color: #0e131a;
}
</style>
