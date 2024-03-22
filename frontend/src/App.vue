<script setup>
</script>

<template>
<body>
<div id="block_container">
    <b-navbar id="nav-block" type="dark" variant="dark">
      <b-navbar-nav>
        <RouterLink class="m-2 btn" to="/">Home</RouterLink>
        <RouterLink class="m-2 btn" to="/calendar">Calendar</RouterLink>
        <RouterLink class="m-2 btn" to="/user">User</RouterLink>
      </b-navbar-nav>
    </b-navbar>
</div>
<main>
  <router-view v-slot="{ Component, route }">
    <div id="wrapper">
      <Transition name="fade" mode="out-in">
        <div style="padding: 20px;" :key="route.name">
          <component :is="Component"></component>
        </div>
      </Transition>
    </div>
  </router-view>
</main>
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
  margin: 0 auto;
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