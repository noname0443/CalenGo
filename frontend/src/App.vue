<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import { ref } from 'vue';
const count = ref(0)

import ApiClient from './api/ApiClient';
import Note from './api/model/Note';
import User from './api/model/User';
import DefaultApi from './api/api/DefaultApi';

var apiClient = new ApiClient('http://localhost:1516');
var xxxSvc = new DefaultApi(apiClient);
var zzz = xxxSvc.getApiV1Note('1', function (error, data, response) {
    console.log(error, data, response)
});
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
    </div>
  </header>

  <main>
    <TheWelcome />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
