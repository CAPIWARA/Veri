<template>
  <section class="layout-container">
    <layout-header />
    <router-view />
  </section>
</template>

<script>
  import LayoutHeader from '@/components/Layout/LayoutHeader';

  export default {
    components: { LayoutHeader },
  };
</script>

<style>
  * {
    box-sizing: border-box;
    margin: 0;
  }

  html {
    height: 100%;
    font-size: 62.5%;
  }

  body {
    height: 100%;
    font-size: 1.8rem;
    font-family: 'Montserrat', sans-serif;
  }

  .layout-container {
    display: flex;
    flex-direction: column;
    height: 100%;
  }
</style>

