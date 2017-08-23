import index from './components/Index.vue'
import users from './components/Users.vue'

export default [
	{ path: '/', component: index },
	{ path: '/users', component: users },
	{ path: '/users/:id', component: users }
]