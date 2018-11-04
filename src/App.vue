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
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import {Component, Prop} from 'vue-property-decorator'
@Component({
})
export default class App extends Vue{
  tabs : Object = {
        About: [
          {name: 'Home', link: '/'},
          {name: 'About', link: '/about'},
        ],
        Company: [
          {name: 'representative', link: '/'}
        ]
  };
  styleObject : Object = {
        height: '60px',
  };
  contentKey : string = '';
  isOpen : boolean = false;
  show : boolean = true;

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
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
