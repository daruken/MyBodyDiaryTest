<template>
  <div class="container">
    <form @submit.prevent="submitForm">
      <div class="wrap_cont">
        <label for="email">E-MAIL : </label>
        <input type="text" id="email" v-model="email" />
      </div>
      <div class="wrap_cont">
        <label for="id">ID : </label>
        <input type="text" id="id" v-model="id" />
      </div>
      <div class="wrap_cont">
        <label for="password">비밀번호 : </label>
        <input type="password" id="password" v-model="password" />
      </div>
      <div class="wrap_cont">
        <label for="passwordConfirm">비밀번호 확인 : </label>
        <input type="password" id="passwordConfirm" v-model="passwordConfirm" />
      </div>
      <button type="submit">회원가입</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      email: '',
      id: '',
      password: ''
    }
  },
  methods: {
    submitForm () {
      if (this.password !== this.passwordConfirm) {
        alert('비밀번호가 일치하지 않습니다.')
        return false
      }

      axios.post('http://localhost:8080/signup', {
        email: this.email,
        id: this.id,
        password: this.password
      }).then(res => {
        this.$router.push({ name: '/login' })
      })
    }
  }
}
</script>

<style scoped lang="stylus">
.container {
  position relative
}

label {
  display inline-block
  text-align right
  width 140px
}
input {
  text-align right
  padding .5em .5em
  margin 5px
  width 200px
}
button {
  width 100px
  margin-top 15px
  padding 10px
}
</style>
