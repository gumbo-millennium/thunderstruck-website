<template>
  <component
    :is="props.link ? 'a' : 'div'"
    ref="noreferrer noopener"
    target="_blank"
    :href="props.link"
    class="group uppercase relative"
  >
    <div class="absolute left-[-.5rem] top-[-1rem] bg-purple w-14 text-center font-archivo text-white group-hover:bg-white group-focus:bg-white group-hover:border-2 group-focus:border-2 group-hover:border-purple group-focus:border-purple group-hover:text-purple group-focus:text-purple transition-all duration-100 delay-75">
      {{ props.time }}
    </div>
    <div class="bg-green-primary w-full px-3 py-2 text-right font-archivo text-white text-xl group-hover:bg-purple group-focus:bg-purple transition-all duration-100">
      <h2>{{ props.title }}</h2>
    </div>
    <div class="hidden group-hover:block group-focus:block bg-green-secondary w-full text-xs text-right py-0.5 px-3 group-hover:bg-white group-focus:bg-white">
      Luister naar deze band op Spotify
    </div>
    <div class="group-hover:hidden group-focus:hidden bg-green-secondary w-full text-xs text-right py-0.5 px-3">
      <span
        v-for="genre, i in props.genres"
        :key="i"
      >
        {{ genre.name }}

        <span
          v-if="props.genres.length > 1 && i < props.genres.length - 1"
        >
          |
        </span>
      </span>
    </div>
  </component>
</template>

<script setup lang="ts">
import type Genre from '~/types/genre';

const props = defineProps<{
  time: string,
  title: string,
  genres: Array<Genre>,
  link?: string,
}>();
</script>
