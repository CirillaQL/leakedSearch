<template>
  <img alt="logo"  src="./assets/leakedSearchLogo.png">
  <div class="blank"/>
  <SearchInput/>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import SearchInput from './components/SearchInput.vue';

@Options({
  components: {
    SearchInput,
  },
})
export default class App extends Vue {}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.blank{
  margin: 5em auto 5em;
}
</style>
