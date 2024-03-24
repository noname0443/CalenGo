<script setup>
import { Modal } from 'usemodal-vue3';
import ApiClient from './api/ApiClient';
import DefaultApi from './api/api/DefaultApi';
import VueCookies from 'vue-cookies'
</script>

<template>
  <Modal v-model:visible="isVisible" title="Status Window" :okButton="{text: 'ok', onclick: () => {authResult = null; isVisible = false;}}" :cancelButton="{text: 'cancel', onclick: () => {authResult = null; isVisible = false;}}">
    <div v-if="authResult == null" class="d-flex justify-content-center mb-3">
      <b-spinner></b-spinner>
    </div>
    <p v-else>{{ authResult }}</p>
  </Modal>
  <body>
    <div class="inner-form rounded">
      <b-button-group style="width: 100%;">
        <b-button class="switch-button custom-color" @click="operationType = 'register'">Register</b-button>
        <b-button class="switch-button custom-color" @click="operationType = 'login'">Login</b-button>
      </b-button-group>
      <template v-if="operationType == 'register'">
      <h1 style="text-align: center; padding: 20px;">Register</h1>
      <b-form style="padding: 20px;" @submit="onRegister" v-if="show">
        <b-form-group
          id="input-group-1"
          label="Your Username:"
          label-for="input-1"
          class="widther"
          required
        >
          <b-form-input
            id="input-1"
            v-model="form.username"
            placeholder="Enter username"
            required
          ></b-form-input>
        </b-form-group>
  
        <b-form-group id="input-group-2" label="Your Password:" label-for="input-2">
          <b-form-input
            id="input-2"
            v-model="form.password"
            placeholder="Enter password"
            class="widther"
            required
          ></b-form-input>
        </b-form-group>

        <b-button class="margined custom-color" type="submit">Submit</b-button>
      </b-form>
      </template>
      <template v-else-if="operationType == 'login'">
      <h1 style="text-align: center; padding: 20px;">Login</h1>
      <b-form style="padding: 20px;" @submit="onLogin" v-if="show">
        <b-form-group
          id="input-group-1"
          label="Your Username:"
          label-for="input-1"
          class="widther"
          required
        >
          <b-form-input
            id="input-1"
            v-model="form.username"
            placeholder="Enter username"
            required
          ></b-form-input>
        </b-form-group>
  
        <b-form-group id="input-group-2" label="Your Password:" label-for="input-2">
          <b-form-input
            id="input-2"
            v-model="form.password"
            placeholder="Enter password"
            class="widther"
            required
          ></b-form-input>
        </b-form-group>

        <b-button class="margined custom-color" type="submit">Submit</b-button>
      </b-form>
      </template>
    </div>
  </body>
</template>

<script>
  export default {
    data() {
      return {
        BACKEND_URL: 'http://localhost:1516',
        operationType: 'register',
        isVisible: false,
        authResult: null,
        form: {
          username: '',
          password: '',
        },
        show: true
      }
    },
    methods: {
      onRegister(event) {
        event.preventDefault();
        this.isVisible = !this.isVisible;

        let THIS_OBJ = this;
        var defaultClient = new ApiClient(this.BACKEND_URL);
        var api = new DefaultApi(defaultClient);
        let username = this.form.username;
        let password = this.form.password;
        
        let req = api.postApiV1User({'user':
          {
            'name': username,
            'password': password,
          }
        }, function (error, data, response) {
          console.log(error, data, response);
          if(error != null) {
            THIS_OBJ.authResult = error;
          } else {
            if (data != null) {
              THIS_OBJ.authResult = data;
            } else {
              THIS_OBJ.authResult = 'success';
            }
            VueCookies.set('credentials', username + ":" + password, 3600 * 24 * 7);
          }
        });
      },
      onLogin(event) {
        event.preventDefault();
        this.isVisible = !this.isVisible;

        let THIS_OBJ = this;

        let username = this.form.username;
        let password = this.form.password;
        let credentials = username + ":" + password

        let apiClient = new ApiClient(this.BACKEND_URL);
        apiClient.authentications['BasicAuth'].username = username;
        apiClient.authentications['BasicAuth'].password = password;

        let defaultApi = new DefaultApi(apiClient);

        let req = defaultApi.getApiV1User(username, function (error, data, response) {
          console.log(error, data, response)
          if(error != null) {
            THIS_OBJ.authResult = error;
          } else {
            THIS_OBJ.authResult = 'success';
            VueCookies.set('credentials', username + ":" + password, 3600 * 24 * 7);
          }
        });
      },
    }
  }
</script>

<style scoped>
.custom-color {
  background: #1abc9c;
  color: white;
  border-color: #18ac8e;
}
.custom-color:hover {
  background: #16a084;
}

.switch-button {
  padding: 20px;
  width: 50%;
}

.widther {
  min-width: 300px;
}
.margined {
  margin-top: 5px;
}

body {
  height: 100%;
  display: flex;
  margin: auto;
  background: transparent;
}

.inner-form {
  background: white;
  margin: auto;
}
</style>