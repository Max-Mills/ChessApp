<template>
  <div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/chess">Chess Game</router-link> |
    <router-link to="/about">About</router-link>
  </div>
  <router-view/>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

.active {
	background-color: rgba(99, 55, 158, 0.50);
	border-collapse: collapse;
}

img {
	transition: transform 0.75s;
	transition-timing-function: ease-in-out;
}

.logo {
  height: 100px;
	width: 100px;
}

.boardBorder {
	width: 650px;
	height: 650px;
	padding-top: 25px;
	background-image: url("http://localhost:8082/assets/redmahogany.jpg");
	background-repeat: no-repeat;
	margin-left: auto;
	margin-right: auto;
	border-radius: 16px;
	align-items: center;
	justify-content: center;
}

.board {
	width: 600px;
	height: 600px;
	border-radius: 2px;
	background-image: url("http://localhost:8082/assets/whitetiles.jpg");
	position: flex;
	margin-left: auto;
	margin-right: auto;
	border-collapse: collapse;
}

table tr td img{
	table-layout: fixed;
	height: 70px;
	width: 70px;
}

td:hover :not(img) {
	background-color: rgba(169, 141, 206, 0.25);
}

table .row:nth-child(odd) .column:nth-child(even) {
background-image: url("http://localhost:8082/assets/darktiles.jpg");
  background-repeat: no-repeat;
}
table .row:nth-child(even) .column:nth-child(odd) {
background-image: url("http://localhost:8082/assets/darktiles.jpg");
  background-repeat: no-repeat;
}

</style>
