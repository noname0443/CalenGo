<script setup>
import { reactive, ref, toRaw } from 'vue';
import moment from 'moment';
import VueCookies from 'vue-cookies'

import "./assets/form-app.css";
import ApiClient from './api/ApiClient';
import DefaultApi from './api/api/DefaultApi';
import { Modal } from 'usemodal-vue3';
import ListApiV1NoteRequest from './api/model/ListApiV1NoteRequest'
import { CSpinner } from '@coreui/vue'
import { Collapse } from 'vue-collapsed'

moment.updateLocale('en', {
  week: {
    dow : 1,
  }
});

let StartSelection = ref(-1);
let EndSelection = ref(-1);

let isVisible = ref(false);
let isWarningMessage = ref(false);

let MONTH_FROM_CHRIST = ref(moment().year() * 12 + moment().month())

let BACKEND_URL = 'http://localhost:1516'
let global_notes = reactive([])
let global_status = ref(null);
let DatesArray = [];

let global_note = reactive({
  name: '',
  description: '',
  startTime: '',
  endTime: ''
});

function DoStartSelection(num) {
  StartSelection._value = num;
  EndSelection._value = num;
  ChangeColor()
}

function ChangeColor() {
  removeColor()
  console.log(StartSelection._value, EndSelection._value)
  if(EndSelection._value == -1 || StartSelection._value == -1) {
    return
  }
  let start = StartSelection._value;
  let end = EndSelection._value;
  if(end < start) {
    let tmp = end;
    end = start;
    start = tmp;
  }
  for(let i = start; i <= end; i++) {
    document.getElementById(`i${i}`).getElementsByTagName('p')[0].style.background = 'rgba(62, 155, 141, 0.3)';
    document.getElementById(`i${i}`).getElementsByTagName('p')[0].style.color = 'black';
  }
}

function removeColor() {
  for(let i = 0; i <= 60; i++) {
    if (document.getElementById(`i${i}`) !== null)
      document.getElementById(`i${i}`).getElementsByTagName('p')[0].style = '';
  }
}

function UnChangeColor() {
  if(EndSelection._value == -1 || StartSelection._value == -1) {
    return
  }
  isVisible.value = true
  listNotes(DatesArray[StartSelection._value], DatesArray[EndSelection._value])
  console.log(isVisible)
  removeColor()
  StartSelection._value = -1;
  EndSelection._value = -1;
}

window.addEventListener('mouseup', function(e) {
  e.preventDefault()
  removeColor()
  StartSelection._value = -1;
  EndSelection._value = -1;
}, false);

function getMonthString(month_from_christ) {
  let year = month_from_christ / 12
  let month = month_from_christ % 12 + 1
  return moment(`${year}-${month}-01`, 'YYYY-M-D').format('MMMM')
}

function getYearString(month_from_christ) {
  let year = month_from_christ / 12
  let month = month_from_christ % 12 + 1
  return moment(`${year}-${month}-01`, 'YYYY-M-D').format('YYYY')
}

function Get(note) {
  let apiClient = new ApiClient(BACKEND_URL);
  let xxxSvc = new DefaultApi(apiClient);
  let zzz = xxxSvc.getApiV1Note(note, function (error, data, response) {
    console.log(error, data, response)
  });
}

function Post() {
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
  let req = defaultApi.postApiV1Note({'note': global_note}, function (error, data, response) {
    console.log(error, data, response)
    if(error != null) {
      global_status.value = "Failed: " + error;
    } else {
      global_status.value = "Success";
    }
    listNotes(moment(global_note.startTime), moment(global_note.endTime))
  });
}

function listNotes(start, end) {
  if(start === undefined || end === undefined) {
    return
  }
  global_note.startTime = start.format("YYYY-MM-DD")
  global_note.endTime = end.format("YYYY-MM-DD")

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

  let ListReqStruct = new ListApiV1NoteRequest(global_note.startTime, global_note.endTime)
  let request = {'listApiV1NoteRequest': ListReqStruct}

  let req = defaultApi.listApiV1Note(request, function (error, data, response) {
    console.log(error, data, response)
    data.forEach(function (element) {
      element.collapse = false;
    });
    global_notes.length = 0;
    global_notes.push(...data);
    console.log(global_notes)
  });
}

function generateCalendar(month_from_christ) {
  let year = month_from_christ / 12
  let month = month_from_christ % 12 + 1
  
  let date = moment(`${year}-${month}-01`, 'YYYY-MM-DD')
  
  const startDay = date.clone().startOf('month').startOf('week');
  const endDay = date.clone().endOf('month').endOf('week');
  let currentMonth = date.clone().month();
  
  let calendar = [];
  let index = startDay.clone().subtract(1, 'day');
  while (index.isBefore(endDay, 'day')) {
    let arr = new Array(7).fill(0).map(function(n, i) {
      return index.add(1, 'day').clone();
    });
      
    if (arr[0].month() === currentMonth || arr[6].month() === currentMonth) {
      for (let i = 0; i < arr.length; i++) {
        calendar.push(arr[i]);
      }
    }
  }
  DatesArray = calendar;
  return calendar;
}
</script>

