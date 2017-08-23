<script>
	import {Bar, mixins} from 'vue-chartjs'
	const { reactiveProp } = mixins
	export default Bar.extend({
		mixins: [reactiveProp],
		data () {
			return {
				options: {
					scales: {
						xAxes: [ {
							gridLines: {
								display: false
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