<template>
	<div>
		<h1 class="h1">{{ indexTitle }}</h1>
		<div class="chart-wrap">
			<chart-bar :chart-data="datacollection" :height="200"></chart-bar>
		</div>
	</div>
</template>

<script>

import ChartBar from './ChartBar.vue'

export default {
	components: {
		'chart-bar': ChartBar,
	},
	data(){
		return{
			indexTitle: 'Index page of my awesome web-application!',
			datacollection: null,
		}
	},
	created(){
		this.BarChartData()
	},
	methods: {
		BarChartData () {
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

<style lang="sass">
	// .chart-wrap
	// 	max-width: 500px
	// 	height: 400px
</style>