<template>
  <section class="travel-screen">
    <travel-map
      :end="end"
      :path="path"
      :begin="path && path[0]"
    />

    <travel-button
      :isLoading="isLoading"
      :isPlaying="isPlaying"
      @stop="stop()"
      @start="start()"
    />
  </section>
</template>

<script>
  import TravelMap from '@/components/Travel/TravelMap';
  import TravelButton from '@/components/Travel/TravelButton';

  export default {
    components: { TravelMap, TravelButton },
    props: {
      isLoading: Boolean,
      isPlaying: Boolean,
    },
    methods: {
      stop () {
        this.$emit('stop');
      },
      start () {
        this.$emit('start');
      },
    },
  }
</script>

