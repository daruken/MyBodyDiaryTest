<template>
  <div>
    <h1>{{ msg }}</h1>
  </div>
  <div class="login">
    <form @submit.prevent="login">
      <div>
      <label for="id">
        ID
      </label>
      <input v-model="id" type="id" name="id" value />
      </div>

      <div>
      <label for="password">
        Password
      </label>
      <input v-model="password" type="password" name="password" value />
      </div>

      <button type="submit" name="button">
        Login
      </button>

      <router-link to="/signup">
        회원 가입
      </router-link>
    </form>
  </div>
</template>

<script>
export default {
  name: 'Login',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.login {
  width 450px
  height 250px
  pading 5px 20px
  top 50%
  left 50%
  margin-left -220px
  margin-top -80px
  position relative
  display flex
  flex-direction columns
  justify-content center
  align-items center
}

.form > div {
  display flex
  justify-content center
  pading-bottom 7px
  align-items center
}

label {
  flex 1
  text-align left
  float left
  margin-top 12px
}
input {
  margin 5px
  float right
  padding 5px
}
button {
  width 85px
  float right
  padding 3px
}
a {
  float left
}
</style>
