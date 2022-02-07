<script setup>
import { useAuth0, AuthState } from "../utils/useAuth0";
const { login, logout, initAuth } = useAuth0(AuthState);

initAuth();
</script>

<template>
<div style="position: absolute; top: 20px; right: 20px;">
  <div v-if="!AuthState.loading">
    <div v-if="!AuthState.isAuthenticated">
      <button @click="login()" class="btn btn-primary">Login</button>
    </div>
    <div v-else>
      <div><img style="border-radius: 80px; border: 5px solid black; width: 80px; height: 80px" :src="AuthState.user.picture"></div>
      <div><strong>{{ AuthState.user.nickname }}</strong></div>
      <div><button @click="logout()" class="btn btn-secondary">Logout</button></div>
    </div>
  </div>

  <div v-else>
    Loading ...
  </div>
</div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.btn {
  padding: 8px 12px;
  margin-bottom: 0;
  font-size: 14px;
  font-weight: 400;
  line-height: 1.5;
  border: none;
  cursor: pointer;
  min-width: 100px;
  border-radius: 4px;
  font-weight: bold;
}

.btn-primary {
  background: #41B883;
  color: white;
}

.btn-secondary {
  background: #aaa;
  color: white;
}
</style>
