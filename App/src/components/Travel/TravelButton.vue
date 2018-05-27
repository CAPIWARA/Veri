<template>
  <fluffy-button type="button" :disabled="isLoading" @click="click()">
    {{ text }}
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
    computed: {
      text () {
        const text = this.isLoading ? 'Carregando'
                   : this.isPlaying ? 'Parar'
                                    : 'Começar';
        return text;
      },
    },
    methods: {
      click () {
        const event = this.isPlaying ? 'stop' : 'start';
        this.$emit(event);
      },
    },
  }
</script>

