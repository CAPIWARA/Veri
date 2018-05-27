<template>
  <section class="travel-screen">
    <travel-map
      v-if="isPlaying"
      :end="end"
      :path="path"
      :begin="begin"
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

