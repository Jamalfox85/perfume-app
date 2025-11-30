<template>
    <div id="app">
        <Header />
        <RouterView class="mt-8 p-4" />
    </div>
</template>

<script>
import Header from "./components/layout/Header.vue";
export default {
    components: {
        Header,
    },
};
</script>

<style scoped>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    min-height: 100vh;
    border: solid 2px blue;
    display: flex;
    flex-direction: column;
}
</style>
