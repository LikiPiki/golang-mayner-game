<script>
	import {Line, mixins} from 'vue-chartjs'
	const { reactiveProp } = mixins
	export default Line.extend({
		mixins: [reactiveProp],
		data () {
			return {
				options: {
					scales: {
						xAxes: [ {
							gridLines: {
								display: false
							},
							ticks: {
								autoSkip: false,
								maxRotation: 90,
								minRotation: 90
							},
							categoryPercentage: 1.0,
							barPercentage: 0.5
						},]
					},
					legend: {
						display: true
					},
					responsive: true,
					maintainAspectRatio: false
				}
			}
		},
		mounted () {
			this.renderChart(this.chartData, this.options)
		}
	})
</script>