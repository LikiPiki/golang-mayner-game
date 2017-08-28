<template>
	<div>
		<p>{{ wow }}</p>
		<chart-line :chart-data="datacollection" :height="200"></chart-line>
	</div>
</template>

<script>

import ChartLine from './ChartLine.vue'

export default {
	components: {
		'chart-line': ChartLine,
		datacollection: null,
	},
	data(){
		return{
			wow: 'Wooooowww!!',
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
			let lineDate = ''
			this.$http.get('https://min-api.cryptocompare.com/data/histoday?fsym=BTC&tsym=USD&limit=30').then(function(data){
				data.body.Data.forEach((day) => {
					lineDate =(new Date(day.time * 1000))
					labels.push(lineDate.toLocaleString("ru", {
						year: '2-digit',
						month: 'numeric',
						day: 'numeric'
					}) + 'г.')
					chartData.push(day.close)
				});
				this.datacollection = {
					labels: labels,
					datasets: [{
						label: 'BTC to USD',
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