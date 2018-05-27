<template>
  <fluffy-button type="button" :isLoading="isLoading" @click="click()">
    <img
      v-if="isPlaying"
      class="image"
      slot="icon"
      src="~@/assets/icons/Pause.png"
      alt="Parar"
    />
    <img
      v-else
      class="image"
      slot="icon"
      src="~@/assets/icons/Play.png"
      alt="Iniciar"
    />
    {{ isPlaying ? null : 'GO' }}
  </fluffy-button>
</template>

<script>
  import FluffyButton from '@/components/Fluffy/FluffyButton';

  export default {
    components: { FluffyButton },
    props: {
      isLoading: Boolean,
      isPlaying: Boolean,
    },
    methods: {
      click () {
        const event = this.isPlaying ? 'stop' : 'start';
        this.$emit(event);
      },
    },
  }
</script>

