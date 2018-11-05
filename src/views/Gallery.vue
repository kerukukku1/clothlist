<template>
  <div>
    <div class="main-content">
      <div class="head-spacer"></div>
      <div class="flexbox">
        Hello World.
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({
  components: {
  },
})
export default class Gallery extends Vue {
}
</script>

<style scoped>
.main-content{
  background: linear-gradient(25deg, rgb(52, 255, 194), rgba(52, 66, 254, 0.308));
  width: 100vw;
  height: 100vh;
  min-height: 100vh;
}

.flexbox {
  display: flex;
  padding: 20px;
}
</style>
