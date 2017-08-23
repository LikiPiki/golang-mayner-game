import Vue from 'vue'
import App from './App.vue'
import Routes from './routes.js'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'

Vue.use(VueRouter);
Vue.use(VueResource);

const router = new VueRouter({
	routes: Routes,
	mode: 'history'
});

new Vue({
  el: '#app',
	render: h => h(App),
	router: router
});
