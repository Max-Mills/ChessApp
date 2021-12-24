<template>
  <div class="hello">
    <h2> Welcome to Chess Online </h2>
    <button @click="loginClick()">Log in</button>
    <p> To get started go to Chess</p>
  </div>
</template>

<script>
import {login} from '../auth/index'
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  methods: {
    // Log the user in
    loginClick() {
      login();
    },
    // Log the user out
    logout() {
      this.$auth.logout({
        returnTo: window.location.origin
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
