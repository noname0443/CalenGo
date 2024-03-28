<script setup>
import { Collapse } from 'vue-collapsed'
import { reactive } from 'vue';
import { ConflictApiV1200ResponseInner } from './api';
import ApiClient from './api/ApiClient';
import DefaultApi from './api/api/DefaultApi';
import VueCookies from 'vue-cookies'

let global_notes = reactive([])
let BACKEND_URL = 'http://localhost:1516'

function getConflicts() {
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

  defaultApi.conflictApiV1(function (error, data, response) {
    console.log(error, data, response)
    if(error !== null) {
      return;
    }
    data.forEach(function (element) {
      element.collapse = false;
    });
    global_notes.length = 0;
    global_notes.push(...data);
    console.log(global_notes)
  });
}
getConflicts()
</script>

<template>
<body>
<div class="screen">
    <h1 class="header">Conflict</h1>
</div>
<div class="text">
    <div v-for="(note, index) in global_notes">
      <h2 @click="global_notes[index].collapse = !global_notes[index].collapse" class="form-signin">
        <template v-if="!global_notes[index].collapse">
          + {{ note.first.name }} | {{ note.second.name }}
        </template>
        <template v-else>
          - {{ note.first.name }} | {{ note.second.name }}
        </template>
      </h2>
      <Collapse :when="global_notes[index].collapse">
        <div style="display: flex;">
          <form class="form-signin" style="width: 100%;">
            <h2>{{ note.first.name }}</h2>
            <textarea disabled="disabled" rows="8" class="form-textarea" v-bind:value="note.first.description" placeholder="description"/>
            <input disabled="disabled" type="text" class="form-control" v-bind:value="note.first.startTime" placeholder="start time"/>
            <input disabled="disabled" type="text" class="form-control" v-bind:value="note.first.endTime" placeholder="end time" />
          </form>
          <form class="form-signin" style="width: 100%;">
            <h2>{{ note.second.name }}</h2>
            <textarea disabled="disabled" rows="8" class="form-textarea" v-bind:value="note.second.description" placeholder="description"/>
            <input disabled="disabled" type="text" class="form-control" v-bind:value="note.second.startTime" placeholder="start time"/>
            <input disabled="disabled" type="text" class="form-control" v-bind:value="note.second.endTime" placeholder="end time" />
          </form>
        </div>
      </Collapse>
    </div>
</div>
</body>
</template>

<style scoped>
.header{
    color: white;
    font-size: 48px;
    padding: 20px;
}

body {
    color: black;    
    margin: auto;
    width: 50%;
}

.text{
    padding: 15px;
    font-size: 18px;
}

.screen {
  background: #1abc9c;
  text-align: center;
  margin: auto;
}
</style>