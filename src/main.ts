import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './registerServiceWorker';

Vue.config.productionTip = false;


// client.connect(function (err){
//   console.log("connected!");
//   const db = client.db("test")
//   console.log(db)
//   client.close();
// })

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
