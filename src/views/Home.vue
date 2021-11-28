<template>
  <div class="home">
    <img class="logo" alt="Chess Online Logo" src="http://localhost:8082/assets/wpawn.png">
    <HomePage />
  </div>
</template>

<script>
// @ is an alias to /src
import HomePage from '@/components/HomePage.vue'

export default {
  name: 'Home',
  components: {
    HomePage
  }
}
</script>
