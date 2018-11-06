<template>
  <div>
    <div class="main-content">
      <div class="head-spacer"></div>
      <div class="title-content">
        Gallery.
      </div>
      <div class="main-gridbox">
        <div class=sidecontent>
          <div class="side-head">
            Clothet.
          </div>
          はろはろ
        </div>
        <div class=maincontent>
          <div class="cloth-gridbox">
            <div class="cloth-item">
              アロハシャツ
            </div>
            <div class="cloth-item">
              パーカー
            </div>
            <div class="cloth-item">
              ジーパン
            </div>
            <div class="cloth-item">
              ソックス
            </div>
          </div>
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

.main-gridbox {
  display: grid;
  grid-gap: 0px;
  grid-template-columns: 2fr 7fr;
}

.sidecontent{
  order: 0;
  background-color: #0000001e;
  height: 620px;
  overflow: scroll;
}

.maincontent{
  order: 1;
  background-color: rgb(230, 222, 130);
  padding: 10px;
  height: 600px;
  overflow: scroll;
}

.side-head{
  width: 100%;
  background-color: azure;
  padding-top: 20px;
  padding-bottom: 20px;
}

.cloth-gridbox{
  display: grid;
  grid-gap: 10px;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
}

.cloth-item{
  background-color: white;
  border-radius: 10px;
  border-color: black;
  padding: 20px;
  text-align: center;
  height: 60px;
}
</style>
