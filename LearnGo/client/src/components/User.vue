<template>
  <div class="list-user">
    <md-table v-model="users" md-card v-if="users">
      <md-table-toolbar>
        <div class="md-toolbar-section-start">
          <h1 class="md-title">Users</h1>
        </div>
        <md-field md-clearable class="md-toolbar-section-end">
          <div>
            <button v-on:click="openInsertDialog(null)" value="Add">Add</button>
          </div>
        </md-field>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="ID" md-sort-by="id" md-numeric>{{
          item.id
        }}</md-table-cell>
        <md-table-cell md-label="UserName" md-sort-by="user_name">{{
          item.user_name
        }}</md-table-cell>
        <md-table-cell md-label="Status" md-sort-by="status">{{
          item.status
        }}</md-table-cell>
        <md-table-cell md-label="Action">
          <button v-on:click="openEditDialog(item)" value="Edit">Edit</button>
          <button v-on:click="deleteUser(item)" value="Delete">
            Delete
          </button>
        </md-table-cell>
      </md-table-row>
    </md-table>
    <md-dialog :md-active.sync="showDialog">
      <upsert-user
        :user="userSelected"
        :users="users"
        :showDialog="showDialog"
        @changeShowDialog="showDialog = $event"
        @insertUser="changeFromDialog($event)"
      ></upsert-user>
    </md-dialog>
    <!-- <transition name="EditUser">
      <upsert-user
        :user="userSelected"
        :showDialog="showDialog"
        v-if="showDialog"
        @changeShowDialog="showDialog = $event"
      ></upsert-user>
    </transition> -->
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import UserService from "../services/userService";
import UpsertUser from "../components/Dialog/UpsertUser.vue";

@Component({
  components: {
    UpsertUser,
  },
})
export default class User extends Vue {
  private users: any = null;
  private userSelected: any = null;
  private showDialog: boolean = false;

  private openEditDialog(user: any) {
    this.userSelected = user;
    this.showDialog = true;
  }

  private openInsertDialog(user: any) {
    this.userSelected = user;
    this.showDialog = true;
  }

  private deleteUser(user: any) {
      if(confirm("Do you really want to delete?")) {
         UserService.DeleteUser(user.id)
         this.users.splice(this.users.indexOf(user), 1)
      }
  }

  private mounted() {
    UserService.GetUser()
      .then((res) => {
        this.users = res.data;
      })
      .catch((error) => {});
  }

  private changeFromDialog(user: any) {
      this.users.push(user);
  }
}
</script>

<style scoped>
.md-dialog-container {
  background-color: #fff;
}
</style>