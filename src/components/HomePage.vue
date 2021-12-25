<script setup>
import { useAuth0, AuthState } from "../auth/index";
const { login, logout, initAuth } = useAuth0(AuthState);

initAuth();
</script>

<template>
  <div class="hello">
    <h2> Welcome to Chess Online </h2>
    <p> To get started go to Chess</p>
    <div v-if="!AuthState.loading">
      <div v-if="!AuthState.isAuthenticated">
        <button class="lButton" @click="login()">Login</button>
      </div>
      <div v-else>
        <p> Welcome to VueAuth <strong>{{ AuthState.user.name }}</strong></p>
        <button class="lButton" @click="logout()">Logout</button>
      </div>
    </div>

    <div v-else>
      Loading ...
    </div>
  </div>
</template>

<script>
export default {
    data() {
    return {
      AuthState: AuthState,
    }
  },
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
