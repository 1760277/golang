<template>
  <div class="container">
    <div class="row">
      <form class="form-login" v-on:submit.prevent="onSubmit">
        <div>
          <span class="title-login"><b>Demo</b></span>
        </div>
        <div v-if="messageError">
          <span class="mes-error">{{ messageError }}</span>
        </div>
        <md-field>
          <label>Username</label>
          <md-input v-model="username"></md-input>
        </md-field>
        <md-field>
          <label>Password</label>
          <md-input v-model="password" type="password"></md-input>
        </md-field>

        <md-button type="submit">Login</md-button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import UserService from "../services/userService";

@Component
export default class Login extends Vue {
  private username: string = "";
  private password: string = "";
  private messageError: string = "";
  private onSubmit() {
    const data = {
      user_name: this.username,
      password: this.password,
    };
    UserService.Login(data)
      .then((res) => {
        localStorage.setItem("demotoken", "test");
        this.$router.push({ name: 'user' });
      })
      .catch((error) => {
        this.messageError = error;
      });
  }
}
</script>
<style scoped>
.container {
  max-width: 300px;
  margin: 200px auto;
}

.form-login {
  border: 1px solid;
  padding: 3em 2em 1em 2em;
}

.title-login {
  font-size: 2.5em;
}
</style>