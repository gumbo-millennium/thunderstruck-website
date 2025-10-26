<template>
  <div class="relative flex gap-4 h-16 overflow-x-hidden">
    <div class="flex absolute w-full h-full top-0">
      <div class="z-1 w-1/2 bg-linear-to-r from-white to-transparent to-30% pointer-events-none" />
      <div class="z-1 w-1/2 bg-linear-to-r from-transparent to-white from-70% pointer-events-none" />
    </div>
    <div class="flex items-center justify-center gap-4 animate-horizontal-scrolling px-2">
      <a
        v-for="partner, i in props.partners"
        :key="i"
        ref="noreferrer noopener"
        target="_blank"
        :href="partner.link"
        :title="partner.name"
        class="flex justify-center items-center w-16 aspect-square"
      >
        <img
          :src="partner.image"
          :alt="partner.name"
        >
      </a>
    </div>
    <div class="absolute top-0 flex items-center justify-center gap-4 animate-horizontal-scrolling-2 px-2">
      <a
        v-for="partner, i in props.partners"
        :key="i"
        ref="noreferrer noopener"
        target="_blank"
        :href="partner.link"
        :title="partner.name"
        class="flex justify-center items-center w-16 aspect-square"
      >
        <img
          :src="partner.image"
          :alt="partner.name"
        >
      </a>
    </div>
    <div class="absolute top-0 flex items-center justify-center gap-4 animate-horizontal-scrolling-3 px-2">
      <a
        v-for="partner, i in props.partners"
        :key="i"
        ref="noreferrer noopener"
        target="_blank"
        :href="partner.link"
        :title="partner.name"
        class="flex justify-center items-center w-16 aspect-square"
      >
        <img
          :src="partner.image"
          :alt="partner.name"
        >
      </a>
    </div>
  </div>
</template>

<script setup lang="ts">
import type Partner from '~/types/partner';

const props = defineProps<{
  partners: Array<Partner>,
}>();
</script>
