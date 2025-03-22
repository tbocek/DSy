<template>
  <div class="about">
    <h1>About Page</h1>
    <p v-if="response">{{ response.hallo }}</p>
    <p v-else>Loading...</p>
  </div>
</template>

<script>
export default {
  name: 'About',
  data() {
    return {
      response: null,
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const result = await fetch('/api');
        if (!result.ok) {
          throw new Error(`HTTP error! status: ${result.status}`);
        }
        const data = await result.json();
        this.response = data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
  },
};
</script>