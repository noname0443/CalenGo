<script setup>
import { ref } from 'vue';
import moment from 'moment';
const count = ref(0)

import ApiClient from './api/ApiClient';
import Note from './api/model/Note';
import User from './api/model/User';
import DefaultApi from './api/api/DefaultApi';
import { Modal } from 'usemodal-vue3';

let isVisible = ref(false);
let counter = 1;

function Send() {
  let apiClient = new ApiClient('http://localhost:1516');
  let xxxSvc = new DefaultApi(apiClient);
  let zzz = xxxSvc.getApiV1Note('1', function (error, data, response) {
    console.log(error, data, response)
  });
}

function generateCalendar() {
  counter = 1;
  const startDay = moment().startOf('month').startOf('week');
  const endDay = moment().endOf('month').endOf('week');
  let currentMonth = moment().month();
  
  let calendar = [];
  let index = startDay.clone();
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
  return calendar;
}
</script>

<template>
<main>
  <div class="month">
    <ul>
      <li class="prev">&#10094;</li>
      <li class="next">&#10095;</li>
      <li>{{ moment().format('MMMM') }}<br><span style="font-size:18px">{{ moment().format('YYYY') }}</span></li>
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
      <template v-for="day in generateCalendar(2024, 3)">
        <li v-if="day.isSame(new Date(), 'day') && counter++">
          <span v-on:click="isVisible = !isVisible" class="active">{{ day.format('D') }}</span>
        </li>
        <li v-else-if="counter % 2 === 0 && counter++">
          <p v-on:click="isVisible = !isVisible">{{ day.format('D') }}</p>
        </li>
        <li v-else-if="counter++">
          <p v-on:click="isVisible = !isVisible" class="grayClass">{{ day.format('D') }}</p>
        </li>
      </template>
    </ul>
  </div>
</main>
<Modal v-model:visible="isVisible">
  <div>your content...</div>
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
  color: black !important
}

.days li span {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
}

.grayClass {
  background: #e5e5e5 !important
}

/* Highlight the "current" day */
.days li .active {
  padding: 5px;
  background: #1abc9c;
  color: white !important
}
</style>
