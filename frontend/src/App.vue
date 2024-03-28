<script setup>
import { ref } from 'vue';
import ApiClient from './api/ApiClient';
import DefaultApi from './api/api/DefaultApi';
import VueCookies from 'vue-cookies'

VueCookies.config('30d')
VueCookies.config('7d','','',true)

let isUser = ref(false);
let BACKEND_URL = 'http://localhost:1516'

function CheckUser() {
  let credentials = VueCookies.get('credentials');
  if (credentials === null) {
    isUser.value = false;
    return;
  }
  let strings = credentials.split(":");

  let apiClient = new ApiClient(BACKEND_URL);
  apiClient.authentications['BasicAuth'].username = strings[0];
  apiClient.authentications['BasicAuth'].password = strings[1];

  let defaultApi = new DefaultApi(apiClient);

  let req = defaultApi.getApiV1User(strings[0], function (error, data, response) {
    console.log(error, data, response)
    if(error != null) {
      isUser.value = false;
    } else {
      isUser.value = true;
    }
  });
}

CheckUser()
// VueCookies.set('test', 'Hello world!', 3600 * 24 * 7);
</script>

<template>
<body class="main-body">
<div id="block_container">
    <b-navbar id="nav-block" type="dark" variant="dark">
      <b-navbar-nav>
        <RouterLink class="m-2 btn" to="/">Home</RouterLink>
        <RouterLink v-if="isUser" class="m-2 btn" to="/calendar">Calendar</RouterLink>
        <RouterLink v-if="isUser" class="m-2 btn" to="/conflict">Conflict</RouterLink>
        <RouterLink v-if="isUser" class="m-2 btn" to="/user">User</RouterLink>
        <RouterLink v-else class="m-2 btn" to="/register">Register</RouterLink>
      </b-navbar-nav>
    </b-navbar>
</div>
<router-view v-slot="{ Component, route }">
    <div id="wrapper">
      <Transition name="fade" mode="out-in">
        <div style="padding: 20px;" :key="route.name">
          <component :is="Component"></component>
        </div>
      </Transition>
    </div>
  </router-view>
</body>
</template>

<style scoped>
a { text-decoration: none; color: white; }

.fade-enter-active, .fade-leave-active {
  transition: opacity .2s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}

body {
  background-color: rgb(48, 48, 48);
}

#wrapper{
  width: 90%;
  height: 100%;
  margin: auto;
  border-radius: 10px;
}

.disabled {
  opacity: 0.5;
  pointer-events: none;
}

#block_container
{
  display: flex;
  justify-content: center;
}

#nav-block
{ 
  width: 100%;
}
</style>