<template>
<main>
  <div class="month">
    <ul>
      <li class="prev prevent-select" v-on:click="MONTH_FROM_CHRIST--">&#10094;</li>
      <li class="next prevent-select" v-on:click="MONTH_FROM_CHRIST++">&#10095;</li>
      <li>{{ getMonthString(MONTH_FROM_CHRIST) }}<br><span style="font-size:18px">{{ getYearString(MONTH_FROM_CHRIST) }}</span></li>
    </ul>
  </div>
  
  <div class="days-view">
    <ul class="weekdays">
      <li>Monday</li>
      <li>Tuesday</li>
      <li>Wednesday</li>
      <li>Thursday</li>
      <li>Friday</li>
      <li>Saturday</li>
      <li>Sunday</li>
    </ul>

    <ul class="days">
      <template v-for="[i, day] in generateCalendar(MONTH_FROM_CHRIST).entries()">
        <li v-if="day.isSame(new Date(), 'day')" :id="['i' + i]" v-on:mouseup="UnChangeColor()" v-on:mousedown="DoStartSelection(i)" v-on:mouseenter="EndSelection=i;ChangeColor()">
          <p class="active prevent-select">{{ day.format('D') }}</p>
        </li>
        <li v-else-if="i % 2 === 0" :id="['i' + i]" v-on:mouseup="UnChangeColor()" v-on:mousedown="DoStartSelection(i)" v-on:mouseenter="EndSelection=i;ChangeColor()">
          <p class="prevent-select">{{ day.format('D') }}</p>
        </li>
        <li v-else :id="['i' + i]" v-on:mouseup="UnChangeColor()" v-on:mousedown="DoStartSelection(i)" v-on:mouseenter="EndSelection=i;ChangeColor()">
          <p class="grayBackground prevent-select">{{ day.format('D') }}</p>
        </li>
      </template>
    </ul>
  </div>
</main>

<Modal title="Notes Window" width="70%" v-model:visible="isVisible" :okButton="{text: 'ok', onclick: () => {isWarningMessage = true; global_status = null; Post()}, loading: false}">
  <ul class="ul-popup">
    <form class="form-signin">
      <h2 class="form-signin-heading">Create Note</h2>
      <input type="text" class="form-control" v-model="global_note.name" placeholder="name" />
      <textarea rows="8" class="form-textarea" v-model="global_note.description" placeholder="description"/>
      <input type="text" class="form-control" v-model="global_note.startTime" placeholder="start time" />
      <input type="text" class="form-control" v-model="global_note.endTime" value="{{global_note.endTime}}" placeholder="end time" />
    </form>
    <div v-for="(note, index) in global_notes" style="padding-top: 15px;">
      <h2 @click="global_notes[index].collapse = !global_notes[index].collapse" class="form-signin">
        <template v-if="!global_notes[index].collapse">
          + {{ note.name }}
        </template>
        <template v-else>
          - {{ note.name }}
        </template>
      </h2>
      <Collapse :when="global_notes[index].collapse">
        <form class="form-signin">
          <textarea disabled="disabled" rows="8" class="form-textarea" v-bind:value="note.description" placeholder="description"/>
          <input disabled="disabled" type="text" class="form-control" v-bind:value="note.startTime" placeholder="start time"/>
          <input disabled="disabled" type="text" class="form-control" v-bind:value="note.endTime" placeholder="end time" />
        </form>
      </Collapse>
    </div>
  </ul>
</Modal>
  <Modal title="Status Window" width="40%" v-model:visible="isWarningMessage" :okButton="{text: 'ok', onclick: () => {isWarningMessage = false}}">
    <div v-if="global_status == null" class="d-flex justify-content-center mb-3">
      <b-spinner></b-spinner>
    </div>
    <p v-else>{{ global_status }}</p>
  </Modal>
</template>
<style scoped>
div {
  color: black;
} 
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

ul {
  list-style-type: none;
}
body {font-family: Verdana, sans-serif;}

.ul-popup {
  list-style-type: decimal;
}

/* Month header */
.month {
  padding: 40px 25px;
  width: 50%;
  background: #1abc9c;
  text-align: center;
  margin: auto;
}

.days-view {
  height: 100%;
  width: 50%;
  margin: auto;
}

/* Month list */
.month ul {
  margin: 0;
  padding: 0;
}

.month ul li {
  color: white;
  font-size: 20px;
  text-transform: uppercase;
  letter-spacing: 3px;
}

/* Previous button inside month header */
.month .prev {
  float: left;
  padding-top: 10px;
}

/* Next button */
.month .next {
  float: right;
  padding-top: 10px;
}

/* Weekdays (Mon-Sun) */
.weekdays {
  margin: 0;
  padding: 10px 0;
  background-color:#ddd;
}

.weekdays li {
  display: inline-block;
  width: 14.2857142857%;
  color: #666;
  text-align: center;
}

/* Days (1-31) */
.days {
  padding: 10px 0;
  background: #eee;
  margin: 0;
  padding-bottom: 0;
  padding-top: 0;
}

.days li {
  aspect-ratio : 1 / 1;
  list-style-type: none;
  display: inline-block;
  width: 14.2857142857%;
  text-align: center;
  font-size: 24px;
  color: #777;
  vertical-align: middle;
}

.days li p {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
  color: black
}

.days li span {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
}

.grayBackground {
  background: #e5e5e5;
}

/* Highlight the "current" day */
.days li .active {
  padding: 5px;
  background: #1abc9c;
  color: white
}

.prevent-select {
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
  cursor: pointer;
}

.v-collapse {
  transition: height 300ms ease-out;
  /* or transition: height var(--vc-auto-duration) ease-in-out */
}
</style>
