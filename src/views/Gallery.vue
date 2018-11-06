<template>
  <div>
    <div class="main-content">
      <div class="head-spacer"></div>
      <div class="title-content">
        Gallery.
      </div>
      <div class="gridbox">
        <div class=sidecontent>
          ここにさいどばーがくるよ
        </div>
        <div class=maincontent>
          ここに洋服がくるよ
        </div>
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
  background: linear-gradient(25deg, rgb(52, 255, 194), rgba(84, 153, 244, 0.308));
  width: 100vw;
  height: 100vh;
  min-height: 100vh;
}

.gridbox {
  display: grid;
  grid-gap: 10px;
  grid-template-columns: 1fr 3fr;
}

.sidecontent{
  order: 0;
}

.maincontent{
  order: 1;
  background-color: red;
}
</style>
