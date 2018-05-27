<template>
  <section class="travel-screen">
    <travel-button
      :isLoading="isLoading"
      :isPlaying="isPlaying"
      @stop="stop()"
      @start="start()"
    />

    <travel-map
      v-if="isPlaying"
      :end="end"
      :path="path"
      :begin="begin"
    />
  </section>
</template>

<script>
  import { toPath } from '@/services/travel';
  import TravelMap from '@/components/Travel/TravelMap';
  import TravelButton from '@/components/Travel/TravelButton';

  export default {
    components: { TravelMap, TravelButton },
    props: {
      isLoading:   Boolean,
      isPlaying:   Boolean,
      end:         { type: Object, default: () => null },
      begin:       { type: Object, default: () => null },
      coordinates: { type: Array,  required: true      },
    },
    computed: {
      path () {
        return this.coordinates.map(toPath);
      }
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

<style>
  .travel-screen {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }
</style>
