<template>
  <div class="container">
    <h1>{{ message }}</h1>
    <a @click="getMessage">Press Me!</a>
  </div>
</template>

<script>
import * as Wails from "@wailsapp/runtime";
export default {
  data() {
    return {
      message: " ",
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.twokey("say hello").then((result) => {
        self.message = result;
      });
      window.backend.Robot.TryEmit().then(console.log);
    },
  },
  mounted() {
    Wails.Events.On("try", (message) => {
      console.log(message);
    });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
