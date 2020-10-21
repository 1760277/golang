<template>
  <div class="form-upsert">
    <div class="update" v-if="user">
      <md-dialog-title>Update User</md-dialog-title>
      <div>
          <span v-if="messageError">{{ messageError }}</span>
      </div>
      <div class="field-input">
        <label>Username: </label>
        <input name="username" type="text" v-model="user.user_name" />
      </div>
      <div class="field-input">
        <label>Status: </label>
        <md-switch v-model="user.status"></md-switch>
      </div>
      <div class="panel-btn">
        <button name="update" value="Update" v-on:click="updateUser(user)">Update</button>
        <button name="close" value="Close" v-on:click="closeDialog()">
          Close
        </button>
      </div>
    </div>
    <div class="insert" v-if="!user">
      <md-dialog-title>Insert User</md-dialog-title>
      <div class="field-input">
        <label>Username: </label>
        <input name="username" type="text" v-model="userInsert.user_name" />
      </div>
      <div class="field-input">
        <label>Password: </label>
        <input name="password" type="password" v-model="userInsert.password" />
      </div>
      <div class="field-input">
        <label>Status: </label>
        <md-switch v-model="userInsert.status"></md-switch>
      </div>
      <div class="panel-btn">
        <button name="insert" value="Insert" v-on:click="insertUser(userInsert)">Insert</button>
        <button name="close" value="Close" v-on:click="closeDialog()">
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import UserService from "../../services/userService";
import User from "../User.vue";

@Component({
  props: ["user", "showDialog", "users"],
})
export default class UpsertUser extends Vue {
  private userInsert: any = {
    user_name: "",
    password: "",
    status: false,
  };
  private messageError: string = '';

  private closeDialog() {
    this.$emit("changeShowDialog", false);
  }

  private updateUser(user: any) {
      UserService.UpdateUser(user)
        .then((res) => {
            if(res.status === 200) {
                this.$emit("changeShowDialog", false);
                this.$emit("insertUser", res.data);
                
            }else{
                this.messageError = res.data.error;
            }
        })
        .catch((error) => {
            console.log(error);
        })
  }

  private insertUser(user: any) {
      UserService.AddUser(user)
        .then((res) => {
            if(res.status === 200) {
                this.$emit("changeShowDialog", false);
                this.$emit("insertUser", res.data);
                
            }else{
                this.messageError = res.data.error;
            }
        })
        .catch((error) => {
            console.log(error);
        })
  }
}
</script>

<style scoped>
.form-upsert {
  margin: 2rem 3rem;
  background-color: #fff;
}
.insert,.update {
    padding: 1rem 1.5rem;
}
.field-input {
    margin-top: 0.5rem;
}

.panel-btn {
    margin: auto;
    text-align: center;
}

.panel-btn button {
    margin: 0 1rem;
}
</style>