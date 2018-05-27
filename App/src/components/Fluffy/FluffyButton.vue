<template>
  <button v-on="$listeners" v-bind="$attrs" class="fluffy-button">
    <slot />
  </button>
</template>
