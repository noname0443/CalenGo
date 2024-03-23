<script setup>
import { Modal } from 'usemodal-vue3';
import ApiClient from './api/ApiClient';
import DefaultApi from './api/api/DefaultApi';
import VueCookies from 'vue-cookies'
</script>

<template>
  <Modal v-model:visible="isVisible" title="Status Window" :okButton="{text: 'ok', onclick: () => {registerResult = null; isVisible = false;}}" :cancelButton="{text: 'cancel', onclick: () => {registerResult = null; isVisible = false;}}">
    <div v-if="registerResult == null" class="d-flex justify-content-center mb-3">
      <b-spinner></b-spinner>
    </div>
    <p v-else>{{ registerResult }}</p>
  </Modal>
  <body>
    <div class="inner-form rounded">
      <b-form @submit="onSubmit" v-if="show">
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

        <b-button class="margined" type="submit" variant="primary">Submit</b-button>
      </b-form>
    </div>
  </body>
</template>

<script>
  export default {
    data() {
      return {
        BACKEND_URL: 'http://localhost:1516',
        isVisible: false,
        registerResult: null,
        form: {
          username: '',
          password: '',
        },
        show: true
      }
    },
    methods: {
      onSubmit(event) {
        event.preventDefault()
        this.isVisible = !this.isVisible;

        let THIS_OBJ = this;
        let apiClient = new ApiClient(this.BACKEND_URL);
        let defaultApi = new DefaultApi(apiClient);
        let req = defaultApi.postApiV1User({'user':
          {
            'username': this.form.username,
            'password': this.form.password,
          }
        }, function (error, data, response) {
          console.log(error, data, response)
          if(error != null) {
            THIS_OBJ.registerResult = error;
          } else {
            THIS_OBJ.registerResult = data;
          }
        });
      },
    }
  }
</script>

<style scoped>
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
  padding: 20px;
}
</style>