<script setup>
import { useAuth0, AuthState } from "../utils/useAuth0";
const { login, logout, initAuth } = useAuth0(AuthState);

initAuth();
</script>

<template>
<div>
  <div v-if="!AuthState.loading">
    <div v-if="!AuthState.isAuthenticated">
      <p> You must login to start playing </p>
    </div>
    <div v-else>
      <p> Welcome to Chess Online, to start navigate to Chess Games </p>
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
