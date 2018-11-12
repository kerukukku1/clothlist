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
          <cloth-box></cloth-box>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import ClothBox from '@/components/ClothBox.vue';
import axios from 'axios'

interface myImage {
  ID : string,
  Title : string,
  Path : string
}

@Component({
  components: {
    'cloth-box' : ClothBox
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
  background-color: #2c3e50;
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
  background-color: #344a5f;
  padding-top: 20px;
  padding-bottom: 20px;
  color: white;
  border-bottom: 2px white solid;
  border-top: 2px white solid;
}

.cloth-gridbox{
  display: grid;
  grid-gap: 10px;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
}

.cloth-item-image{
  width: 100%;
  height: 100%;
  min-height: 100%;
  background-size: cover;
  background-position: center;
  background-image: url("../assets/users/mahito6/IMG_7696.jpg");
}

.cloth-item-detail{

}
</style>
