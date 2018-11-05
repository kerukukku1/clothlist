<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <div class="imagebox">
      <div class="subtext">
        Your closet is comfortable.
      </div>
    </div>
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

@Component({
  components: {
    HelloWorld,
  },
})

export default class Home extends Vue {
}
</script>

<style scoped>
.imagebox{
  width: 100vws;
  height: 100vh;
  min-height: 100vh;
  background-size: cover;
  background-position: center;
  background-image: url("../assets/cloth.jpg");
  display:  table-cell;
  vertical-align:  middle; 
}

.subtext{
  width: 100vw;
  vertical-align: center;
  color: rgba(255, 255, 255, 0.678);
  font-weight: bold;
  font-size: 100px;
  -webkit-text-stroke: 2px #2c3e50;
}
</style>
