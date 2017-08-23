<template>
	<header>
		<div class="header__content">
			<router-link to="/" exact>
				<img src="../assets/logo.jpg" alt="Logo" width="100">
			</router-link>
			<router-link to="/" exact>Главная</router-link>
			<router-link to="/users">Топ майнеры</router-link>
			<router-link to="/quotations">Котировки</router-link>
			<router-link to="/contact">Контакты</router-link>
		</div>
	</header>
</template>

<script>

export default {
	data(){
		return{

		}
	}
}
</script>

<style lang="sass">
	*
		box-sizing: border-box
	body
		margin: 0
		font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif
		color: #333
	.router-link-active
		color: tomato
</style>