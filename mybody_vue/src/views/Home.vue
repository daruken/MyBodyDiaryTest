<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <Login msg="Welcome to MyBodyDiary!"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Login from '@/views/Login.vue'

export default {
  name: 'Home',
  components: {
    Login
  }
}
</script>
