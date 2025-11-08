<template>
  <a
    ref="noreferrer noopener"
    target="_blank"
    :href="props.url"
    class="font-archivo group flex justify-center items-center relative hover:cursor-pointer"
  >
    <div class="flex justify-center items-center z-1 bg-purple w-80 h-12 text-white uppercase font-[1000] text-3xl group-hover:bg-green-secondary group-focus:bg-white group-hover:text-purple group-focus:text-purple group-hover:z-0 group-focus:z-0 transition-all duration-100">
      <slot />
    </div>
    <div class="group-hover:ml-0 group-focus:ml-0 group-hover:mt-0 group-focus:mt-0 group-hover:z-1 group-focus:z-1 absolute ml-2 mt-2 border-purple border-1 w-80 h-12 transition-all duration-100 delay-100"/>
  </a>
</template>

<script setup lang="ts">
const props = defineProps<{
  url: string,
}>();
</script>
