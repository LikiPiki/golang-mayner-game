<template>
	<div>
		<h1 class="h1">{{ indexTitle }}</h1>
		<chart :chart-data="datacollection"></chart>
	</div>
</template>

<script>

import Chart from './Chart.vue'

export default {
	components: {
		'chart': Chart,
	},
	data(){
		return{
			indexTitle: 'Index page of my awesome web-application!',
			datacollection: null,
		}
	},
	created(){
		this.fillData()
	},
	methods: {
		fillData () {
			let labels = []
			let chartData = []
			this.$http.get('https://jsonplaceholder.typicode.com/users').then(function(data){
				data.body.forEach((user) => {
					labels.push(user.username)
					chartData.push(user.id)
				});
				this.datacollection = {
					labels: labels,
					datasets: [{
						label: 'Намайнено BTC',
						backgroundColor: '#f87979',
						data: chartData
					}]
				}
			});
		}
	}
}
</script>

<style>

</style>