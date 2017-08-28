import index from './components/Index.vue'
import users from './components/Users.vue'
import quotations from './components/Quotations.vue'

export default [
	{ path: '/', component: index },
	{ path: '/quotations', component: quotations },
	{ path: '/users', component: users },
	{ path: '/users/:id', component: users }
]