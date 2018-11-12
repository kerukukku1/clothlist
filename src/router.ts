import Vue from 'vue';
import Router from 'vue-router';
import Top from './views/Top.vue';
import Gallery from './views/Gallery.vue'
import Upload from '@/views/Upload.vue'
import Detail from '@/views/Detail.vue'

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'top',
      component: Top,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
    },
    {
      path: '/gallery',
      name: 'gallery',
      component : Gallery
    },
    {
      path: '/upload',
      name: 'upload',
      component : Upload
    },
    {
      path: '/gallery/:id',
      name: 'detail',
      component : Detail
    }
  ],
});
