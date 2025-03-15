/*
1) Install Vue CLI using the command: npm install @vue/cli.
2) Create a new Vue project using the command npx vue create ost-app. Choose the default preset when prompted.
3) Navigate to the newly created ost-app directory using the command cd ost-app.
4) Open the src/App.vue file and replace its content with the file ost-app/src/App.vue
5) Start the Go API server by running go run api.go in the terminal, where the api.go file is located.
6) Start the Vue.js development server by running npm run serve in the terminal, where the ost-app directory is located.
7) Open a web browser and navigate to http://localhost:8081. You should see the "OST" text fetched from the Go API and displayed in the Vue.js SPA.
*/

<template>
  <div id="app">
    <h1>{{ text }}</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: "",
    };
  },
  async mounted() {
    try {
      const response = await fetch("http://localhost:8080/api");
      const data = await response.json();
      this.text = data.text;
    } catch (err) {
      console.error("Error fetching data:", err);
    }
  },
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>