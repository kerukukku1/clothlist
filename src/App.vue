<template>
  <div id="app">
    <div class="coolnav" @mouseleave="closeAccordion()" :style="styleObject" v-if="show">
      <div :style="{marginRight: '15px'}">
        <navigation-item v-for="(tab, key) in tabs" :key="tab.id" @open-accordion="openAccordion(key)"> 
          {{key}}
        </navigation-item>
      </div>
      <transition name="ac-content">
        <div v-if="isOpen" class="accordion-content" ref="body">
          <h1>{{contentKey}}</h1>
          <div v-for="tab in tabs[contentKey]" :key="tab.id">
            <router-link :to="tab.link">{{tab.name}}</router-link>
          </div>
        </div>
      </transition>
      <!-- <router-link v-for="tab in tabs" :key=tab.id :to="tab.link" :style="{width: tabWidth}">{{tab.name}}</router-link> -->
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Watch, Component } from 'vue-property-decorator'
import navigationItem from '@/components/NavigationItem.vue'

interface content {
  name : string,
  link : string
}

interface tabs {
  About : Array<content>,
  Management : Array<content>
}
@Component({
  components: {
    'navigation-item' : navigationItem
  }
})
export default class App extends Vue{
  tabs : tabs = {
        About: [
          {name: 'Toppage.', link: '/'},
          {name: 'About.', link: '/about'}
        ],
        Management: [
          {name: 'Gallery.', link: '/gallery'},
          {name: 'Upload.', link: '/upload'},
        ]
  };
  styleObject : {height: string} = {
        height: '60px',
  };
  contentKey : string = '';
  isOpen : boolean = false;
  show : boolean = true;

  @Watch('contentKey')
  onchange(){
    this.$nextTick(function() {
        this.$set(this.styleObject, 'height', (this.isOpen)?this.$refs.body.getBoundingClientRect().height+'px':'60px')
    })
  }

  openAccordion(key : string) {
    this.contentKey = key
    this.isOpen = true
  }

  closeAccordion() {
    this.isOpen = false
    this.contentKey = ''
  }
}
</script>


<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  margin: 0 auto;
}
#nav {
  padding: 20px;
  background: #2c3e50;
  /* margin-bottom: 50px; */
  bottom: 0;
  position: fixed;
  width: 100%;
}

#nav a {
  font-weight: 300;
  /* color: #2c3e50; */
  color: #dd5511;
  padding: 20px;
}

#nav a:hover{
  background: hsl(203, 23%, 27%);
  transition: 400ms;
}

a{
  color: white;
}

a.router-link-exact-active {
  font-weight: 800;
  /* color: #42b983; */
  color: #dd5511;
}

footer{
  height: 100px;
}

.coolnav{
  background: #2c3e50;
  color: white;
  font-weight: bold;
  position: fixed;
  width: 100%;
  transition: height 500ms cubic-bezier(0.68, -0.55,  0.265, 1.55 );
}

.head-spacer{
  height: 40px;
  background-color: #2c3e50;
}

.title-content{
    height: 100px;
    background-color: #2c3e50;
    text-align: center;
    color: aquamarine;
    font-size: 40px;
    font-weight: bold;
    padding: 20px;
}

@media screen and (min-width: 750px){
  .coolnav{
    top: 2%;
    width: 95%;
    left: calc(2.5%);
    border-radius: 35px;
  }
  .head-spacer{
    height: 60px;
    background-color: #2c3e50;
  }
}

.content{
  margin-top: 120px;
}

.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, .5);
  display: table;
  transition: opacity .3s ease;
}

.accordion-content::before{
  content: "";
  width:94%;
  border-top: 2px solid white;
  left: calc(3%);
  top: 80px;
  position: fixed;
}

.accordion-content{
  padding: 70px;
}

.view-enter-active, .view-leave-active {
  animation: bounce-in 500ms;
}

.view-enter, .view-leave-to {
  animation: bounce-in 500ms reverse;
}

.ac-content-enter-active{
  transition: opacity 500ms ease 200ms;
}

.ac-content-enter, .ac-content-leave-to {
  opacity: 0;
}

.accordion-enter-active, .accordion-leave-active {
  transition: opacity 500ms ease 200ms;
}

.accordion-enter, .accordion-leave-to {
  opacity: 0;
}

.expand-transition {
  transition: all 300ms ease;
  height: 30px;
  padding: 10px;
  background-color: #eee;
  overflow: hidden;
}
.expand-enter, .expand-leave {
  height: 0;
  padding: 0 10px;
  opacity: 0;
}


@keyframes bounce-in {
  0%{
    transform: scale(0.5)
  }
  50%{
    transform: scale(1.5)
  }
  100%{
    transform: scale(1.0)
  }
  /* 0%{
    transform: scale(0.5) rotate(0deg);
  }
  100%{
    transform: scale(1.0) rotate(360deg);
  } */
}

</style>