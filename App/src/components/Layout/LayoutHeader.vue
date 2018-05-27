<template>
  <header class="layout-header">
    <img  class="back" src="~@/assets/icons/Caret-Left.png" alt="Voltar" />
    <h1 class="title">{{ $route.name }}</h1>
  </header>
</template>

<style>
  .layout-header {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    height: 47px;
    background-color: #442B7B;
  }

  .layout-header > .back {
    margin-left: 16px;
  }

  .layout-header > .title {
    color: #ffffff;
    margin-left: calc(50% - (12px + 16px));
    font-size: 20pt;
    font-weight: 400;
    transform: translateX(-50%);
  }
</style>

