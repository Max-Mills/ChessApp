<template>
  <div class="home">
    <img class="logo" alt="Chess Online Logo" src="https://chessapp.s3.us-west-2.amazonaws.com/wpawn.png">
    <HomePage />
  </div>
</template>

<script>
// @ is an alias to /src
import HomePage from '../components/HomePage.vue'

export default {
  name: 'Home',
  components: {
    HomePage
  }
}
</script>
